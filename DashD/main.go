package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:"+misc.DashDPort)
	util.FatalIfErr(err)

	defer listener.Close()
	fmt.Println(listener.Addr())

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(con)
	}
}

func handleConnection(con net.Conn) {
	fmt.Println("New Connection:", con.RemoteAddr())

	var out string
	for {
		if _, err := fmt.Fscan(con, &out); err != nil {
			log.Println(err)
			return
		}

		fmt.Println("|", con.RemoteAddr(), "|>", out)
	}
}
