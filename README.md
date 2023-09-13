# CLI

CLI app based on Cobra library.

## Create New CLI App

```go
import "github.com/gomig/cli"
app := cli.NewCLI("myApp", "My App Description")
```

## Usage

CLI interface contains following methods:

### AddCommand

Add new command to cli

```go
// Signature:
AddCommand(cmd *cobra.Command)
```

### RootCommand

Return pointer to cli root command

```go
// Signature:
RootCommand() *cobra.Command
```

### Run

Run cli

```go
// Signature:
Run()
```
