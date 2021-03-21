package cmd

import (
	"github.com/gofabian/denv/cfg"
	"github.com/urfave/cli/v2"
)

func loadDenvConfig(c *cli.Context) (*cfg.DenvConfig, error) {
	if c.IsSet("config") {
		return cfg.LoadConfigFrom(c.String("config"))
	} else {
		return cfg.LoadConfigFromDefaultDirs()
	}
}
