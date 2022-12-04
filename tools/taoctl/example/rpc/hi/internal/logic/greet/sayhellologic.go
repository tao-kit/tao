package greetlogic

import (
	"context"

	"github.com/sllt/tao/tools/taoctl/example/rpc/hi/internal/svc"
	"github.com/sllt/tao/tools/taoctl/example/rpc/hi/pb/hi"

	"github.com/sllt/tao/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *hi.HelloReq) (*hi.HelloResp, error) {
	// todo: add your logic here and delete this line

	return &hi.HelloResp{}, nil
}
