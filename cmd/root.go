package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(migrationCmd)
}

var rootCmd = &cobra.Command{Use: "app"}

func Execute() error {
	return rootCmd.Execute()
}
