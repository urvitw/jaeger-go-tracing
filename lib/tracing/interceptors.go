package tracing

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.

// request to span ke tracer me dalta hai
func Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.

// tracer se span context nikalta hai
func Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
}
