// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/org-service/v1/services/org.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/organization-service/api/org-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvOrgV1_Ping_FullMethodName = "/saas.api.org.servicev1.SrvOrgV1/Ping"
)

// SrvOrgV1Client is the client API for SrvOrgV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvOrgV1Client interface {
	// Ping ping
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
}

type srvOrgV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvOrgV1Client(cc grpc.ClientConnInterface) SrvOrgV1Client {
	return &srvOrgV1Client{cc}
}

func (c *srvOrgV1Client) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvOrgV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvOrgV1Server is the server API for SrvOrgV1 service.
// All implementations must embed UnimplementedSrvOrgV1Server
// for forward compatibility
type SrvOrgV1Server interface {
	// Ping ping
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	mustEmbedUnimplementedSrvOrgV1Server()
}

// UnimplementedSrvOrgV1Server must be embedded to have forward compatible implementations.
type UnimplementedSrvOrgV1Server struct {
}

func (UnimplementedSrvOrgV1Server) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvOrgV1Server) mustEmbedUnimplementedSrvOrgV1Server() {}

// UnsafeSrvOrgV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvOrgV1Server will
// result in compilation errors.
type UnsafeSrvOrgV1Server interface {
	mustEmbedUnimplementedSrvOrgV1Server()
}

func RegisterSrvOrgV1Server(s grpc.ServiceRegistrar, srv SrvOrgV1Server) {
	s.RegisterService(&SrvOrgV1_ServiceDesc, srv)
}

func _SrvOrgV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvOrgV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvOrgV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvOrgV1Server).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvOrgV1_ServiceDesc is the grpc.ServiceDesc for SrvOrgV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvOrgV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.org.servicev1.SrvOrgV1",
	HandlerType: (*SrvOrgV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvOrgV1_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/org-service/v1/services/org.service.v1.proto",
}
