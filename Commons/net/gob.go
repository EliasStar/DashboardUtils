package net

import (
	"encoding/gob"

	. "github.com/EliasStar/DashboardUtils/Commons/command"
	. "github.com/EliasStar/DashboardUtils/Commons/command/display"
	. "github.com/EliasStar/DashboardUtils/Commons/command/launch"
	. "github.com/EliasStar/DashboardUtils/Commons/command/ledstrip"
	. "github.com/EliasStar/DashboardUtils/Commons/command/schedule"
	. "github.com/EliasStar/DashboardUtils/Commons/command/screen"
)

func InitGOB() {
	gob.Register(ErrorRst{})
	gob.Register(OKRst{})

	gob.Register(DisplayCmd{})

	gob.Register(LaunchCmd{})
	gob.Register(LaunchRst(""))

	gob.Register(LedstripCmd{})
	gob.Register(LedstripRst{})

	gob.Register(ScheduleCmd{})

	gob.Register(ScreenCmd{})
	gob.Register(ScreenRst(false))
}
