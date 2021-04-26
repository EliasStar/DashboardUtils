package ledstrip

import (
	"image/color"

	"github.com/EliasStar/DashboardUtils/Commons/command"
)

type LedstripRst struct {
	command.ErrorRst
	Colors []color.Color
}
