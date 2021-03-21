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
	config, err := loadRunConfig(c)
	if err != nil {
		return err
	}

	args := c.Args().Slice()
	return execDockerRun(config.Image, args)
}

func loadRunConfig(c *cli.Context) (*cfg.NamedConfig, error) {
	if c.IsSet("image") {
		config := &cfg.NamedConfig{Image: c.String("image")}
		return config, nil
	}

	denvConfig, err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}

	config := denvConfig.GetByName("")
	if config.Image == "" {
		return nil, fmt.Errorf("missing image")
	}
	return config, nil
}
