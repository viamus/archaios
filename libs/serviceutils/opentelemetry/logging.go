package opentelemetry

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

func LoggerWithTrace(ctx context.Context) *slog.Logger {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	return slog.Default().With(
		slog.String("trace_id", sc.TraceID().String()),
		slog.String("span_id", sc.SpanID().String()),
	)
}
