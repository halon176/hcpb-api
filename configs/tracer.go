package configs

import (
	"context"
	"log/slog"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer(ctx context.Context) (shutdown func(context.Context) error, err error) {
	if LOGFIRE_TOKEN == "" {
		slog.Warn("LOGFIRE_TOKEN not set, tracing disabled")
		return func(context.Context) error { return nil }, nil
	}

	endpoint := "logfire-eu.pydantic.dev"
	if strings.Contains(LOGFIRE_TOKEN, "_us_") {
		endpoint = "logfire-us.pydantic.dev"
	}

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithHeaders(map[string]string{
			"Authorization": LOGFIRE_TOKEN,
		}),
	)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("hcpb-api"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	slog.Info("OpenTelemetry tracer initialized", "endpoint", endpoint)

	return tp.Shutdown, nil
}
