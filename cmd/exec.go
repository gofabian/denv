package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gofabian/denv/cfg"
	"github.com/urfave/cli/v2"
)

var ExecCommand = &cli.Command{
	Name:  "exec",
	Usage: "execute steps in Docker container",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "configuration `NAME` defined in '.denv.yml'",
			Value:   cli.NewStringSlice(""), // default: empty name
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "execute all configurations",
		},
	},
	Action: runExec,
}

func runExec(c *cli.Context) error {
	denvConfig, err := loadDenvConfig(c)
	if err != nil {
		return err
	}

	var namedConfigs []cfg.NamedConfig
	if c.Bool("all") {
		namedConfigs = denvConfig.GetAll()
	} else {
		names := c.StringSlice("name")
		namedConfigs = denvConfig.GetByNames(names...)
	}

	if len(namedConfigs) == 0 {
		return fmt.Errorf("no executable configuration")
	}

	for _, namedConfig := range namedConfigs {
		if len(namedConfig.Exec) == 0 {
			continue
		}

		err = executeNamedConfig(&namedConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func executeNamedConfig(namedConfig *cfg.NamedConfig) error {
	// create script on host
	file, err := ioutil.TempFile("", "denv-exec-*.sh")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	fmt.Fprintf(file, scriptTemplate, strings.Join(namedConfig.Exec, "\n"))

	// run script in Docker container
	dockerArgs := []string{"-v", file.Name() + ":/denv/exec.sh"}
	dockerCmd := []string{"/bin/sh", "/denv/exec.sh"}
	cmd, err := createDockerRunCmd(namedConfig.Image, dockerCmd, dockerArgs)
	if err != nil {
		return err
	}
	executeCmd(
		cmd,
		NewPrefixWriter("\x1b[32m[stdout]\x1b[0m ", os.Stdout),
		NewPrefixWriter("\x1b[31m[stderr]\x1b[0m ", os.Stderr),
	)
	return nil
}

var scriptTemplate string = `set -e
set -x
%s`
