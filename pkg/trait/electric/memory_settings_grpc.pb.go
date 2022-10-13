// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pkg/trait/electric/memory_settings.proto

package electric

import (
	context "context"
	traits "github.com/smart-core-os/sc-api/go/traits"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MemorySettingsApiClient is the client API for MemorySettingsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemorySettingsApiClient interface {
	UpdateDemand(ctx context.Context, in *UpdateDemandRequest, opts ...grpc.CallOption) (*traits.ElectricDemand, error)
	// Create a new mode in the device.
	// Do not specify the id, this will be provided by the server.
	// This will not validate that exactly one normal mode exists,
	// the ClearActiveMode method will choose the first normal mode found.
	CreateMode(ctx context.Context, in *CreateModeRequest, opts ...grpc.CallOption) (*traits.ElectricMode, error)
	UpdateMode(ctx context.Context, in *UpdateModeRequest, opts ...grpc.CallOption) (*traits.ElectricMode, error)
	// Delete an existing mode.
	// Returns NOT_FOUND if the mode, identified by id, is not found. See allow_missing.
	DeleteMode(ctx context.Context, in *DeleteModeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type memorySettingsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMemorySettingsApiClient(cc grpc.ClientConnInterface) MemorySettingsApiClient {
	return &memorySettingsApiClient{cc}
}

func (c *memorySettingsApiClient) UpdateDemand(ctx context.Context, in *UpdateDemandRequest, opts ...grpc.CallOption) (*traits.ElectricDemand, error) {
	out := new(traits.ElectricDemand)
	err := c.cc.Invoke(ctx, "/smartcore.go.trait.electric.MemorySettingsApi/UpdateDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memorySettingsApiClient) CreateMode(ctx context.Context, in *CreateModeRequest, opts ...grpc.CallOption) (*traits.ElectricMode, error) {
	out := new(traits.ElectricMode)
	err := c.cc.Invoke(ctx, "/smartcore.go.trait.electric.MemorySettingsApi/CreateMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memorySettingsApiClient) UpdateMode(ctx context.Context, in *UpdateModeRequest, opts ...grpc.CallOption) (*traits.ElectricMode, error) {
	out := new(traits.ElectricMode)
	err := c.cc.Invoke(ctx, "/smartcore.go.trait.electric.MemorySettingsApi/UpdateMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memorySettingsApiClient) DeleteMode(ctx context.Context, in *DeleteModeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/smartcore.go.trait.electric.MemorySettingsApi/DeleteMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemorySettingsApiServer is the server API for MemorySettingsApi service.
// All implementations must embed UnimplementedMemorySettingsApiServer
// for forward compatibility
type MemorySettingsApiServer interface {
	UpdateDemand(context.Context, *UpdateDemandRequest) (*traits.ElectricDemand, error)
	// Create a new mode in the device.
	// Do not specify the id, this will be provided by the server.
	// This will not validate that exactly one normal mode exists,
	// the ClearActiveMode method will choose the first normal mode found.
	CreateMode(context.Context, *CreateModeRequest) (*traits.ElectricMode, error)
	UpdateMode(context.Context, *UpdateModeRequest) (*traits.ElectricMode, error)
	// Delete an existing mode.
	// Returns NOT_FOUND if the mode, identified by id, is not found. See allow_missing.
	DeleteMode(context.Context, *DeleteModeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMemorySettingsApiServer()
}

// UnimplementedMemorySettingsApiServer must be embedded to have forward compatible implementations.
type UnimplementedMemorySettingsApiServer struct {
}

func (UnimplementedMemorySettingsApiServer) UpdateDemand(context.Context, *UpdateDemandRequest) (*traits.ElectricDemand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDemand not implemented")
}
func (UnimplementedMemorySettingsApiServer) CreateMode(context.Context, *CreateModeRequest) (*traits.ElectricMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMode not implemented")
}
func (UnimplementedMemorySettingsApiServer) UpdateMode(context.Context, *UpdateModeRequest) (*traits.ElectricMode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMode not implemented")
}
func (UnimplementedMemorySettingsApiServer) DeleteMode(context.Context, *DeleteModeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMode not implemented")
}
func (UnimplementedMemorySettingsApiServer) mustEmbedUnimplementedMemorySettingsApiServer() {}

// UnsafeMemorySettingsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemorySettingsApiServer will
// result in compilation errors.
type UnsafeMemorySettingsApiServer interface {
	mustEmbedUnimplementedMemorySettingsApiServer()
}

func RegisterMemorySettingsApiServer(s grpc.ServiceRegistrar, srv MemorySettingsApiServer) {
	s.RegisterService(&MemorySettingsApi_ServiceDesc, srv)
}

func _MemorySettingsApi_UpdateDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDemandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemorySettingsApiServer).UpdateDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.trait.electric.MemorySettingsApi/UpdateDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemorySettingsApiServer).UpdateDemand(ctx, req.(*UpdateDemandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemorySettingsApi_CreateMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemorySettingsApiServer).CreateMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.trait.electric.MemorySettingsApi/CreateMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemorySettingsApiServer).CreateMode(ctx, req.(*CreateModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemorySettingsApi_UpdateMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemorySettingsApiServer).UpdateMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.trait.electric.MemorySettingsApi/UpdateMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemorySettingsApiServer).UpdateMode(ctx, req.(*UpdateModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemorySettingsApi_DeleteMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemorySettingsApiServer).DeleteMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.go.trait.electric.MemorySettingsApi/DeleteMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemorySettingsApiServer).DeleteMode(ctx, req.(*DeleteModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemorySettingsApi_ServiceDesc is the grpc.ServiceDesc for MemorySettingsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemorySettingsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.go.trait.electric.MemorySettingsApi",
	HandlerType: (*MemorySettingsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDemand",
			Handler:    _MemorySettingsApi_UpdateDemand_Handler,
		},
		{
			MethodName: "CreateMode",
			Handler:    _MemorySettingsApi_CreateMode_Handler,
		},
		{
			MethodName: "UpdateMode",
			Handler:    _MemorySettingsApi_UpdateMode_Handler,
		},
		{
			MethodName: "DeleteMode",
			Handler:    _MemorySettingsApi_DeleteMode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/trait/electric/memory_settings.proto",
}
