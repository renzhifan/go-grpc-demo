package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"helloworld/controller"
	pb "helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &controller.HelloworldService{})
	pb.RegisterMemberServer(s, &controller.PassportService{})
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
