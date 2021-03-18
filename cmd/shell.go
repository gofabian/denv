package cmd

import (
	"github.com/urfave/cli/v2"
)

var ShellCommand = &cli.Command{
	Name:  "shell",
	Usage: "open shell in Docker container",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "image",
			Aliases:  []string{"i"},
			Usage:    "Docker `IMAGE`, e.g. 'busybox:1'",
			Required: true,
		},
	},
	Action: shell,
}

func shell(c *cli.Context) error {
	image := c.String("image")
	return execDockerRun(image, []string{"/bin/sh"})
}
