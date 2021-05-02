package ledstrip

import (
	"image/color"
)

type LedstripRst []color.Color

func (l LedstripRst) IsOK() bool {
	return true
}

func (l LedstripRst) Err() error {
	return nil
}
