package util

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// RunCommand runs an arbitrary os/exec.Cmd command as if you were in a terminal at the given path
func RunCommand(cmd *exec.Cmd, path string, noOutput bool) string {
	cmd.Dir = path

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if !noOutput {
		if err != nil {
			log.Println(fmt.Sprint(err) + ": " + stderr.String())
		}
		fmt.Printf(out.String())
	}

	return out.String() + stderr.String()
}
