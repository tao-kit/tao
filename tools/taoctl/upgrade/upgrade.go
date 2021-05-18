package upgrade

import (
	"fmt"

	"manlu.org/tao/tools/taoctl/rpc/execx"
	"github.com/urfave/cli"
)

// Upgrade gets the latest taoctl by
// go get -u manlu.org/tao/tools/taoctl
func Upgrade(_ *cli.Context) error {
	info, err := execx.Run("GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u manlu.org/tao/tools/taoctl", "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
