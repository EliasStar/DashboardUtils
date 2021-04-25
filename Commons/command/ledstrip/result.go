package ledstrip

import "image/color"

type LedstripRst struct {
	Colors []color.Color
	Error  error
}

func (s LedstripRst) OK() bool {
	return s.Error == nil
}

func (s LedstripRst) Err() error {
	return s.Error
}
