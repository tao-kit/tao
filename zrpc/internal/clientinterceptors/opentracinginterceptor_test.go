package clientinterceptors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"manlu.org/tao/core/trace/opentelemetry"
	"google.golang.org/grpc"
)

func TestOpenTracingInterceptor(t *testing.T) {
	opentelemetry.StartAgent(opentelemetry.Config{
		Name:     "go-zero-test",
		Endpoint: "http://localhost:14268/api/traces",
		Batcher:  "jaeger",
		Sampler:  1.0,
	})

	cc := new(grpc.ClientConn)
	err := UnaryOpenTracingInterceptor()(context.Background(), "/ListUser", nil, nil, cc,
		func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
			opts ...grpc.CallOption) error {
			return nil
		})
	assert.Nil(t, err)
}
