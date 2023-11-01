package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"

	"github.com/sijoma/camundactl/pkg/console"
)

var stage string
var clientID string
var clientSecret string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate to Camunda Console and store the accessToken in the configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		if clientID == "" || clientSecret == "" {
			fmt.Println("using device login")
			c := console.NewConsole(stage)
			c.Auth()

			err := c.UpdateProfile(cmd.Context())
			if err != nil {
				fmt.Printf("unable to update profile: %v\n", err)
			}

			fmt.Printf("Active Org to Name: %s | ID: %s\n", c.ActiveOrg.Name, c.ActiveOrg.Uuid)
			myFigure := figure.NewFigure("CAMUNDA", "isometric1", true)
			myFigure.Print()
			return
		}

		// Client credentials were passed we can use this
		fmt.Println("credentials detected, using machine login")
		c := console.NewMachineConsole(stage, clientID)
		err := c.MachineLogin(clientID, clientSecret)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("set active Org to Name: %s | ID: %s\n", c.ActiveOrg.Name, c.ActiveOrg.Uuid)
		return
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.PersistentFlags().StringVar(&stage, "stage", "prod", "the console stage to be used, either 'dev'. 'int' or 'prod'")
	rootCmd.PersistentFlags().StringVar(&clientID, "client_id", "", "the id of the client")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "client_secret", "", "the secret of the client")
}
