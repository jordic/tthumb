package main

import (
	"io"
	"os"
	"os/exec"
)

var baseCmd = []string{
	"-colorspace",
	"gray",
	"-level-colors",
}

type ImgCmd struct {
	Preset string
	Params []string
	File   string
	Out    string
	Color1 string
	Color2 string
}

func (i *ImgCmd) Run() error {
	cmd := make([]string, 0)
	cmd = append(cmd, i.File)
	cmd = append(cmd,
		i.File,
		"-colorspace",
		"gray",
		"-level-colors",
		i.Color1,
		i.Color2)

	//cmd := string[]{ i.File }
	return nil
}

func runOut(w io.Writer, c string, args ...string) error {
	cmd := exec.Command(c, args...)
	cmd.Stdin = nil
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
