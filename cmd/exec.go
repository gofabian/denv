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
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "configuration `NAME` defined in '.denv.yml'",
		},
	},
	Action: runExec,
}

func runExec(c *cli.Context) error {
	denvConfig, err := loadDenvConfig(c)
	if err != nil {
		return err
	}

	name := c.String("name")
	namedConfigs := denvConfig.GetByName(name)
	if len(namedConfigs) == 0 {
		return fmt.Errorf("unknown configuration name '%s'", name)
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
	return execDockerRun(namedConfig.Image, dockerArgs, dockerCmd)
}

var scriptTemplate string = `set -e
set -x
%s`
