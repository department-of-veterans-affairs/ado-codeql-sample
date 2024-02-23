package main

import (
	"os"
	"os/exec"
	"strings"
)

// write an example main function that executes an arbitrary command passed in as the first argument to the program
func main() {
	param := os.Args[1]
	split := strings.Split(param, " ")
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
