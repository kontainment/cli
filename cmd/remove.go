package cmd

import (
	"context"
	"log"

	"github.com/kontainment/engine/api/types"
	"github.com/kontainment/engine/client"
	"github.com/kontainment/engine/containertools"
	"github.com/spf13/cobra"
)

// TODO:
// - Add ability to specify the workspace to remove
// - Update UX and use Charm for pretty terminal output

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove will remove a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		remove()
	},
}

func remove() {
	// TODO: for now hardcode the workspace to run
	wkspc := types.NewWorkspace(
		types.WithImage(
			containertools.NewImage(
				containertools.WithRepository("bpalmer/cade-example"),
				containertools.WithTag("latest")),
		),
		types.WithName("testspace"),
	)

	kCli := client.NewClient()

	err := kCli.DeleteWorkspace(context.Background(), wkspc)
	if err != nil {
		log.Fatalf("removing workspace: %v", err)
	}
}
