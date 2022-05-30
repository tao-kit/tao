package cmd

import (
	"github.com/spf13/cobra"
	"manlu.org/tao/tools/taoctl/compare/testdata"
	"manlu.org/tao/tools/taoctl/util/console"
)

var rootCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare the taoctl commands generated results between urfave and cobra",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		testdata.MustRun(dir)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		console.Error("%+v", err)
	}
}
