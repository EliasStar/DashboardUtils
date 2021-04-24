package ledstrip

import (
	"context"
	"errors"
	"image/color"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/hardware"
)

type LedstripCmd struct {
	Animation       LedstripAnimation
	LEDs            []uint
	Colors          []color.Color
	AnimationLength time.Duration
}

func (l LedstripCmd) IsValid(ctx context.Context) bool {
	strip, ok := ctx.Value("strip").(hardware.Ledstrip)
	if !ok {
		return false
	}

	for _, v := range l.LEDs {
		if v >= strip.Length() {
			return false
		}
	}

	a := l.Animation.IsValid()
	b := len(l.Colors) == 1 || (len(l.Colors) > 1 && len(l.LEDs) == len(l.Colors))
	c := 0 <= l.AnimationLength
	d := l.AnimationLength.Minutes() <= 5

	return a && b && c && d
}

func (l LedstripCmd) Execute(ctx *context.Context) error {
	_, ok := (*ctx).Value("strip").(hardware.Ledstrip)
	if !ok {
		return errors.New("ledstrip not initialized")
	}

	return nil
}
