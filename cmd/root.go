package cmd

import (
	"github.com/spf13/cobra"

	"github.com/sijoma/camundactl/cmd/cluster"
	"github.com/sijoma/camundactl/internal/config"
)

var (
	cfgFile     string
	accessToken string
	version     = "dev"
	commit      = "none"
	date        = "unknown"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "camundactl",
	Short: "camunda ctl allows to provison Camunda SaaS resources",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Command with subcommands
	rootCmd.AddCommand(cluster.RootCmd)

	// Not sure if that default makes sense, the usage is then also not correct
	// However if the default is set, it seamlessly creates the file if it does not exist
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.camundactl.yaml)")
	rootCmd.PersistentFlags().StringVar(&accessToken, "accessToken", "", "console access token")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.CreateConfig(cfgFile)
}
