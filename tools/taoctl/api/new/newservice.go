package new

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli"
	"manlu.org/tao/tools/taoctl/api/gogen"
	conf "manlu.org/tao/tools/taoctl/config"
	"manlu.org/tao/tools/taoctl/util"
	"manlu.org/tao/tools/taoctl/util/pathx"
)

//go:embed api.tpl
var apiTemplate string

// CreateServiceCommand fast create service
func CreateServiceCommand(c *cli.Context) error {
	args := c.Args()
	dirName := args.First()
	if len(dirName) == 0 {
		dirName = "greet"
	}

	dirStyle := c.String("style")
	if len(dirStyle) == 0 {
		dirStyle = conf.DefaultFormat
	}
	if strings.Contains(dirName, "-") {
		return errors.New("api new command service name not support strikethrough, because this will used by function name")
	}

	abs, err := filepath.Abs(dirName)
	if err != nil {
		return err
	}

	err = pathx.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	dirName = filepath.Base(filepath.Clean(abs))
	filename := dirName + ".api"
	apiFilePath := filepath.Join(abs, filename)
	fp, err := os.Create(apiFilePath)
	if err != nil {
		return err
	}

	defer fp.Close()

	home := c.String("home")
	remote := c.String("remote")
	branch := c.String("branch")
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterTaoctlHome(home)
	}

	text, err := pathx.LoadTemplate(category, apiTemplateFile, apiTemplate)
	if err != nil {
		return err
	}

	t := template.Must(template.New("template").Parse(text))
	if err := t.Execute(fp, map[string]string{
		"name":    dirName,
		"handler": strings.Title(dirName),
	}); err != nil {
		return err
	}

	err = gogen.DoGenProject(apiFilePath, abs, dirStyle)
	return err
}
