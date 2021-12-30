// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VnfmUIClient is the client API for VnfmUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VnfmUIClient interface {
	// ListVimSpecs returns a list of Virtualized Infrastructure Manager(s) that can be started through the UI.
	ListVimSpecs(ctx context.Context, in *ListVimSpecsRequest, opts ...grpc.CallOption) (VnfmUI_ListVimSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type vnfmUIClient struct {
	cc grpc.ClientConnInterface
}

func NewVnfmUIClient(cc grpc.ClientConnInterface) VnfmUIClient {
	return &vnfmUIClient{cc}
}

func (c *vnfmUIClient) ListVimSpecs(ctx context.Context, in *ListVimSpecsRequest, opts ...grpc.CallOption) (VnfmUI_ListVimSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &VnfmUI_ServiceDesc.Streams[0], "/v1.VnfmUI/ListVimSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &vnfmUIListVimSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VnfmUI_ListVimSpecsClient interface {
	Recv() (*ListVimSpecsResponse, error)
	grpc.ClientStream
}

type vnfmUIListVimSpecsClient struct {
	grpc.ClientStream
}

func (x *vnfmUIListVimSpecsClient) Recv() (*ListVimSpecsResponse, error) {
	m := new(ListVimSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vnfmUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.VnfmUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VnfmUIServer is the server API for VnfmUI service.
// All implementations must embed UnimplementedVnfmUIServer
// for forward compatibility
type VnfmUIServer interface {
	// ListVimSpecs returns a list of Virtualized Infrastructure Manager(s) that can be started through the UI.
	ListVimSpecs(*ListVimSpecsRequest, VnfmUI_ListVimSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedVnfmUIServer()
}

// UnimplementedVnfmUIServer must be embedded to have forward compatible implementations.
type UnimplementedVnfmUIServer struct {
}

func (UnimplementedVnfmUIServer) ListVimSpecs(*ListVimSpecsRequest, VnfmUI_ListVimSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListVimSpecs not implemented")
}
func (UnimplementedVnfmUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedVnfmUIServer) mustEmbedUnimplementedVnfmUIServer() {}

// UnsafeVnfmUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VnfmUIServer will
// result in compilation errors.
type UnsafeVnfmUIServer interface {
	mustEmbedUnimplementedVnfmUIServer()
}

func RegisterVnfmUIServer(s grpc.ServiceRegistrar, srv VnfmUIServer) {
	s.RegisterService(&VnfmUI_ServiceDesc, srv)
}

func _VnfmUI_ListVimSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListVimSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VnfmUIServer).ListVimSpecs(m, &vnfmUIListVimSpecsServer{stream})
}

type VnfmUI_ListVimSpecsServer interface {
	Send(*ListVimSpecsResponse) error
	grpc.ServerStream
}

type vnfmUIListVimSpecsServer struct {
	grpc.ServerStream
}

func (x *vnfmUIListVimSpecsServer) Send(m *ListVimSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VnfmUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VnfmUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VnfmUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VnfmUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VnfmUI_ServiceDesc is the grpc.ServiceDesc for VnfmUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VnfmUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VnfmUI",
	HandlerType: (*VnfmUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _VnfmUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListVimSpecs",
			Handler:       _VnfmUI_ListVimSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vnfm-ui.proto",
}
