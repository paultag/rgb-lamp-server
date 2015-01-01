package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn, s io.ReadWriteCloser) {
	buff := bufio.NewReader(conn)
    for {
        line, _, err := buff.ReadLine()
        if err != nil {
            return
        }

        _, err = s.Write([]byte(line))
        fmt.Printf("Got: %s\n", line)
        if err != nil {
            return
        }
    }
}

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":2017")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			go handleConnection(conn, s)
		}
	}

}
