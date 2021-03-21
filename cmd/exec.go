package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	config := denvConfig.GetByName(name)
	if config == nil {
		return fmt.Errorf("unknown configuration name '%s'", name)
	}
	if len(config.Exec) == 0 {
		return fmt.Errorf("Missing exec steps")
	}

	file, err := ioutil.TempFile("", "denv-exec-*.sh")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	fmt.Fprintf(file, scriptTemplate, strings.Join(config.Exec, "\n"))

	dockerArgs := []string{"-v", file.Name() + ":/denv/exec.sh"}
	dockerCmd := []string{"/bin/sh", "/denv/exec.sh"}
	return execDockerRun(config.Image, dockerArgs, dockerCmd)
}

var scriptTemplate string = `set -e
set -x
%s`
