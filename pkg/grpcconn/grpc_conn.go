package grpcconn

import (
	"helloworld/pkg/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect() *grpc.ClientConn {
	conn, err := grpc.Dial(config.Get("grpc.host")+":"+config.Get("grpc.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn

}
