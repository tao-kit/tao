package main

import (
	"github.com/sllt/tao/core/load"
	"github.com/sllt/tao/core/logx"
	"github.com/sllt/tao/tools/taoctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
