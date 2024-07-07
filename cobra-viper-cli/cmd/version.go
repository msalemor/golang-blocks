package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ASCII ART @ https://www.asciiart.eu/text-to-ascii-art
const art string = `
                               ████   ███ 
                              ░░███  ░░░  
  ██████  █████ █████  ██████  ░███  ████ 
 ███░░███░░███ ░░███  ███░░███ ░███ ░░███ 
░███ ░░░  ░███  ░███ ░███ ░░░  ░███  ░███ 
░███  ███ ░░███ ███  ░███  ███ ░███  ░███ 
░░██████   ░░█████   ░░██████  █████ █████
 ░░░░░░     ░░░░░     ░░░░░░  ░░░░░ ░░░░░ 
`

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
