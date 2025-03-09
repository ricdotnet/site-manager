package main

import (
	"bytes"
	"github.com/charmbracelet/log"
	"io"
	"os"
	"os/exec"
	"strings"
)

func execute(cmd string, args []string) (string, error) {
	log.Infof("executing: %s %s", cmd, strings.Join(args, " "))

	command := exec.Command(cmd, args...)

	var errBuf, outBuf bytes.Buffer
	command.Stderr = io.MultiWriter(os.Stderr, &errBuf)
	command.Stdout = &outBuf
	err := command.Run()

	return outBuf.String(), err
}
