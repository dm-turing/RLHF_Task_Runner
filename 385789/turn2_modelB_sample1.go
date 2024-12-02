package main

import (
	"github.com/jaegertracing/jaeger-client-go/config"
)

func main() {
	// Initialise Jaeger Tracer
	cfg := config.Configuration{
		ServiceName: "your-service-name",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1.0, // Sample 100% of requests
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "localhost:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()
}
