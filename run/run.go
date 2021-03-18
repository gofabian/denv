package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	args := []string{"docker", "run", "--rm", "-it", "-v", cwd + ":/denv/workdir",
		"-w", "/denv/workdir", c.String("image")}
	args = append(args, c.Args().Slice()...)

	exitCode := execCommand(args)
	os.Exit(exitCode)
	return nil
}

func execCommand(args []string) int {
	fmt.Printf("+ %s\n", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}
