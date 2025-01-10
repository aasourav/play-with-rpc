package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/aasourav/go-rpc-server/service"
)

func main() {
	arith := new(service.Arithmetic)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalf("Error registering service: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Println("Server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
