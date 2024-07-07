package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var releases = `cbcli v0.0.0 (yyyy-mm-dd):
  - Initial release
`

func init() {
	rootCmd.AddCommand(releasesCmd)
}

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Prints the release information",
	Long:  `Prints the release information`,
	Run: func(cmd *cobra.Command, args []string) {
		// print release information
		printVersion()
		fmt.Println(releases)
	},
}
