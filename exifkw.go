package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		exit(1, fmt.Sprintln("Usage:", os.Args[0], "\"keyword[,keyword...]\" paths"))
	}
	args := []string{"-P", "-overwrite_original"}

	for _, kw := range strings.Split(os.Args[1], ",") {
		kw = strings.TrimSpace(kw)
		if kw == "" {
			continue
		}
		args = append(args, fmt.Sprintf("-keywords-=%s", kw), fmt.Sprintf("-keywords+=%s", kw))
	}
	if len(args) == 2 {
		exit(2, "No keywords supplied")
	}

	args = append(args, os.Args[2:]...)
	cmd := exec.Command("exiftool", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			os.Exit(e.ExitCode())
		}
		exit(3, err.Error())
	}
}

func exit(status int, message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Stderr.Sync()
	os.Exit(status)
}
