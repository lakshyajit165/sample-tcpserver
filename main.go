package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	// reading a sample buffer; typically a request is read
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing the request")
	// mimicking server processing time
	time.Sleep(8 * time.Second)

	/* added this message as it is the format for curl request
		Note: Can add formatted message also
	*/
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	// close the connection after returning the response
	conn.Close()
}

func main() {
	// server listening on a particular port
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	// wait for clients inside an infinite loop
	for {
		log.Println("waiting for a client to connect")
		// executed whenever a client tries to connect to the server
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")
		// spin up a new thread(go-routine) for each connection
		go do(conn)
	}
}