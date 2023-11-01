package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version info",
	Long:  `This command prints the commit hash and version info of the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		myFigure := figure.NewFigure("CAMUNDACTL", "isometric1", true)
		myFigure.Print()

		fmt.Println("Version: ", version)
		fmt.Println("Commit: ", commit)
		fmt.Println("Build Time: ", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
