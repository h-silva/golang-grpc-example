// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package equipment

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

// EquipmentServiceClient is the client API for EquipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EquipmentServiceClient interface {
	Create(ctx context.Context, in *EquipmentRequest, opts ...grpc.CallOption) (*EquipmentResponse, error)
	Get(ctx context.Context, in *EquipmentCode, opts ...grpc.CallOption) (*EquipmentResponse, error)
}

type equipmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquipmentServiceClient(cc grpc.ClientConnInterface) EquipmentServiceClient {
	return &equipmentServiceClient{cc}
}

func (c *equipmentServiceClient) Create(ctx context.Context, in *EquipmentRequest, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/equipment.EquipmentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *equipmentServiceClient) Get(ctx context.Context, in *EquipmentCode, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/equipment.EquipmentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquipmentServiceServer is the server API for EquipmentService service.
// All implementations must embed UnimplementedEquipmentServiceServer
// for forward compatibility
type EquipmentServiceServer interface {
	Create(context.Context, *EquipmentRequest) (*EquipmentResponse, error)
	Get(context.Context, *EquipmentCode) (*EquipmentResponse, error)
	mustEmbedUnimplementedEquipmentServiceServer()
}

// UnimplementedEquipmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEquipmentServiceServer struct {
}

func (UnimplementedEquipmentServiceServer) Create(context.Context, *EquipmentRequest) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEquipmentServiceServer) Get(context.Context, *EquipmentCode) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEquipmentServiceServer) mustEmbedUnimplementedEquipmentServiceServer() {}

// UnsafeEquipmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EquipmentServiceServer will
// result in compilation errors.
type UnsafeEquipmentServiceServer interface {
	mustEmbedUnimplementedEquipmentServiceServer()
}

func RegisterEquipmentServiceServer(s grpc.ServiceRegistrar, srv EquipmentServiceServer) {
	s.RegisterService(&EquipmentService_ServiceDesc, srv)
}

func _EquipmentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/equipment.EquipmentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).Create(ctx, req.(*EquipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EquipmentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipmentCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquipmentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/equipment.EquipmentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquipmentServiceServer).Get(ctx, req.(*EquipmentCode))
	}
	return interceptor(ctx, in, info, handler)
}

// EquipmentService_ServiceDesc is the grpc.ServiceDesc for EquipmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EquipmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "equipment.EquipmentService",
	HandlerType: (*EquipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EquipmentService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EquipmentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/equipment/equipment.proto",
}
