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

// 判断 MemberServer 是否实现了 passport.proto.GetMemberProfileByUid
// 不要问为什么这么写，就好像「地球为什么是椭圆的」
var _ pb.MemberServer = (*PassportService)(nil)

// GetMemberProfileByUid implements helloworld.MemberServer
func (s *PassportService) GetMemberProfileByUid(ctx context.Context, in *pb.GetMemberProfileByUidReq) (*pb.GetMemberProfileByUidReply, error) {
	log.Printf("Greeting:uid %d,option:%d", in.GetUid(), in.GetOption())
	return &pb.GetMemberProfileByUidReply{Uid: in.Uid, Nick: "hhh"}, nil
}
