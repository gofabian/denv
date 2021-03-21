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
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "configuration `NAME` defined in '.denv.yml'",
		},
	},
	Action: run,
}

func run(c *cli.Context) error {
	config, err := loadRunConfig(c)
	if err != nil {
		return err
	}

	dockerCmd := c.Args().Slice()
	return execDockerRun(config.Image, nil, dockerCmd)
}

func loadRunConfig(c *cli.Context) (*cfg.NamedConfig, error) {
	if c.IsSet("image") {
		config := &cfg.NamedConfig{Image: c.String("image")}
		return config, nil
	}

	denvConfig, err := loadDenvConfig(c)
	if err != nil {
		return nil, err
	}

	name := c.String("name")
	config := denvConfig.GetByName(name)
	if config == nil {
		return nil, fmt.Errorf("unknown configuration name '%s'", name)
	}
	if config.Image == "" {
		return nil, fmt.Errorf("missing image")
	}
	return config, nil
}
