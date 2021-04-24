package hardware

import (
	"image/color"

	"github.com/EliasStar/DashboardUtils/Commons/util"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func MakeLedstrip(pin Pin, ledCount uint, addBurnerLED bool) (Ledstrip, error) {
	opt := ws2811.DefaultOptions
	channel := &opt.Channels[0]

	channel.GpioPin = int(pin)
	channel.LedCount = int(ledCount)
	channel.Brightness = 255

	if addBurnerLED {
		channel.LedCount++
	}

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		return Ledstrip{}, err
	}

	return Ledstrip{dev, addBurnerLED}, nil
}

type Ledstrip struct {
	*ws2811.WS2811

	hasBurnerLED bool
}

func (l *Ledstrip) GetLEDs() []uint32 {
	if l.hasBurnerLED {
		return l.Leds(0)[1:]
	}

	return l.Leds(0)
}

func (l *Ledstrip) SetSingleLEDColor(index uint, c color.Color) {
	r, g, b, _ := c.RGBA()
	l.GetLEDs()[index] = r<<16 | g<<8 | b
}

func (l *Ledstrip) SetLEDColor(indicies []uint, c color.Color) {
	leds := l.GetLEDs()
	r, g, b, _ := c.RGBA()

	for _, v := range indicies {
		leds[v] = r<<16 | g<<8 | b
	}
}

func (l *Ledstrip) SetLEDColors(indicies []uint, c []color.Color) {
	leds := l.GetLEDs()

	for i := 0; i < util.Min(len(indicies), len(c)); i++ {
		r, g, b, _ := c[i].RGBA()
		leds[indicies[i]] = r<<16 | g<<8 | b
	}
}

func (l *Ledstrip) SetStripColor(c color.Color) {
	leds := l.GetLEDs()
	r, g, b, _ := c.RGBA()

	for i := 0; i < len(leds); i++ {
		leds[i] = r<<16 | g<<8 | b
	}
}

func (l *Ledstrip) Length() uint {
	return uint(len(l.GetLEDs()))
}
