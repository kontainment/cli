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
// - Add ability to reference a particular image
// - Add ability to reference a kontainment.yaml
// -- Both locally and remote
// - Update UX and use Charm for pretty terminal output

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create will create a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func create() {
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

	err := kCli.CreateWorkspace(context.Background(), wkspc)
	if err != nil {
		log.Fatalf("creating workspace: %v", err)
	}
}
