package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const art string = `<INSERT ASCII ART HERE>`

var version string = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion() {
	fmt.Print(art)
	fmt.Printf("Version: %s\n\n", version)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  `Prints the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}
