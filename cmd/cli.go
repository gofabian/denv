package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func RunCliApp() {
	app := &cli.App{
		Name:  "denv",
		Usage: "Use working directory in docker image",
		Commands: []*cli.Command{
			RunCommand,
			ShellCommand,
			ExecCommand,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "`PATH` to configuration file, e.g. 'path/to/config.denv.yml'",
			},
		},
		Action: runShell, // fallback to shell command
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
