package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/nsf/termbox-go"
	"net"
	"time"
)

func main() {
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleConnection(listener)
	}

}

func handleConnection(con *net.UDPConn) {

	for {

		buf := make([]byte, 2000)
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		buff := bytes.NewReader(buf[0:n])

		var cord struct {
			X int32
			Y int32
		}

		err = binary.Read(buff, binary.LittleEndian, &cord)
		if err != nil {
			fmt.Println(err)
			return
		}

		x := int(cord.X)
		y := int(cord.Y)

		termbox.Init()
		termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, y, '*', termbox.ColorWhite, termbox.ColorBlack)
		termbox.Flush()
		time.Sleep(1 * time.Second)
	}
	termbox.Close()
}
