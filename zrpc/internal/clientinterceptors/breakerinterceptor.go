package clientinterceptors

import (
	"context"
	"path"

	"github.com/tao-kit/tao/core/breaker"
	"github.com/tao-kit/tao/zrpc/internal/codes"
	"google.golang.org/grpc"
)

// BreakerInterceptor is an interceptor that acts as a circuit breaker.
func BreakerInterceptor(ctx context.Context, method string, req, reply any,
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	breakerName := path.Join(cc.Target(), method)
	return breaker.DoWithAcceptableCtx(ctx, breakerName, func() error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}, codes.Acceptable)
}
