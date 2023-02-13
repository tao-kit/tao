package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	ztrace "github.com/sllt/tao/core/trace"
	"github.com/sllt/tao/rest/chain"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func TestOtelHandler(t *testing.T) {
	ztrace.StartAgent(ztrace.Config{
		Name:     "go-tao-test",
		Endpoint: "http://localhost:14268/api/traces",
		Batcher:  "jaeger",
		Sampler:  1.0,
	})
	defer ztrace.StopAgent()

	for _, test := range []string{"", "bar"} {
		t.Run(test, func(t *testing.T) {
			h := chain.New(TracingHandler("foo", test)).Then(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					span := trace.SpanFromContext(r.Context())
					assert.True(t, span.SpanContext().IsValid())
					assert.True(t, span.IsRecording())
				}))
			ts := httptest.NewServer(h)
			defer ts.Close()

			client := ts.Client()
			err := func(ctx context.Context) error {
				ctx, span := otel.Tracer("httptrace/client").Start(ctx, "test")
				defer span.End()

				req, _ := http.NewRequest("GET", ts.URL, nil)
				otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

				res, err := client.Do(req)
				assert.Nil(t, err)
				return res.Body.Close()
			}(context.Background())

			assert.Nil(t, err)
		})
	}
}

func TestDontTracingSpan(t *testing.T) {
	ztrace.StartAgent(ztrace.Config{
		Name:     "go-tao-test",
		Endpoint: "http://localhost:14268/api/traces",
		Batcher:  "jaeger",
		Sampler:  1.0,
	})
	defer ztrace.StopAgent()

	DontTraceSpan("bar")
	defer notTracingSpans.Delete("bar")

	for _, test := range []string{"", "bar", "foo"} {
		t.Run(test, func(t *testing.T) {
			h := chain.New(TracingHandler("foo", test)).Then(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					span := trace.SpanFromContext(r.Context())
					spanCtx := span.SpanContext()
					if test == "bar" {
						assert.False(t, spanCtx.IsValid())
						assert.False(t, span.IsRecording())
						return
					}

					assert.True(t, span.IsRecording())
					assert.True(t, spanCtx.IsValid())
				}))
			ts := httptest.NewServer(h)
			defer ts.Close()

			client := ts.Client()
			err := func(ctx context.Context) error {
				ctx, span := otel.Tracer("httptrace/client").Start(ctx, "test")
				defer span.End()

				req, _ := http.NewRequest("GET", ts.URL, nil)
				otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

				res, err := client.Do(req)
				assert.Nil(t, err)
				return res.Body.Close()
			}(context.Background())

			assert.Nil(t, err)
		})
	}
}

func TestTraceResponseWriter(t *testing.T) {
	ztrace.StartAgent(ztrace.Config{
		Name:     "go-tao-test",
		Endpoint: "http://localhost:14268/api/traces",
		Batcher:  "jaeger",
		Sampler:  1.0,
	})
	defer ztrace.StopAgent()

	for _, test := range []int{0, 200, 300, 400, 401, 500, 503} {
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			h := chain.New(TracingHandler("foo", "bar")).Then(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					span := trace.SpanFromContext(r.Context())
					spanCtx := span.SpanContext()
					assert.True(t, span.IsRecording())
					assert.True(t, spanCtx.IsValid())
					if test != 0 {
						w.WriteHeader(test)
					}
					w.Write([]byte("hello"))
				}))
			ts := httptest.NewServer(h)
			defer ts.Close()

			client := ts.Client()
			err := func(ctx context.Context) error {
				ctx, span := otel.Tracer("httptrace/client").Start(ctx, "test")
				defer span.End()

				req, _ := http.NewRequest("GET", ts.URL, nil)
				otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

				res, err := client.Do(req)
				assert.Nil(t, err)
				resBody := make([]byte, 5)
				_, err = res.Body.Read(resBody)
				assert.Equal(t, io.EOF, err)
				assert.Equal(t, []byte("hello"), resBody, "response body fail")
				return res.Body.Close()
			}(context.Background())

			assert.Nil(t, err)
		})
	}
}
