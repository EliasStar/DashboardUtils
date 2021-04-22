package command

import (
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

type ScreenCmd struct {
	GeneralCmd

	btn hardware.Pin
	//action   string
	msToggle uint
}

func (c ScreenCmd) Execute() {
	c.GeneralCmd.Execute()

	val, err := c.btn.Read()
	if err != nil {
		c.err = err
		return
	}

	c.err = c.btn.Write(!val)
	if c.err != nil {
		return
	}

	time.Sleep(time.Duration(c.msToggle) * time.Millisecond)

	c.err = c.btn.Write(val)
}
