package zrpc

import (
	"github.com/alicebob/miniredis/v2"
	"testing"
	"time"

	"github.com/sllt/tao/core/discov"
	"github.com/sllt/tao/core/logx"
	"github.com/sllt/tao/core/service"
	"github.com/sllt/tao/core/stat"
	"github.com/sllt/tao/core/stores/redis"
	"github.com/sllt/tao/zrpc/internal"
	"github.com/sllt/tao/zrpc/internal/serverinterceptors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestServer_setupInterceptors(t *testing.T) {
	rds, err := miniredis.Run()
	assert.NoError(t, err)
	defer rds.Close()

	server := new(mockedServer)
	conf := RpcServerConf{
		Auth: true,
		Redis: redis.RedisKeyConf{
			RedisConf: redis.RedisConf{
				Host: rds.Addr(),
				Type: redis.NodeType,
			},
			Key: "foo",
		},
		CpuThreshold: 10,
		Timeout:      100,
		Middlewares: ServerMiddlewaresConf{
			Trace:      true,
			Recover:    true,
			Stat:       true,
			Prometheus: true,
			Breaker:    true,
		},
	}
	err = setupInterceptors(server, conf, new(stat.Metrics))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(server.unaryInterceptors))
	assert.Equal(t, 1, len(server.streamInterceptors))

	rds.SetError("mock error")
	err = setupInterceptors(server, conf, new(stat.Metrics))
	assert.Error(t, err)
}

func TestServer(t *testing.T) {
	DontLogContentForMethod("foo")
	SetServerSlowThreshold(time.Second)
	svr := MustNewServer(RpcServerConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				ServiceName: "foo",
				Mode:        "console",
			},
		},
		ListenOn:      "localhost:8080",
		Etcd:          discov.EtcdConf{},
		Auth:          false,
		Redis:         redis.RedisKeyConf{},
		StrictControl: false,
		Timeout:       0,
		CpuThreshold:  0,
		Middlewares: ServerMiddlewaresConf{
			Trace:      true,
			Recover:    true,
			Stat:       true,
			Prometheus: true,
			Breaker:    true,
		},
	}, func(server *grpc.Server) {
	})
	svr.AddOptions(grpc.ConnectionTimeout(time.Hour))
	svr.AddUnaryInterceptors(serverinterceptors.UnaryRecoverInterceptor)
	svr.AddStreamInterceptors(serverinterceptors.StreamRecoverInterceptor)
	go svr.Start()
	svr.Stop()
}

func TestServerError(t *testing.T) {
	_, err := NewServer(RpcServerConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				ServiceName: "foo",
				Mode:        "console",
			},
		},
		ListenOn: "localhost:8080",
		Etcd: discov.EtcdConf{
			Hosts: []string{"localhost"},
		},
		Auth:  true,
		Redis: redis.RedisKeyConf{},
		Middlewares: ServerMiddlewaresConf{
			Trace:      true,
			Recover:    true,
			Stat:       true,
			Prometheus: true,
			Breaker:    true,
		},
	}, func(server *grpc.Server) {
	})
	assert.NotNil(t, err)
}

func TestServer_HasEtcd(t *testing.T) {
	svr := MustNewServer(RpcServerConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				ServiceName: "foo",
				Mode:        "console",
			},
		},
		ListenOn: "localhost:8080",
		Etcd: discov.EtcdConf{
			Hosts: []string{"notexist"},
			Key:   "any",
		},
		Redis: redis.RedisKeyConf{},
		Middlewares: ServerMiddlewaresConf{
			Trace:      true,
			Recover:    true,
			Stat:       true,
			Prometheus: true,
			Breaker:    true,
		},
	}, func(server *grpc.Server) {
	})
	svr.AddOptions(grpc.ConnectionTimeout(time.Hour))
	svr.AddUnaryInterceptors(serverinterceptors.UnaryRecoverInterceptor)
	svr.AddStreamInterceptors(serverinterceptors.StreamRecoverInterceptor)
	go svr.Start()
	svr.Stop()
}

func TestServer_StartFailed(t *testing.T) {
	svr := MustNewServer(RpcServerConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				ServiceName: "foo",
				Mode:        "console",
			},
		},
		ListenOn: "localhost:aaa",
		Middlewares: ServerMiddlewaresConf{
			Trace:      true,
			Recover:    true,
			Stat:       true,
			Prometheus: true,
			Breaker:    true,
		},
	}, func(server *grpc.Server) {
	})

	assert.Panics(t, svr.Start)
}

type mockedServer struct {
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

func (m *mockedServer) AddOptions(_ ...grpc.ServerOption) {
}

func (m *mockedServer) AddStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) {
	m.streamInterceptors = append(m.streamInterceptors, interceptors...)
}

func (m *mockedServer) AddUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) {
	m.unaryInterceptors = append(m.unaryInterceptors, interceptors...)
}

func (m *mockedServer) SetName(_ string) {
}

func (m *mockedServer) Start(_ internal.RegisterFn) error {
	return nil
}
