// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package memory

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

// MemoryPowerSupplySettingsApiClient is the client API for MemoryPowerSupplySettingsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoryPowerSupplySettingsApiClient interface {
	GetSettings(ctx context.Context, in *GetMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (*MemoryPowerSupplySettings, error)
	UpdateSettings(ctx context.Context, in *UpdateMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (*MemoryPowerSupplySettings, error)
	PullSettings(ctx context.Context, in *PullMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (MemoryPowerSupplySettingsApi_PullSettingsClient, error)
}

type memoryPowerSupplySettingsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryPowerSupplySettingsApiClient(cc grpc.ClientConnInterface) MemoryPowerSupplySettingsApiClient {
	return &memoryPowerSupplySettingsApiClient{cc}
}

func (c *memoryPowerSupplySettingsApiClient) GetSettings(ctx context.Context, in *GetMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (*MemoryPowerSupplySettings, error) {
	out := new(MemoryPowerSupplySettings)
	err := c.cc.Invoke(ctx, "/smartcore.go.mem.MemoryPowerSupplySettingsApi/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryPowerSupplySettingsApiClient) UpdateSettings(ctx context.Context, in *UpdateMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (*MemoryPowerSupplySettings, error) {
	out := new(MemoryPowerSupplySettings)
	err := c.cc.Invoke(ctx, "/smartcore.go.mem.MemoryPowerSupplySettingsApi/UpdateSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memoryPowerSupplySettingsApiClient) PullSettings(ctx context.Context, in *PullMemoryPowerSupplySettingsReq, opts ...grpc.CallOption) (MemoryPowerSupplySettingsApi_PullSettingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MemoryPowerSupplySettingsApi_ServiceDesc.Streams[0], "/smartcore.go.mem.MemoryPowerSupplySettingsApi/PullSettings", opts...)
	if err != nil {
		return nil, err
	}
	x := &memoryPowerSupplySettingsApiPullSettingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MemoryPowerSupplySettingsApi_PullSettingsClient interface {
	Recv() (*PullMemoryPowerSupplySettingsRes, error)
	grpc.ClientStream
}

type memoryPowerSupplySettingsApiPullSettingsClient struct {
	grpc.ClientStream
}

func (x *memoryPowerSupplySettingsApiPullSettingsClient) Recv() (*PullMemoryPowerSupplySettingsRes, error) {
	m := new(PullMemoryPowerSupplySettingsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MemoryPowerSupplySettingsApiServer is the server API for MemoryPowerSupplySettingsApi service.
// All implementations must embed UnimplementedMemoryPowerSupplySettingsApiServer
// for forward compatibility
type MemoryPowerSupplySettingsApiServer interface {
	GetSettings(context.Context, *GetMemoryPowerSupplySettingsReq) (*MemoryPowerSupplySettings, error)
	UpdateSettings(context.Context, *UpdateMemoryPowerSupplySettingsReq) (*MemoryPowerSupplySettings, error)
	PullSettings(*PullMemoryPowerSupplySettingsReq, MemoryPowerSupplySettingsApi_PullSettingsServer) error
	mustEmbedUnimplementedMemoryPowerSupplySettingsApiServer()
}

// UnimplementedMemoryPowerSupplySettingsApiServer must be embedded to have forward compatible implementations.
type UnimplementedMemoryPowerSupplySettingsApiServer struct {
}

func (UnimplementedMemoryPowerSupplySettingsApiServer) GetSettings(context.Context, *GetMemoryPowerSupplySettingsReq) (*MemoryPowerSupplySettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (UnimplementedMemoryPowerSupplySettingsApiServer) UpdateSettings(context.Context, *UpdateMemoryPowerSupplySettingsReq) (*MemoryPowerSupplySettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSettings not implemented")
}
func (UnimplementedMemoryPowerSupplySettingsApiServer) PullSettings(*PullMemoryPowerSupplySettingsReq, MemoryPowerSupplySettingsApi_PullSettingsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullSettings not implemented")
}
func (UnimplementedMemoryPowerSupplySettingsApiServer) mustEmbedUnimplementedMemoryPowerSupplySettingsApiServer() {
}

// UnsafeMemoryPowerSupplySettingsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoryPowerSupplySettingsApiServer will
// result in compilation errors.
type UnsafeMemoryPowerSupplySettingsApiServer interface {
	mustEmbedUnimplementedMemoryPowerSupplySettingsApiServer()
}

func RegisterMemoryPowerSupplySettingsApiServer(s grpc.ServiceRegistrar, srv MemoryPowerSupplySettingsApiServer) {
	s.RegisterService(&MemoryPowerSupplySettingsApi_ServiceDesc, srv)
}

func _MemoryPowerSupplySettingsApi_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoryPowerSupplySettingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryPowerSupplySettingsApiServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.mem.MemoryPowerSupplySettingsApi/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryPowerSupplySettingsApiServer).GetSettings(ctx, req.(*GetMemoryPowerSupplySettingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryPowerSupplySettingsApi_UpdateSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemoryPowerSupplySettingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryPowerSupplySettingsApiServer).UpdateSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.mem.MemoryPowerSupplySettingsApi/UpdateSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryPowerSupplySettingsApiServer).UpdateSettings(ctx, req.(*UpdateMemoryPowerSupplySettingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemoryPowerSupplySettingsApi_PullSettings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullMemoryPowerSupplySettingsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MemoryPowerSupplySettingsApiServer).PullSettings(m, &memoryPowerSupplySettingsApiPullSettingsServer{stream})
}

type MemoryPowerSupplySettingsApi_PullSettingsServer interface {
	Send(*PullMemoryPowerSupplySettingsRes) error
	grpc.ServerStream
}

type memoryPowerSupplySettingsApiPullSettingsServer struct {
	grpc.ServerStream
}

func (x *memoryPowerSupplySettingsApiPullSettingsServer) Send(m *PullMemoryPowerSupplySettingsRes) error {
	return x.ServerStream.SendMsg(m)
}

// MemoryPowerSupplySettingsApi_ServiceDesc is the grpc.ServiceDesc for MemoryPowerSupplySettingsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoryPowerSupplySettingsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.go.mem.MemoryPowerSupplySettingsApi",
	HandlerType: (*MemoryPowerSupplySettingsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler:    _MemoryPowerSupplySettingsApi_GetSettings_Handler,
		},
		{
			MethodName: "UpdateSettings",
			Handler:    _MemoryPowerSupplySettingsApi_UpdateSettings_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullSettings",
			Handler:       _MemoryPowerSupplySettingsApi_PullSettings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "power_supply_settings.proto",
}
