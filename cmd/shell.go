package cmd

import (
	"github.com/urfave/cli/v2"
)

var ShellCommand = &cli.Command{
	Name:  "shell",
	Usage: "open shell in Docker container",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "Docker `IMAGE`, e.g. 'busybox:1'",
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "configuration `NAME` defined in '.denv.yml'",
		},
	},
	Action: shell,
}

func shell(c *cli.Context) error {
	cfg, err := loadRunConfig(c)
	if err != nil {
		return err
	}

	return execDockerRun(cfg.Image, []string{"/bin/sh"})
}
