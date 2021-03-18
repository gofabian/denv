package main

import (
	"fmt"
	"os"

	"github.com/gofabian/denv/run"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "denv",
		Usage: "Use working directory in docker image",
		Commands: []*cli.Command{
			run.Command,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
