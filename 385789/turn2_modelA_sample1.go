package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	// Configure Jaeger
	jaegerConfig := &config.Configuration{
		ServiceName: "my-service",
		// Agent: &config.Agent{
		// Endpoint: "localhost:6831",
		// },
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1, // Sample 100% of requests
		},
	}

	tracer, closer, err := jaegerConfig.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	// Your code goes here
	processData := func(ctx context.Context, data []byte) {
		span, ctx := opentracing.StartSpanFromContext(ctx, "processData")
		defer span.Finish()

		span.SetTag("data.size", len(data))

		// Process data here
	}

	// Example usage
	data := []byte("Some large dataset")
	ctx := context.Background()
	processData(ctx, data)
}
