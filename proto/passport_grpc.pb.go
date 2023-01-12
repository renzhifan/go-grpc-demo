// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/passport.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MemberClient is the client API for Member service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberClient interface {
	GetMemberProfileByUid(ctx context.Context, in *GetMemberProfileByUidReq, opts ...grpc.CallOption) (*GetMemberProfileByUidReply, error)
}

type memberClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberClient(cc grpc.ClientConnInterface) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) GetMemberProfileByUid(ctx context.Context, in *GetMemberProfileByUidReq, opts ...grpc.CallOption) (*GetMemberProfileByUidReply, error) {
	out := new(GetMemberProfileByUidReply)
	err := c.cc.Invoke(ctx, "/proto.Member/GetMemberProfileByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServer is the server API for Member service.
// All implementations must embed UnimplementedMemberServer
// for forward compatibility
type MemberServer interface {
	GetMemberProfileByUid(context.Context, *GetMemberProfileByUidReq) (*GetMemberProfileByUidReply, error)
	mustEmbedUnimplementedMemberServer()
}

// UnimplementedMemberServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServer struct {
}

func (UnimplementedMemberServer) GetMemberProfileByUid(context.Context, *GetMemberProfileByUidReq) (*GetMemberProfileByUidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberProfileByUid not implemented")
}
func (UnimplementedMemberServer) mustEmbedUnimplementedMemberServer() {}

// UnsafeMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServer will
// result in compilation errors.
type UnsafeMemberServer interface {
	mustEmbedUnimplementedMemberServer()
}

func RegisterMemberServer(s grpc.ServiceRegistrar, srv MemberServer) {
	s.RegisterService(&Member_ServiceDesc, srv)
}

func _Member_GetMemberProfileByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberProfileByUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).GetMemberProfileByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Member/GetMemberProfileByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).GetMemberProfileByUid(ctx, req.(*GetMemberProfileByUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Member_ServiceDesc is the grpc.ServiceDesc for Member service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Member_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMemberProfileByUid",
			Handler:    _Member_GetMemberProfileByUid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/passport.proto",
}