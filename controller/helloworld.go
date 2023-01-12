package controller

import (
	"context"
	pb "helloworld/proto"
	"log"
)

//HelloworldService is used to implement helloworld.GreeterServer.
type HelloworldService struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *HelloworldService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
