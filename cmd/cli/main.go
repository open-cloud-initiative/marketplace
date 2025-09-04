package main

import (
	"context"

	"github.com/spf13/cobra"
)

func Init() error {
	ctx := context.Background()

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	err := RootCmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli",
}
