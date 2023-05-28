package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func queryEscape(s string) string {
	split := strings.Split(s, "://")
	usernamePassword := strings.Split(split[1], "@")
	password := strings.Split(usernamePassword[0], ":")
	replaced := strings.Replace(s, password[1], url.QueryEscape(usernamePassword[0]), 3)

	return fmt.Sprintf("\n %s", replaced)
}

func main() {
	args := os.Args

	arguments := []string{
		"-path",
		args[1],
		"-database",
		queryEscape(args[2]),
		"-prefetch",
		args[3],
		"-lock-timeout",
		args[4],
	}

	fmt.Printf("Arguments are: migrate %v", args)

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	arguments = append(arguments, args[7])

	fmt.Printf("Command is: migrate %v", arguments)
	cmd := exec.Command("migrate", arguments...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}
}
