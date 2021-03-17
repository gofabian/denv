package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"docker", "run", "--rm", "-it",
		"-v", cwd + ":/denv/workdir", "-w", "/denv/workdir"}
	args = append(args, os.Args[1:]...)

	exitCode := execCommand(args)
	os.Exit(exitCode)
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
