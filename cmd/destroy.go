package cmd

import (
	"context"
	"log"

	"github.com/prodingerd/nessus-on-demand/core"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy [flags] DEPLOYMENT [DEPLOYMENT...]",
	Short: "Destroy a deployment",
	Args:  cobra.MinimumNArgs(1),
	Run:   runDestroy,
}

func runDestroy(cmd *cobra.Command, args []string) {
	core.StartSpinner("Initializing Terraform")
	tf := core.GetTerraformInstance()
	core.StopSpinner("Terraform initialized")

	for _, deploymentId := range args {
		if err := tf.WorkspaceSelect(context.Background(), deploymentId); err != nil {
			// TODO This should not error but maybe raise a warning.
			log.Fatalf("error selecting Terraform workspace: %s", err)
		}

		if err := tf.Destroy(context.Background()); err != nil {
			log.Fatalf("error destroying Terraform deployment: %s", err)
		}

		if err := tf.WorkspaceDelete(context.Background(), deploymentId); err != nil {
			log.Fatalf("error deleting Terraform workspace: %s", err)
		}
	}
}

func init() {
	deploymentCmd.AddCommand(destroyCmd)
}
