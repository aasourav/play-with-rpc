package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/aasourav/go-rpc-client/client"
	"github.com/aasourav/go-rpc-client/types"
)

func main() {
	clientConn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer clientConn.Close()

	args := types.Args{A: 10, B: 12}

	reply, err := client.Multiply(clientConn, args)
	if err != nil {
		log.Fatalf("Error calling Multiply: %v", err)
	}
	fmt.Printf("10 * 2 = %d\n", reply)

	reply, err = client.Divide(clientConn, args)
	if err != nil {
		log.Fatalf("Error calling Divide: %v", err)
	}
	fmt.Printf("10 / 2 = %d\n", reply)
}
