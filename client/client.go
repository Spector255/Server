package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var cord struct {
		X int32
		Y int32
	}
	for {

		cord.X = int32(rand.Intn(100))
		cord.Y = int32(rand.Intn(100))
		
		time.Sleep(1 * time.Second)

		var buf bytes.Buffer
		err = binary.Write(&buf, binary.LittleEndian, cord)

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Close()
	}
}
