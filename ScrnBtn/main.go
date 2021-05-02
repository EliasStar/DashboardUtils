package main

import (
	"context"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/EliasStar/DashboardUtils/Commons/command"
	"github.com/EliasStar/DashboardUtils/Commons/command/screen"
	. "github.com/EliasStar/DashboardUtils/Commons/command/screen"
	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	con, conErr := net.Dial("tcp", "127.0.0.1:"+misc.DashDPort)
	if conErr != nil {
		for _, b := range ScreenButtons() {
			util.PanicIfErr(b.Pin().Mode(true))
		}
	} else {
		defer con.Close()
	}

	cmd := parseCommand()

	var rst command.Result
	if conErr == nil {
		gob.Register(command.ErrorRst{})
		gob.Register(command.OKRst{})
		gob.Register(screen.ScreenCmd{})
		gob.Register(screen.ScreenRst(false))

		util.PanicIfErr(gob.NewEncoder(con).Encode(&cmd))
		util.PanicIfErr(gob.NewDecoder(con).Decode(&rst))
	} else {
		rst = cmd.Execute(context.Background())
	}

	if !rst.IsOK() {
		log.Fatal(rst.Err())
	}

	snRst, ok := rst.(ScreenRst)
	if ok {
		fmt.Println(snRst)
	}
}

func parseCommand() command.Command {
	switch os.Args[1] {
	case "read", "press", "release":
		return ScreenCmd{
			Action:      ScreenAction(os.Args[1]),
			Button:      parseButton(os.Args[2]),
			ToggleDelay: 0,
		}

	case "toggle":
		set := flag.NewFlagSet("toggle", flag.ContinueOnError)
		delay := set.Uint("delay", 250, "`ms` between pressing and releasing on toggle")
		util.PanicIfErr(set.Parse(os.Args[2:]))

		return ScreenCmd{
			Action:      ActionToggle,
			Button:      parseButton(set.Arg(0)),
			ToggleDelay: time.Duration(*delay) * time.Millisecond,
		}

	case "reset":
		for _, b := range ScreenButtons() {
			b.Pin().Write(false)
			b.Pin().Mode(false)
		}
		os.Exit(0)

	default:
		log.Panic("screen {read|press|release|toggle|reset} [-delay=<ms>] [{power|menu|plus|minus|source}]")
	}

	return nil
}

func parseButton(pin string) ScreenButton {
	switch pin {
	case "power":
		return ButtonPower

	case "menu":
		return ButtonMenu

	case "plus":
		return ButtonPlus

	case "minus":
		return ButtonMinus

	case "source":
		return ButtonSource
	}

	log.Fatal("possible pin names: power, menu, plus, minus, source")
	return 0
}
