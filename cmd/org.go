package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sijoma/camundactl/pkg/console"
)

var orgCmd = &cobra.Command{
	Use:   "org [ID of organization]",
	Short: "Set your current Camunda org",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a org id argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updating active Organization")
		c := console.NewConsole(stage)
		if c.AccessToken.AccessToken == "" {
			fmt.Println("run camundactl login first OR put your accessToken in the config file")
			return
		}

		setOrgID := args[0]
		err := c.SetOrg(setOrgID)
		if err != nil {
			fmt.Printf("Unable to set Org with ID: %s", setOrgID)
			return
		}
		fmt.Printf("Active Org to Name: %s | ID: %s\n", c.ActiveOrg.Name, c.ActiveOrg.Uuid)
		return
	},
}

var orgListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Camunda orgs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("org list called")
		c := console.NewConsole(stage)
		if c.AccessToken.AccessToken == "" {
			fmt.Println("run camundactl login first OR put your accessToken in the config file")
			return
		}

		c.PrintOrgs()
		fmt.Printf("Active Org to Name: %s | ID: %s\n", c.ActiveOrg.Name, c.ActiveOrg.Uuid)
		return
	},
}

func init() {
	orgCmd.AddCommand(orgListCmd)
	rootCmd.AddCommand(orgCmd)
}
