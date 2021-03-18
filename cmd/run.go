package cmd

import (
	"github.com/urfave/cli/v2"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "run command in Docker container",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "image",
			Aliases:  []string{"i"},
			Usage:    "Docker `IMAGE`, e.g. 'busybox:1'",
			Required: true,
		},
	},
	Action: run,
}

func run(c *cli.Context) error {
	image := c.String("image")
	args := c.Args().Slice()
	return execDockerRun(image, args)
}
