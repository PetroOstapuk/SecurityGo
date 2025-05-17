package main

import (
	"io"
	"log"
	"net"
)

func handler(src net.Conn) {
	dst, err := net.Dial("tcp", "cyberdev.space:443")
	if err != nil {
		log.Fatalln("Unable to connect ot our unreachable host")
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handler(conn)
	}
}
