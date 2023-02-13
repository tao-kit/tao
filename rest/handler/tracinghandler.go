package handler

import (
	"github.com/sllt/tao/rest/internal/response"
	"net/http"
	"sync"

	"github.com/sllt/tao/core/lang"
	"github.com/sllt/tao/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var notTracingSpans sync.Map

// DontTraceSpan disable tracing for the specified span name.
func DontTraceSpan(spanName string) {
	notTracingSpans.Store(spanName, lang.Placeholder)
}

// TracingHandler return a middleware that process the opentelemetry.
func TracingHandler(serviceName, path string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		propagator := otel.GetTextMapPropagator()
		tracer := otel.GetTracerProvider().Tracer(trace.TraceName)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			spanName := path
			if len(spanName) == 0 {
				spanName = r.URL.Path
			}

			if _, ok := notTracingSpans.Load(spanName); ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			spanCtx, span := tracer.Start(
				ctx,
				spanName,
				oteltrace.WithSpanKind(oteltrace.SpanKindServer),
				oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(
					serviceName, spanName, r)...),
			)
			defer span.End()

			// convenient for tracking error messages
			propagator.Inject(spanCtx, propagation.HeaderCarrier(w.Header()))

			trw := &response.WithCodeResponseWriter{Writer: w, Code: http.StatusOK}
			next.ServeHTTP(trw, r.WithContext(spanCtx))

			span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(trw.Code)...)
			span.SetStatus(semconv.SpanStatusFromHTTPStatusCodeAndSpanKind(trw.Code, oteltrace.SpanKindServer))
		})
	}
}
