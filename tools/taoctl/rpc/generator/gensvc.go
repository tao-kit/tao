package generator

import (
	"fmt"
	"path/filepath"

	conf "manlu.org/tao/tools/taoctl/config"
	"manlu.org/tao/tools/taoctl/rpc/parser"
	"manlu.org/tao/tools/taoctl/util"
	"manlu.org/tao/tools/taoctl/util/format"
	"manlu.org/tao/tools/taoctl/util/pathx"
)

const svcTemplate = `package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:c,
	}
}
`

// GenSvc generates the servicecontext.go file, which is the resource dependency of a service,
// such as rpc dependency, model dependency, etc.
func (g *DefaultGenerator) GenSvc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetSvc()
	svcFilename, err := format.FileNamingFormat(cfg.NamingFormat, "service_context")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, svcFilename+".go")
	text, err := pathx.LoadTemplate(category, svcTemplateFile, svcTemplate)
	if err != nil {
		return err
	}

	return util.With("svc").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"imports": fmt.Sprintf(`"%v"`, ctx.GetConfig().Package),
	}, fileName, false)
}
