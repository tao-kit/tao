package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	conf "manlu.org/tao/tools/taoctl/config"
	"manlu.org/tao/tools/taoctl/rpc/parser"
	"manlu.org/tao/tools/taoctl/util"
	"manlu.org/tao/tools/taoctl/util/format"
	"manlu.org/tao/tools/taoctl/util/stringx"
)

const mainTemplate = `package main

import (
	"flag"
	"fmt"

	{{.imports}}

	"manlu.org/tao/core/conf"
	"manlu.org/tao/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.New{{.serviceNew}}Server(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		{{.pkg}}.Register{{.service}}Server(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
`

// GenMain generates the main file of the rpc service, which is an rpc service program call entry
func (g *DefaultGenerator) GenMain(ctx DirContext, proto parser.Proto, cfg *conf.Config) error {
	mainFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(ctx.GetMain().Filename, fmt.Sprintf("%v.go", mainFilename))
	imports := make([]string, 0)
	pbImport := fmt.Sprintf(`"%v"`, ctx.GetPb().Package)
	svcImport := fmt.Sprintf(`"%v"`, ctx.GetSvc().Package)
	remoteImport := fmt.Sprintf(`"%v"`, ctx.GetServer().Package)
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	imports = append(imports, configImport, pbImport, remoteImport, svcImport)
	text, err := util.LoadTemplate(category, mainTemplateFile, mainTemplate)
	if err != nil {
		return err
	}

	etcFileName, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	return util.With("main").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"serviceName": etcFileName,
		"imports":     strings.Join(imports, util.NL),
		"pkg":         proto.PbPackage,
		"serviceNew":  stringx.From(proto.Service.Name).ToCamel(),
		"service":     parser.CamelCase(proto.Service.Name),
	}, fileName, false)
}