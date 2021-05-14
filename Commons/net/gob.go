package net

import (
	"encoding/gob"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/display"
	"github.com/EliasStar/DashboardUtils/Commons/command/launch"
	"github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	"github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util/color"
)

func InitGOB() {
	gob.Register(color.RGB{})
	gob.Register(color.RGBA32{})

	gob.Register(command.ErrorRst{})
	gob.Register(command.OKRst{})

	gob.Register(display.DisplayCmd{})

	gob.Register(launch.LaunchCmd{})
	gob.Register(launch.LaunchRst(""))

	gob.Register(ledstrip.LedstripCmd{})
	gob.Register(ledstrip.LedstripRst{})

	gob.Register(schedule.ScheduleCmd{})

	gob.Register(screen.ScreenCmd{})
	gob.Register(screen.ScreenRst(false))
}
