package main

import (
	"bufio"
	"log"
	"net"
)

func echo(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %s bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(connection)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
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
