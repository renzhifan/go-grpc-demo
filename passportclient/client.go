package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultUid = 123456
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	uid  = flag.Uint64("uid", defaultUid, "UId to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMemberClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMemberProfileByUid(ctx, &pb.GetMemberProfileByUidReq{Uid: *uid, Option: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting:uid %d,nick:%s", r.GetUid(), r.GetNick())
}
