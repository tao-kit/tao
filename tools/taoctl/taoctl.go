package main

import (
	"manlu.org/tao/core/load"
	"manlu.org/tao/core/logx"
	"manlu.org/tao/tools/taoctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
