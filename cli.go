package cli

import "github.com/spf13/cobra"

// CLI interface
type CLI interface {
	// AddCommand add new command to cli
	AddCommand(cmd *cobra.Command)
	// RootCommand return pointer to cli root command
	RootCommand() *cobra.Command
	// Run cli
	Run()
}
