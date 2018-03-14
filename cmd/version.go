package cmd

import (
	"fmt"

	"github.com/jonaklab/snk/version"
	"github.com/spf13/cobra"
)

var vCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of app",
	Long:  `Print the version of app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.AddCommand(vCmd)
}
