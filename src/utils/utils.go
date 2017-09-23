package utils

import (
    "fmt"
	  "google.golang.org/grpc"
)

func MustDial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("failed to dial: %v", err)
		panic(err)
	}
	return conn
}
