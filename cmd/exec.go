package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// TODO:
// - Implement this

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec will instantiate a terminal for a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		exec()
	},
}

func exec() {
	log.Fatal("not implemented")
}
