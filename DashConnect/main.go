package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/EliasStar/DashboardUtils/Commons/util"
	"github.com/EliasStar/DashboardUtils/Commons/util/misc"
)

func main() {
	con, err := net.Dial("tcp", os.Args[1]+":"+misc.DashDPort)
	util.FatalIfErr(err)

	defer con.Close()

	var in string
	for {
		fmt.Scan(&in)

		if _, err := fmt.Fprintln(con, in); err != nil {
			log.Println(err)
			return
		}
	}
}
