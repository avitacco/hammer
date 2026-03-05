package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "hammer",
	Short: "A tool for building and publishing Puppet modules",
}

func Execute() error {
	return rootCmd.Execute()
}
