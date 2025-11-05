package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(newRootCmd().Execute())
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "integration-app",
		Short:   "Integration app for Bitrix24 and Facebook",
		Version: "0.0.1",
	}

	rootCmd.AddCommand(
		newServerCmd(),
		newMigrateCmd(),
		newHealthCmd(),
	)

	return rootCmd
}

func newServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start API server",
		RunE:  runServer,
	}
}

func newMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		RunE:  runMigrate,
	}
}

func newHealthCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Check application health",
		RunE:  runHealth,
	}
}
