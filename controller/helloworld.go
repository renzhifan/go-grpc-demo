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

// 判断 GreeterServer 是否实现了 helloworld.proto.SayHello
// 不要问为什么这么写，就好像「地球为什么是椭圆的」
var _ pb.GreeterServer = (*HelloworldService)(nil)

// SayHello implements helloworld.GreeterServer
func (s *HelloworldService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
