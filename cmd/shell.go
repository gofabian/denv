package cmd

import (
	"github.com/google/shlex"
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

	args := []string{"/bin/sh"}
	if cfg.Shell != "" {
		customArgs, err := shlex.Split(cfg.Shell)
		if err != nil {
			return err
		}
		if len(args) > 0 {
			args = customArgs
		}
	}

	return execDockerRun(cfg.Image, args)
}
