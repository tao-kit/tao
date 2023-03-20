package internal

import (
	"sync"
	"testing"

	"github.com/sllt/tao/core/stat"
	"github.com/sllt/tao/zrpc/internal/mock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestRpcServer(t *testing.T) {
	metrics := stat.NewMetrics("foo")
	server := NewRpcServer("localhost:54321", ServerMiddlewaresConf{
		Trace:      true,
		Recover:    true,
		Stat:       true,
		Prometheus: true,
		Breaker:    true,
	}, WithMetrics(metrics), WithRpcHealth(true))
	server.SetName("mock")
	var wg sync.WaitGroup
	var grpcServer *grpc.Server
	var lock sync.Mutex
	wg.Add(1)
	go func() {
		err := server.Start(func(server *grpc.Server) {
			lock.Lock()
			mock.RegisterDepositServiceServer(server, new(mock.DepositServer))
			grpcServer = server
			lock.Unlock()
			wg.Done()
		})
		assert.Nil(t, err)
	}()

	wg.Wait()
	lock.Lock()
	grpcServer.GracefulStop()
	lock.Unlock()
}

func TestRpcServer_WithBadAddress(t *testing.T) {
	server := NewRpcServer("localhost:111111", ServerMiddlewaresConf{
		Trace:      true,
		Recover:    true,
		Stat:       true,
		Prometheus: true,
		Breaker:    true,
	}, WithRpcHealth(true))
	server.SetName("mock")
	err := server.Start(func(server *grpc.Server) {
		mock.RegisterDepositServiceServer(server, new(mock.DepositServer))
	})
	assert.NotNil(t, err)
}
