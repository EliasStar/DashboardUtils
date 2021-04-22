package command

import (
	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

type LedstripCmd struct {
	GeneralCmd

	colors []uint32
	leds   []uint
}

var strip hardware.Ledstrip

func (c LedstripCmd) Execute() {
	c.GeneralCmd.Execute()

	if len(c.leds) == 0 {
		c.err = strip.SetStripColor(c.colors[0])
	}

	for i, v := range c.leds {
		c.err = strip.SetLEDColor(v, c.colors[i])
		if c.err != nil {
			return
		}
	}
}
