package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(len(os.Args), os.Args)
	args := os.Args

	arguments := []string{args[1], url.QueryEscape(args[2]), args[3], args[4]}

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	fmt.Printf("%v", arguments)
	cmd := exec.Command(fmt.Sprintf("migrate %s", args[7]), arguments...)

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
