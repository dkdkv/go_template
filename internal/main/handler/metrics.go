package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"

	"github.com/go-faster/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Metrics struct {
	meterProvider  metric.MeterProvider
	tracerProvider trace.TracerProvider
	prometheus     http.Handler
	resource       *resource.Resource
	srv            *http.Server
}

type Config struct {
	Addr string // Metrics server address
	Name string // Service name
}

func (m *Metrics) MeterProvider() metric.MeterProvider {
	return m.meterProvider
}

func (m *Metrics) TracerProvider() trace.TracerProvider {
	return m.tracerProvider
}

func NewMetrics(log *zap.Logger, cfg Config) (*Metrics, error) {
	if cfg.Addr == "" {
		cfg.Addr = os.Getenv("METRICS_ADDR")
	}
	if cfg.Addr == "" {
		cfg.Addr = "0.0.0.0:8080"
	}

	res, err := resource.New(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create resource")
	}

	// Prometheus setup
	exp, err := prometheus.New()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Prometheus exporter")
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(exp),
	)
	otel.SetMeterProvider(meterProvider)

	// Tracing setup
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tracerProvider)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var b strings.Builder
		b.WriteString("Service is up and running.\n\n")
		b.WriteString("Available endpoints:\n")
		b.WriteString("/metrics - Prometheus metrics\n")
		b.WriteString("/debug/pprof - Profiling\n")
		_, _ = fmt.Fprintln(w, b.String())
	})

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	log.Info("Metrics server initialized", zap.String("addr", cfg.Addr))

	return &Metrics{
		meterProvider:  meterProvider,
		tracerProvider: tracerProvider,
		prometheus:     promhttp.Handler(),
		resource:       res,
		srv:            srv,
	}, nil
}

func (m *Metrics) Run(ctx context.Context) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return m.srv.ListenAndServe()
	})

	wg.Go(func() error {
		<-ctx.Done()
		return m.srv.Shutdown(context.Background())
	})

	return wg.Wait()
}
