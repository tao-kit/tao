package mongo

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	"manlu.org/tao/tools/taoctl/config"
	"manlu.org/tao/tools/taoctl/model/mongo/generate"
	file "manlu.org/tao/tools/taoctl/util"
)

// Action provides the entry for taoctl mongo code generation.
func Action(ctx *cli.Context) error {
	tp := ctx.StringSlice("type")
	c := ctx.Bool("cache")
	o := strings.TrimSpace(ctx.String("dir"))
	s := ctx.String("style")
	home := ctx.String("home")

	if len(home) > 0 {
		file.RegisterTaoctlHome(home)
	}

	if len(tp) == 0 {
		return errors.New("missing type")
	}

	cfg, err := config.NewConfig(s)
	if err != nil {
		return err
	}

	a, err := filepath.Abs(o)
	if err != nil {
		return err
	}

	return generate.Do(&generate.Context{
		Types:  tp,
		Cache:  c,
		Output: a,
		Cfg:    cfg,
	})
}
