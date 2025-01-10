// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/recs/recs_service.proto

package recs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RecsService_GetPreferences_FullMethodName    = "/recs.RecsService/GetPreferences"
	RecsService_UpdatePreferences_FullMethodName = "/recs.RecsService/UpdatePreferences"
	RecsService_GetNewRec_FullMethodName         = "/recs.RecsService/GetNewRec"
	RecsService_DeletePeferences_FullMethodName  = "/recs.RecsService/DeletePeferences"
)

// RecsServiceClient is the client API for RecsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecsServiceClient interface {
	GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*GetPreferencesResponse, error)
	UpdatePreferences(ctx context.Context, in *UpdatePreferencesRequest, opts ...grpc.CallOption) (*UpdatePreferencesResponse, error)
	GetNewRec(ctx context.Context, in *GetNewRecRequest, opts ...grpc.CallOption) (*GetNewRecResponse, error)
	DeletePeferences(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type recsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecsServiceClient(cc grpc.ClientConnInterface) RecsServiceClient {
	return &recsServiceClient{cc}
}

func (c *recsServiceClient) GetPreferences(ctx context.Context, in *GetPreferencesRequest, opts ...grpc.CallOption) (*GetPreferencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPreferencesResponse)
	err := c.cc.Invoke(ctx, RecsService_GetPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recsServiceClient) UpdatePreferences(ctx context.Context, in *UpdatePreferencesRequest, opts ...grpc.CallOption) (*UpdatePreferencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePreferencesResponse)
	err := c.cc.Invoke(ctx, RecsService_UpdatePreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recsServiceClient) GetNewRec(ctx context.Context, in *GetNewRecRequest, opts ...grpc.CallOption) (*GetNewRecResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewRecResponse)
	err := c.cc.Invoke(ctx, RecsService_GetNewRec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recsServiceClient) DeletePeferences(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RecsService_DeletePeferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecsServiceServer is the server API for RecsService service.
// All implementations must embed UnimplementedRecsServiceServer
// for forward compatibility.
type RecsServiceServer interface {
	GetPreferences(context.Context, *GetPreferencesRequest) (*GetPreferencesResponse, error)
	UpdatePreferences(context.Context, *UpdatePreferencesRequest) (*UpdatePreferencesResponse, error)
	GetNewRec(context.Context, *GetNewRecRequest) (*GetNewRecResponse, error)
	DeletePeferences(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedRecsServiceServer()
}

// UnimplementedRecsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecsServiceServer struct{}

func (UnimplementedRecsServiceServer) GetPreferences(context.Context, *GetPreferencesRequest) (*GetPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreferences not implemented")
}
func (UnimplementedRecsServiceServer) UpdatePreferences(context.Context, *UpdatePreferencesRequest) (*UpdatePreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreferences not implemented")
}
func (UnimplementedRecsServiceServer) GetNewRec(context.Context, *GetNewRecRequest) (*GetNewRecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewRec not implemented")
}
func (UnimplementedRecsServiceServer) DeletePeferences(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeferences not implemented")
}
func (UnimplementedRecsServiceServer) mustEmbedUnimplementedRecsServiceServer() {}
func (UnimplementedRecsServiceServer) testEmbeddedByValue()                     {}

// UnsafeRecsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecsServiceServer will
// result in compilation errors.
type UnsafeRecsServiceServer interface {
	mustEmbedUnimplementedRecsServiceServer()
}

func RegisterRecsServiceServer(s grpc.ServiceRegistrar, srv RecsServiceServer) {
	// If the following call pancis, it indicates UnimplementedRecsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RecsService_ServiceDesc, srv)
}

func _RecsService_GetPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecsServiceServer).GetPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecsService_GetPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecsServiceServer).GetPreferences(ctx, req.(*GetPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecsService_UpdatePreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecsServiceServer).UpdatePreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecsService_UpdatePreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecsServiceServer).UpdatePreferences(ctx, req.(*UpdatePreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecsService_GetNewRec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewRecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecsServiceServer).GetNewRec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecsService_GetNewRec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecsServiceServer).GetNewRec(ctx, req.(*GetNewRecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecsService_DeletePeferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecsServiceServer).DeletePeferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecsService_DeletePeferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecsServiceServer).DeletePeferences(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RecsService_ServiceDesc is the grpc.ServiceDesc for RecsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recs.RecsService",
	HandlerType: (*RecsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreferences",
			Handler:    _RecsService_GetPreferences_Handler,
		},
		{
			MethodName: "UpdatePreferences",
			Handler:    _RecsService_UpdatePreferences_Handler,
		},
		{
			MethodName: "GetNewRec",
			Handler:    _RecsService_GetNewRec_Handler,
		},
		{
			MethodName: "DeletePeferences",
			Handler:    _RecsService_DeletePeferences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/recs/recs_service.proto",
}
