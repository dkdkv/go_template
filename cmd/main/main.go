package main

import (
	"context"
	"errors"
	"flag"
	"go_template/internal/main/handler"
	"go_template/internal/oas"
	"go_template/internal/pkg/repository"
	"log"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	handler.Run(func(ctx context.Context, lg *zap.Logger) error {
		dbPool, err := repository.ConnectDB()
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		defer dbPool.Close()
		var arg struct {
			Addr        string
			MetricsAddr string
		}
		flag.StringVar(&arg.Addr, "addr", "0.0.0.0:8000", "listen address")
		flag.StringVar(&arg.MetricsAddr, "metrics.addr", "0.0.0.0:8080", "metrics listen address")
		flag.Parse()

		lg.Info("Initializing",
			zap.String("http.addr", arg.Addr),
			zap.String("metrics.addr", arg.MetricsAddr),
		)

		m, err := handler.NewMetrics(lg, handler.Config{
			Addr: arg.MetricsAddr,
			Name: "api",
		})
		if err != nil {
			return err
		}

		params := handler.Params{
			Logger: lg,
		}
		h := handler.New(params, dbPool)

		oasServer, err := oas.NewServer(h,
			oas.WithTracerProvider(m.TracerProvider()),
			oas.WithMeterProvider(m.MeterProvider()),
			oas.WithPathPrefix("/api"),
		)

		if err != nil {
			return err
		}
		httpServer := http.Server{
			Addr:    arg.Addr,
			Handler: oasServer,
		}

		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			return m.Run(ctx)
		})
		g.Go(func() error {
			<-ctx.Done()
			return httpServer.Shutdown(ctx)
		})
		g.Go(func() error {
			defer lg.Info("Server stopped")
			if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		})

		return g.Wait()
	})
}
