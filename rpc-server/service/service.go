package service

import (
	"errors"

	"github.com/aasourav/go-rpc-server/types"
)

type Arithmetic struct{}

func (t *Arithmetic) Multiply(args types.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arithmetic) Divide(args types.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}
