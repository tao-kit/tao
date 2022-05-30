package upgrade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"manlu.org/tao/tools/taoctl/rpc/execx"
)

// Upgrade gets the latest taoctl by
// go install manlu.org/tao/tools/taoctl@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install manlu.org/tao/tools/taoctl@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install manlu.org/tao/tools/taoctl@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
