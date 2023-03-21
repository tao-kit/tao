package migrate

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sllt/tao/core/stringx"
	"github.com/sllt/tao/tools/taoctl/rpc/execx"
	"github.com/sllt/tao/tools/taoctl/util/console"
	"github.com/sllt/tao/tools/taoctl/util/ctx"
)

const (
	deprecatedGoTaoMod  = "github.com/sllt/tao"
	deprecatedBuilderx  = "github.com/sllt/tao/tools/taoctl/model/sql/builderx"
	replacementBuilderx = "github.com/sllt/tao/core/stores/builder"
	goTaoMod            = "github.com/sllt/tao"
)

var errInvalidGoMod = errors.New("it's only working for go module")

func editMod(version string, verbose bool) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	isGoMod, _ := ctx.IsGoMod(wd)
	if !isGoMod {
		return nil
	}

	latest, err := getLatest(goTaoMod, verbose)
	if err != nil {
		return err
	}

	if !stringx.Contains(latest, version) {
		return fmt.Errorf("release version %q is not found", version)
	}

	mod := fmt.Sprintf("%s@%s", goTaoMod, version)
	err = removeRequire(deprecatedGoTaoMod, verbose)
	if err != nil {
		return err
	}

	return addRequire(mod, verbose)
}

func addRequire(mod string, verbose bool) error {
	if verbose {
		console.Info("adding require %s ...", mod)
		time.Sleep(200 * time.Millisecond)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	isGoMod, _ := ctx.IsGoMod(wd)
	if !isGoMod {
		return errInvalidGoMod
	}

	_, err = execx.Run("go mod edit -require "+mod, wd)
	return err
}

func removeRequire(mod string, verbose bool) error {
	if verbose {
		console.Info("remove require %s ...", mod)
		time.Sleep(200 * time.Millisecond)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = execx.Run("go mod edit -droprequire "+mod, wd)
	return err
}

func tidy(verbose bool) error {
	if verbose {
		console.Info("go mod tidy ...")
		time.Sleep(200 * time.Millisecond)
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	isGoMod, _ := ctx.IsGoMod(wd)
	if !isGoMod {
		return nil
	}

	_, err = execx.Run("go mod tidy", wd)
	return err
}
