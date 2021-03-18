package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/gofabian/denv/cfg"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "run command in Docker container",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "Docker `IMAGE`, e.g. 'busybox:1'",
		},
	},
	Action: run,
}

func run(c *cli.Context) error {
	cfg, err := loadRunConfig(c)
	if err != nil {
		return err
	}

	args := c.Args().Slice()
	return execDockerRun(cfg.Image, args)
}

func loadRunConfig(c *cli.Context) (*cfg.DenvConfig, error) {
	cfg, err := cfg.ReadConfigFromFile()
	if err != nil {
		return nil, err
	}

	if c.IsSet("image") {
		cfg.Image = c.String("image")
	}
	if cfg.Image == "" {
		return nil, fmt.Errorf("missing image")
	}
	return cfg, nil
}
