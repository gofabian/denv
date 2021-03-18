package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execDockerRun(image string, dockerCmd []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	args := []string{"docker", "run", "--rm", "-it", "-v", cwd + ":/denv/workdir",
		"-w", "/denv/workdir", image}
	args = append(args, dockerCmd...)

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
