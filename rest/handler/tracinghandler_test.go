package handler

import (
	"manlu.org/tao/core/stringx"
	"manlu.org/tao/core/trace"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"manlu.org/tao/core/trace/tracespec"
)

func TestTracingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)

	traceId := stringx.RandId()
	req.Header.Set(trace.TraceIdKey, traceId)

	handler := TracingHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ok := r.Context().Value(tracespec.TracingKey).(tracespec.Trace)
		assert.True(t, ok)
		assert.Equal(t, traceId, span.TraceId())
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, traceId, resp.Header().Get(trace.TraceIdKey))
}
