package upgrade

import "github.com/spf13/cobra"

// Cmd describes an upgrade command.
var Cmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade taoctl to latest version",
	RunE:  upgrade,
}
