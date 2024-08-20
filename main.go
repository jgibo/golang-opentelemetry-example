package main

import (
	"context"

	"github.com/lainio/err2/try"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	provider := trace.NewTracerProvider(
		// The span processor. WithBatcher initialises a BatchSpanProcessor
		trace.WithBatcher(
			// The SpanProcessor needs an exporter, where it exports the spans after processing them
			// This exporter uses the open telemetry GRPC exporter. HTTP json exporter is the other commonly used exporter type.
			try.To1(otlptracegrpc.New(
				context.TODO(),
				otlptracegrpc.WithEndpointURL("http://localhost:4317"),
			)),
		),
		// The sampler to use. AlwaysSample samples all spans.
		trace.WithSampler(trace.AlwaysSample()),
	)
	defer provider.Shutdown(context.Background())

	// Register the provider as the global tracer provider. Similiar pattern to how a global logger is set and used.
	otel.SetTracerProvider(provider)

	// get a tracer
	ctx := context.Background() // this would usually be the request context
	tracer := otel.Tracer("main")

	// start a trace span
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	// Call some functions that will create child spans in this trace. The context contains the trace context that otel will use to connect the spans.
	a(ctx)
	b(ctx)
}

func a(ctx context.Context) {
	tracer := otel.Tracer("main")
	ctx, span := tracer.Start(ctx, "a")
	defer span.End()

	aa(ctx)
}

func aa(ctx context.Context) {
	tracer := otel.Tracer("main")
	ctx, span := tracer.Start(ctx, "aa")
	defer span.End()
}

func b(ctx context.Context) {
	tracer := otel.Tracer("main")
	ctx, span := tracer.Start(ctx, "b")
	defer span.End()

	ba(ctx)
}

func ba(ctx context.Context) {
	tracer := otel.Tracer("main")
	ctx, span := tracer.Start(ctx, "ba")
	defer span.End()
}
