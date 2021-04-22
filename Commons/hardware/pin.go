package hardware

import (
	"os/exec"
	"strconv"
)

type Pin uint

func (p Pin) Mode(out bool) error {
	val := "in"
	if out {
		val = "out"
	}

	return exec.Command("gpio", "-g", "mode", p.String(), val).Run()
}

func (p Pin) Write(value bool) error {
	val := "0"
	if value {
		val = "1"
	}

	return exec.Command("gpio", "-g", "write", p.String(), val).Run()
}

func (p Pin) Read() (value bool, err error) {
	out, err := exec.Command("gpio", "-g", "read", p.String()).Output()
	value = string(out[0]) == "1"
	return
}

func (p Pin) String() string {
	return strconv.Itoa(int(p))
}
