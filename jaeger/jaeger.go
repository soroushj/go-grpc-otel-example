// See https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go

package jaeger

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// TracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func TracerProvider(url, service string, attrs ...attribute.KeyValue) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	attrs = append(attrs, semconv.ServiceNameKey.String(service))
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production
		trace.WithBatcher(exp),
		// Record information about this application in a Resource
		trace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, attrs...)),
	)
	return tp, nil
}
