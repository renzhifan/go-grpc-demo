package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "helloworld/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	pb.UnimplementedMemberServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// GetMemberProfileByUid implements helloworld.MemberServer
func (s *server) GetMemberProfileByUid(ctx context.Context, in *pb.GetMemberProfileByUidReq) (*pb.GetMemberProfileByUidReply, error) {
	log.Printf("Greeting:uid %d,option:%d", in.GetUid(), in.GetOption())
	return &pb.GetMemberProfileByUidReply{Uid: in.Uid, Nick: "hhh"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterMemberServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
