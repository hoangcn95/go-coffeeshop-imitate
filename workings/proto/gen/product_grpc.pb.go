// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: product.proto

package gen

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetItemTypes(ctx context.Context, in *GetItemTypesRequest, opts ...grpc.CallOption) (*GetItemTypesResponse, error)
	GetItemsByType(ctx context.Context, in *GetItemsByTypeRequest, opts ...grpc.CallOption) (*GetItemsByTypeResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetItemTypes(ctx context.Context, in *GetItemTypesRequest, opts ...grpc.CallOption) (*GetItemTypesResponse, error) {
	out := new(GetItemTypesResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetItemTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetItemsByType(ctx context.Context, in *GetItemsByTypeRequest, opts ...grpc.CallOption) (*GetItemsByTypeResponse, error) {
	out := new(GetItemsByTypeResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductService/GetItemsByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	GetItemTypes(context.Context, *GetItemTypesRequest) (*GetItemTypesResponse, error)
	GetItemsByType(context.Context, *GetItemsByTypeRequest) (*GetItemsByTypeResponse, error)
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetItemTypes(context.Context, *GetItemTypesRequest) (*GetItemTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemTypes not implemented")
}
func (UnimplementedProductServiceServer) GetItemsByType(context.Context, *GetItemsByTypeRequest) (*GetItemsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemsByType not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetItemTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetItemTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetItemTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetItemTypes(ctx, req.(*GetItemTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetItemsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetItemsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/GetItemsByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetItemsByType(ctx, req.(*GetItemsByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItemTypes",
			Handler:    _ProductService_GetItemTypes_Handler,
		},
		{
			MethodName: "GetItemsByType",
			Handler:    _ProductService_GetItemsByType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
