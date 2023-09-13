package cli

import (
	"github.com/spf13/cobra"
)

// NewCLI create new cli app
func NewCLI(name string, desc string) CLI {
	cli := new(cliDriver)
	cli.rootCmd = new(cobra.Command)
	cli.rootCmd.Use = name
	cli.rootCmd.Short = desc
	return cli
}
