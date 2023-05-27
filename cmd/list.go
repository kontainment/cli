package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/kontainment/engine/client"
	"github.com/spf13/cobra"
)

// TODO:
// - Update UX and use Charm for pretty terminal output

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list will list all workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func list() {
	kCli := client.NewClient()

	wkspcs, err := kCli.ListWorkspaces(context.Background())
	if err != nil {
		log.Fatalf("listing workspaces: %v", err)
	}

	for _, wkspc := range wkspcs.Workspaces {
		fmt.Println(wkspc.Name)
	}
}
