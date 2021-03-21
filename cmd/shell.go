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
	Action: runShell,
}

func runShell(c *cli.Context) error {
	cfg, err := loadRunConfig(c)
	if err != nil {
		return err
	}

	dockerCmd := []string{"/bin/sh"}
	if cfg.Shell != "" {
		customCmd, err := shlex.Split(cfg.Shell)
		if err != nil {
			return err
		}
		if len(dockerCmd) > 0 {
			dockerCmd = customCmd
		}
	}

	return execDockerRun(cfg.Image, nil, dockerCmd)
}
