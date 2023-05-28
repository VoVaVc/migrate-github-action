package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

func quote(s string) string {
	return (&url.URL{Path: s}).RequestURI()
}

func main() {
	args := os.Args

	arguments := []string{
		"-path",
		args[1],
		"-database",
		fmt.Sprintf("%q", quote(args[2])),
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
