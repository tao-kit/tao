package ktgen

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/sllt/tao/tools/taoctl/api/spec"
)

var (
	//go:embed apibase.tpl
	apiBaseTemplate string
	//go:embed api.tpl
	apiTemplate string
)

func genBase(dir, pkg string, api *spec.ApiSpec) error {
	e := os.MkdirAll(dir, 0o755)
	if e != nil {
		return e
	}
	path := filepath.Join(dir, "BaseApi.kt")
	if _, e := os.Stat(path); e == nil {
		fmt.Println("BaseApi.kt already exists, skipped it.")
		return nil
	}

	file, e := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if e != nil {
		return e
	}
	defer file.Close()

	t, e := template.New("n").Parse(apiBaseTemplate)
	if e != nil {
		return e
	}
	e = t.Execute(file, pkg)
	if e != nil {
		return e
	}
	return nil
}

func genApi(dir, pkg string, api *spec.ApiSpec) error {
	properties := api.Info.Properties
	fmt.Println(properties)
	if properties == nil {
		return fmt.Errorf("none properties")
	}

	title := properties["title"]
	if len(title) == 0 {
		return fmt.Errorf("none title")
	}

	desc := properties["desc"]
	if len(desc) == 0 {
		return fmt.Errorf("none desc")
	}

	name := strcase.ToCamel(title + "Api")
	path := filepath.Join(dir, name+".kt")
	api.Info.Title = name
	api.Info.Desc = desc

	e := os.MkdirAll(dir, 0o755)
	if e != nil {
		return e
	}

	file, e := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if e != nil {
		return e
	}
	defer file.Close()

	t, e := template.New("api").Funcs(funcsMap).Parse(apiTemplate)
	if e != nil {
		return e
	}
	type data struct {
		*spec.ApiSpec
		Pkg string
	}
	return t.Execute(file, data{api, pkg})
}
