package command

import (
	"os/exec"
)

type LaunchCmd struct {
	GeneralCmd

	exe  string
	args []string
}

func (c LaunchCmd) Execute() {
	c.GeneralCmd.Execute()
	c.err = exec.Command(c.exe, c.args...).Run()
}
