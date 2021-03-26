package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func createDockerRunCmd(image string, dockerCmd []string, dockerArgs []string) (args []string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cmd := []string{"docker", "run", "--rm", "-it", "-v", cwd + ":/denv/workdir", "-w", "/denv/workdir"}
	cmd = append(cmd, dockerArgs...)
	cmd = append(cmd, image)
	cmd = append(cmd, dockerCmd...)
	return cmd, nil
}

func executeCmd(cmd []string, stdout io.Writer, stderr io.Writer) {
	printCmd(cmd)
	cliCmd := exec.Command(cmd[0], cmd[1:]...)
	cliCmd.Stdin = os.Stdin
	cliCmd.Stdout = stdout
	cliCmd.Stderr = stderr
	cliCmd.Run()
	exitCode := cliCmd.ProcessState.ExitCode()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func printCmd(cmd []string) {
	escapedArgs := make([]string, len(cmd))
	for i, a := range cmd {
		if strings.ContainsRune(a, ' ') {
			escapedArgs[i] = "\"" + a + "\""
		} else {
			escapedArgs[i] = a
		}
	}
	fmt.Printf("+ %s\n", strings.Join(escapedArgs, " "))
}
