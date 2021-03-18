package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/gofabian/denv/cmd"
)

func main() {
	app := &cli.App{
		Name:  "denv",
		Usage: "Use working directory in docker image",
		Commands: []*cli.Command{
			cmd.RunCommand,
			cmd.ShellCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
