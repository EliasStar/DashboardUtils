package main

import (
	"time"

	"github.com/EliasStar/DashboardUtils/common"
)

func main() {

	strip, err := newLedstrip(data, 61, true)
	common.FatalIfErr(err)

	common.FatalIfErr(strip.Init())
	defer strip.Fini()

	strip.display(0xff0000, 1000*time.Millisecond)
	strip.display(0x00ff00, 1000*time.Millisecond)
	strip.display(0x0000ff, 1000*time.Millisecond)
	strip.display(0x000000, 500*time.Millisecond)
}

func (ws *ledstrip) display(color uint32, sleepTime time.Duration) error {
	for i := 0; i < len(ws.Leds(0)); i++ {
		ws.Leds(0)[i] = color

		if err := ws.Render(); err != nil {
			return err
		}

		time.Sleep(sleepTime)
	}

	return nil
}

func (ws *ledstrip) wipeFromStartToEnd() {

}

func (ws *ledstrip) wipeFromEndToStart() {

}
