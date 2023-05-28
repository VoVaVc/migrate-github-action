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

	arguments := []string{
		"-path",
		fmt.Sprintf("\"%s\"", args[1]),
		"-database",
		fmt.Sprintf("\"%s\"", url.QueryEscape(args[2])),
		"-prefetch",
		fmt.Sprintf("\"%s\"", args[3]),
		"-lock-timeout",
		fmt.Sprintf("\"%s\"", args[4]),
	}

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	arguments = append(arguments, args[7])

	fmt.Printf("%v", arguments)
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
