package client

import (
	"net/rpc"

	"github.com/aasourav/go-rpc-client/types"
)

func Multiply(client *rpc.Client, args types.Args) (int, error) {
	var reply int
	err := client.Call("Arithmetic.Multiply", args, &reply)
	return reply, err
}

func Divide(client *rpc.Client, args types.Args) (int, error) {
	var reply int
	err := client.Call("Arithmetic.Divide", args, &reply)
	return reply, err
}
