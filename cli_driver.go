package cli

import "github.com/spf13/cobra"

type cliDriver struct {
	rootCmd *cobra.Command
}

// AddCommand add new command to cli
func (cli *cliDriver) AddCommand(cmd *cobra.Command) {
	cli.rootCmd.AddCommand(cmd)
}

// RootCommand return pointer to cli root command
func (cli *cliDriver) RootCommand() *cobra.Command {
	return cli.rootCmd
}

// Run cli
func (cli *cliDriver) Run() {
	cli.rootCmd.Execute()
}
