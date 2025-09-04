package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(CatalogCmd)

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli",
}
