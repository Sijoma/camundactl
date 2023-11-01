package cluster

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sijoma/camundactl/pkg/console"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [cluster id]",
	Short: "Delete a Camunda SaaS cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a id argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("delete called on stage: %s\n", stage)
		c := console.NewConsole(stage)
		if c.AccessToken.AccessToken == "" {
			fmt.Println("run camundactl login first OR put your accessToken in the config file")
			return
		}

		clusterID := args[0]
		err := c.DeleteCluster(
			cmd.Context(), clusterID,
		)
		if err != nil {
			fmt.Printf("❌: Failed to delete cluster with ID %s; %v\n", clusterID, err)
		} else {
			fmt.Printf("✅: Cluster deleted with ID %s\n", clusterID)
		}
	},
}
