package controller

import (
	"context"
	pb "helloworld/proto"
	"log"
)

//PassportService is used to implement helloworld.MemberServer.
type PassportService struct {
	pb.UnimplementedMemberServer
}

// GetMemberProfileByUid implements helloworld.MemberServer
func (s *PassportService) GetMemberProfileByUid(ctx context.Context, in *pb.GetMemberProfileByUidReq) (*pb.GetMemberProfileByUidReply, error) {
	log.Printf("Greeting:uid %d,option:%d", in.GetUid(), in.GetOption())
	return &pb.GetMemberProfileByUidReply{Uid: in.Uid, Nick: "hhh"}, nil
}
