package internal

import (
	"github.com/sllt/tao/core/discov"
	"testing"

	"github.com/sllt/tao/core/netx"
	"github.com/stretchr/testify/assert"
)

func TestNewRpcPubServer(t *testing.T) {
	s, err := NewRpcPubServer(discov.EtcdConf{
		User: "user",
		Pass: "pass",
	}, "", ServerMiddlewaresConf{})
	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		s.Start(nil)
	})
}

func TestFigureOutListenOn(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "192.168.0.5:1234",
			expect: "192.168.0.5:1234",
		},
		{
			input:  "0.0.0.0:8080",
			expect: netx.InternalIp() + ":8080",
		},
		{
			input:  ":8080",
			expect: netx.InternalIp() + ":8080",
		},
		{
			input:  "",
			expect: netx.InternalIp(),
		},
	}

	for _, test := range tests {
		val := figureOutListenOn(test.input)
		assert.Equal(t, test.expect, val)
	}
}
