package main

import (
	"io"
	"log"
	"net"
)

func echo(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		size, err := connection.Read(buffer[0:])
		if err == io.EOF {
			log.Println("Client disconnected.")
			break
		}
		if err != nil {
			log.Println("Unexpected error.")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(buffer))

		log.Println("Writing data")
		if _, err = connection.Write(buffer[0:size]); err != nil {
			log.Println("Unable to write data.")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}
	log.Println("Listening on " + listener.Addr().String())
	for {
		connection, err := listener.Accept()
		log.Println("Client connected.")
		if err != nil {
			log.Fatalln("Unable to accept connection.")
		}

		go echo(connection)
	}
}
