package cmd

import (
	"github.com/sllt/tao/tools/taoctl/compare/testdata"
	"github.com/sllt/tao/tools/taoctl/util/console"
	"github.com/spf13/cobra"
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
