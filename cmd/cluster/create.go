package cluster

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sijoma/camundactl/pkg/console"
)

var stage string

var RootCmd = &cobra.Command{
	Use:   "cluster",
	Short: "CRUD cluster commands",
}

const defaultStage = "prod"

func init() {
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.PersistentFlags().StringVar(&stage, "stage", defaultStage, "the console stage to be used, either 'dev'. 'int' or 'prod'")

	createCmd.Flags().StringVar(&planType, "type", "Trial Cluster", "the cluster type for example `Trial Cluster`")
	createCmd.Flags().StringVar(&channel, "channel", "Alpha", "the channel of the cluster for example `Alpha`")
	createCmd.Flags().StringVar(&generation, "gen", "Camunda 8.3.1", "the cluster type for example `Camunda 8.3.1`")
	createCmd.Flags().StringVar(&region, "region", "europe-west1", "the cluster type for example `europe-west1`")
	createCmd.Flags().StringVar(&stageLabel, "stagelabel", "dev", "the stage label, one of `dev`, `test`, `stage`, `prod`")
	createCmd.Flags().BoolVar(&autoUpdate, "auto", true, "whether auto updates are active, defaults to `true`")
}

var (
	planType   string
	channel    string
	generation string
	region     string
	stageLabel string
	autoUpdate bool
)

var createCmd = &cobra.Command{
	Use:   "create [cluster name]",
	Short: "Create a Camunda SaaS cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("create called on stage: %s\n", stage)
		c := console.NewConsole(cmd.Context(), stage)
		if !c.IsLoggedIn() {
			fmt.Println("run camundactl login first OR put your accessToken in the config file")
			return
		}

		clusterName := args[0]
		cluster := console.NamedClusterCreateRequest{
			Name:       clusterName,
			PlanType:   planType,
			Channel:    channel,
			Generation: generation,
			Region:     region,
			AutoUpdate: autoUpdate,
			StageLabel: stageLabel,
		}
		resp, err := c.CreateCluster(
			cmd.Context(), cluster,
		)
		if err != nil {
			fmt.Printf("❌: %s\n", err)
		} else {
			fmt.Printf("✅: Cluster created with ID %s\n", resp)
		}
	},
}
