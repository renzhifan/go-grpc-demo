package grpcconn

import (
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var once sync.Once

// Conn 全局 Conn
var Conn *grpc.ClientConn

func Connect(addr string) {
	once.Do(func() {
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		Conn = conn
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

	})
}
