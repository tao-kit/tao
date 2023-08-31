package testdata

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/sllt/tao/tools/taoctl/util/pathx"
)

type (
	File struct {
		IsDir        bool
		Path         string
		AbsolutePath string
		Content      string
		Cmd          string
	}

	Files []File
)

func (f File) execute(taoctl string) error {
	printDir := f.Path
	dir := f.AbsolutePath
	if !f.IsDir {
		printDir = filepath.Dir(printDir)
		dir = filepath.Dir(dir)
	}
	printCommand := strings.ReplaceAll(fmt.Sprintf("cd %s && %s", printDir, f.Cmd), "taoctl", filepath.Base(taoctl))
	command := strings.ReplaceAll(fmt.Sprintf("cd %s && %s", dir, f.Cmd), "taoctl", taoctl)
	fmt.Println(color.LightGreen.Render(printCommand))
	cmd := exec.Command("sh", "-c", command)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (fs Files) execute(taoctl string) error {
	for _, f := range fs {
		err := f.execute(taoctl)
		if err != nil {
			return err
		}
	}
	return nil
}

func mustGetTestData(baseDir string) (Files, Files) {
	if len(baseDir) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		baseDir = dir
	}
	baseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, nil
	}
	createFile := func(baseDir string, data File) (File, error) {
		fp := filepath.Join(baseDir, data.Path)
		dir := filepath.Dir(fp)
		if data.IsDir {
			dir = fp
		}
		if err := pathx.MkdirIfNotExist(dir); err != nil {
			return data, err
		}
		data.AbsolutePath = fp
		if data.IsDir {
			return data, nil
		}

		return data, os.WriteFile(fp, []byte(data.Content), os.ModePerm)
	}
	oldDir := filepath.Join(baseDir, "old_fs")
	newDir := filepath.Join(baseDir, "new_fs")
	os.RemoveAll(oldDir)
	os.RemoveAll(newDir)
	var oldFiles, newFiles []File
	for _, data := range list {
		od, err := createFile(oldDir, data)
		if err != nil {
			log.Fatalln(err)
		}
		oldFiles = append(oldFiles, od)
		nd, err := createFile(newDir, data)
		if err != nil {
			log.Fatalln(err)
		}
		newFiles = append(newFiles, nd)
	}
	return oldFiles, newFiles
}

func MustRun(baseDir string) {
	oldFiles, newFiles := mustGetTestData(baseDir)
	taoctlOld, err := exec.LookPath("taoctl.old")
	must(err)
	taoctlNew, err := exec.LookPath("taoctl")
	must(err)
	fmt.Println(color.LightBlue.Render("========================taoctl.old======================="))
	must(oldFiles.execute(taoctlOld))
	fmt.Println()
	fmt.Println(color.LightBlue.Render("========================taoctl.new======================="))
	must(newFiles.execute(taoctlNew))
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
