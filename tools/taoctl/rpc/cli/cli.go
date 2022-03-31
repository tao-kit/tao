package cli

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli"
	"manlu.org/tao/tools/taoctl/rpc/generator"
)

// RPCNew is to generate rpc greet service, this greet service can speed
// up your understanding of the zrpc service structure
func RPCNew(c *cli.Context) error {
	rpcname := c.Args().First()
	ext := filepath.Ext(rpcname)
	if len(ext) > 0 {
		return fmt.Errorf("unexpected ext: %s", ext)
	}
	style := c.String("style")
	verbose := c.Bool("verbose")
	validate := c.Bool("validate")

	protoName := rpcname + ".proto"
	filename := filepath.Join(".", rpcname, protoName)
	src, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	err = generator.ProtoTmpl(src, validate)
	if err != nil {
		return err
	}

	var ctx generator.ZRpcContext
	ctx.Src = src
	ctx.GoOutput = filepath.Dir(src)
	ctx.GrpcOutput = filepath.Dir(src)
	ctx.IsGooglePlugin = true
	ctx.Output = filepath.Dir(src)
	ctx.ProtocCmd = fmt.Sprintf("protoc -I=%s %s --go_out=%s --go-grpc_out=%s", filepath.Dir(src), filepath.Base(src), filepath.Dir(src), filepath.Dir(src))
	g := generator.NewGenerator(style, verbose)
	return g.Generate(&ctx)
}

// RPCTemplate is the entry for generate rpc template
func RPCTemplate(c *cli.Context) error {
	protoFile := c.String("o")
	validate := c.Bool("validate")

	if len(protoFile) == 0 {
		return errors.New("missing -o")
	}

	return generator.ProtoTmpl(protoFile, validate)
}
