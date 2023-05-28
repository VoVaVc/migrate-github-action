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
	fmt.Printf("%v", arguments)

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	arguments = append(arguments, args[7])
	cmd := exec.Command("migrate", arguments...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}
