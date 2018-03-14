package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snk",
	Short: "Sync n kill",
	Long: `This software will give you the utility
To sync your device to remote and clear the current one`,
}

// Execute is the root executor
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
