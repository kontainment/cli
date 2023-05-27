package cmd

import "github.com/spf13/cobra"

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "workspace is a subcommand to manage workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	workspaceCmd.AddCommand(createCmd)
	workspaceCmd.AddCommand(removeCmd)
	workspaceCmd.AddCommand(listCmd)
	workspaceCmd.AddCommand(execCmd)
}
