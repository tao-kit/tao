package rpc

import (
	"github.com/sllt/tao/tools/taoctl/rpc/cli"
	"github.com/spf13/cobra"
)

var (
	// Cmd describes a rpc command.
	Cmd = &cobra.Command{
		Use:   "rpc",
		Short: "Generate rpc code",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.RPCTemplate(true)
		},
	}

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Generate rpc demo service",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:  cli.RPCNew,
	}

	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Generate proto template",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.RPCTemplate(false)
		},
	}

	protocCmd = &cobra.Command{
		Use:     "protoc",
		Short:   "Generate grpc code",
		Example: "taoctl rpc protoc xx.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.",
		Args:    cobra.ExactValidArgs(1),
		RunE:    cli.ZRPC,
	}
)

func init() {
	Cmd.Flags().StringVar(&cli.VarStringOutput, "o", "", "Output a sample proto file")
	Cmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The taoctl home path of "+
		"the template, --home and --remote cannot be set at the same time, if they are, --remote has"+
		" higher priority")
	Cmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote git repo"+
		" of the template, --home and --remote cannot be set at the same time, if they are, --remote"+
		" has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/sllt/tao-template directory structure")
	Cmd.Flags().StringVar(&cli.VarStringBranch, "branch", "", "The branch of the "+
		"remote repo, it does work with --remote")

	newCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOpt, "go_opt", nil, "")
	newCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOpt, "go-grpc_opt", nil, "")
	newCmd.Flags().StringVar(&cli.VarStringStyle, "style", "gotao", "The file "+
		"naming format, see [https://github.com/sllt/tao/tree/master/tools/taoctl/config/readme.md]")
	newCmd.Flags().BoolVar(&cli.VarBoolIdea, "idea", false, "Whether the command "+
		"execution environment is from idea plugin.")
	newCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The taoctl home path "+
		"of the template, --home and --remote cannot be set at the same time, if they are, --remote "+
		"has higher priority")
	newCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote git "+
		"repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/sllt/tao-template directory structure")
	newCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "",
		"The branch of the remote repo, it does work with --remote")
	newCmd.Flags().BoolVarP(&cli.VarBoolVerbose, "verbose", "v", false, "Enable log output")
	newCmd.Flags().MarkHidden("go_opt")
	newCmd.Flags().MarkHidden("go-grpc_opt")

	protocCmd.Flags().BoolVarP(&cli.VarBoolMultiple, "multiple", "m", false,
		"Generated in multiple rpc service mode")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOut, "go_out", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOut, "go-grpc_out", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoOpt, "go_opt", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSliceGoGRPCOpt, "go-grpc_opt", nil, "")
	protocCmd.Flags().StringSliceVar(&cli.VarStringSlicePlugin, "plugin", nil, "")
	protocCmd.Flags().StringSliceVarP(&cli.VarStringSliceProtoPath, "proto_path", "I", nil, "")
	protocCmd.Flags().StringVar(&cli.VarStringZRPCOut, "zrpc_out", "", "The zrpc output directory")
	protocCmd.Flags().StringVar(&cli.VarStringStyle, "style", "gotao", "The file "+
		"naming format, see [https://github.com/sllt/tao/tree/master/tools/taoctl/config/readme.md]")
	protocCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The taoctl home "+
		"path of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority")
	protocCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote "+
		"git repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/sllt/tao-template directory structure")
	protocCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "",
		"The branch of the remote repo, it does work with --remote")
	protocCmd.Flags().BoolVarP(&cli.VarBoolVerbose, "verbose", "v", false, "Enable log output")
	protocCmd.Flags().MarkHidden("go_out")
	protocCmd.Flags().MarkHidden("go-grpc_out")
	protocCmd.Flags().MarkHidden("go_opt")
	protocCmd.Flags().MarkHidden("go-grpc_opt")
	protocCmd.Flags().MarkHidden("plugin")
	protocCmd.Flags().MarkHidden("proto_path")

	templateCmd.Flags().StringVar(&cli.VarStringOutput, "o", "", "Output a sample proto file")
	templateCmd.Flags().StringVar(&cli.VarStringHome, "home", "", "The taoctl home"+
		" path of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority")
	templateCmd.Flags().StringVar(&cli.VarStringRemote, "remote", "", "The remote "+
		"git repo of the template, --home and --remote cannot be set at the same time, if they are, "+
		"--remote has higher priority\n\tThe git repo directory must be consistent with the "+
		"https://github.com/sllt/tao-template directory structure")
	templateCmd.Flags().StringVar(&cli.VarStringBranch, "branch", "", "The branch"+
		" of the remote repo, it does work with --remote")

	Cmd.AddCommand(newCmd)
	Cmd.AddCommand(protocCmd)
	Cmd.AddCommand(templateCmd)
}
