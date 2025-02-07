package main

import (
	"context"
	"errors"
	"flag"
	"go_template/config"
	"go_template/internal/main/handler"
	"go_template/internal/oas"
	"go_template/internal/pkg/repository"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	handler.Run(func(ctx context.Context, lg *zap.Logger) error {
		// Parse CLI arguments
		configPath := flag.String("config", "config/config.yml", "path to config file")
		addr := flag.String("addr", "0.0.0.0:8000", "server listen address")
		metricsAddr := flag.String("metrics.addr", "0.0.0.0:8080", "metrics listen address")
		flag.Parse()

		// Load configuration
		_, err := config.New(*configPath)
		if err != nil {
			lg.Fatal("Failed to load config", zap.Error(err))
		}

		// Connect to databases
		dbPool, err := repository.ConnectDB()
		if err != nil {
			lg.Fatal("Unable to connect to database: %v\n", zap.Error(err))
		}
		defer dbPool.Close()

		// Initialize metrics
		m, err := handler.NewMetrics(lg, handler.Config{
			Addr: *metricsAddr,
			Name: "api",
		})
		if err != nil {
			return err
		}

		// Set up OpenTelemetry global providers
		otel.SetTracerProvider(m.TracerProvider())
		otel.SetMeterProvider(m.MeterProvider())

		// Set up handler
		params := handler.Params{
			Logger: lg,
		}
		h := handler.New(params, dbPool)

		oasServer, err := oas.NewServer(h,
			oas.WithTracerProvider(m.TracerProvider()),
			oas.WithMeterProvider(m.MeterProvider()),
			oas.WithPathPrefix("/api/v1"),
		)
		if err != nil {
			return err
		}

		httpServer := http.Server{
			Addr:    *addr,
			Handler: oasServer,
		}

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			return m.Run(ctx)
		})
		g.Go(func() error {
			<-ctx.Done()
			lg.Info("Shutting down HTTP server...")
			return httpServer.Shutdown(ctx)
		})
		g.Go(func() error {
			lg.Info("Starting HTTP server", zap.String("address", *addr))
			if err := httpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
				lg.Error("HTTP server error", zap.Error(err))
				return err
			}
			return nil
		})

		return g.Wait()
	})
}
