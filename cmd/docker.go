package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execDockerRun(image string, dockerArgs []string, dockerCmd []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cliCmd := []string{"docker", "run", "--rm", "-it", "-v", cwd + ":/denv/workdir",
		"-w", "/denv/workdir"}
	cliCmd = append(cliCmd, dockerArgs...)
	cliCmd = append(cliCmd, image)
	cliCmd = append(cliCmd, dockerCmd...)

	exitCode := execCommand(cliCmd)
	os.Exit(exitCode)
	return nil
}

func execCommand(args []string) int {
	printCommand(args)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return cmd.ProcessState.ExitCode()
}

func printCommand(args []string) {
	escapedArgs := make([]string, len(args))
	for i, a := range args {
		if strings.ContainsRune(a, ' ') {
			escapedArgs[i] = "\"" + a + "\""
		} else {
			escapedArgs[i] = a
		}
	}
	fmt.Printf("+ %s\n", strings.Join(escapedArgs, " "))
}
