package generator

import (
	_ "embed"
	"path/filepath"
	"strings"

	"manlu.org/tao/tools/taoctl/util"
	"manlu.org/tao/tools/taoctl/util/pathx"
	"manlu.org/tao/tools/taoctl/util/stringx"
)

//go:embed rpc.tpl
var rpcTemplateText string

// ProtoTmpl returns a sample of a proto file
func ProtoTmpl(out string, validate bool) error {
	protoFilename := filepath.Base(out)
	serviceName := stringx.From(strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename)))
	text, err := pathx.LoadTemplate(category, rpcTemplateFile, rpcTemplateText)
	if err != nil {
		return err
	}

	dir := filepath.Dir(out)
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = util.With("t").Parse(text).SaveTo(map[string]string{
		"package":     serviceName.Untitle(),
		"serviceName": serviceName.Title(),
	}, out, false)

	if validate {
		_ = ProtoValidateImpl()
	}

	return err
}
