// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: orderitems.proto

package genprotos

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

// OrderItemServiceClient is the client API for OrderItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderItemServiceClient interface {
	AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error)
	DeleteItems(ctx context.Context, in *DeleItemsRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	ListOrderItems(ctx context.Context, in *GetByUseridItems, opts ...grpc.CallOption) (*GetAllItems, error)
}

type orderItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderItemServiceClient(cc grpc.ClientConnInterface) OrderItemServiceClient {
	return &orderItemServiceClient{cc}
}

func (c *orderItemServiceClient) AddItems(ctx context.Context, in *AddItemsRequest, opts ...grpc.CallOption) (*AddItemsResponse, error) {
	out := new(AddItemsResponse)
	err := c.cc.Invoke(ctx, "/protos.OrderItemService/AddItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) DeleteItems(ctx context.Context, in *DeleItemsRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, "/protos.OrderItemService/DeleteItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemServiceClient) ListOrderItems(ctx context.Context, in *GetByUseridItems, opts ...grpc.CallOption) (*GetAllItems, error) {
	out := new(GetAllItems)
	err := c.cc.Invoke(ctx, "/protos.OrderItemService/ListOrderItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderItemServiceServer is the server API for OrderItemService service.
// All implementations must embed UnimplementedOrderItemServiceServer
// for forward compatibility
type OrderItemServiceServer interface {
	AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error)
	DeleteItems(context.Context, *DeleItemsRequest) (*DeleteProductResponse, error)
	ListOrderItems(context.Context, *GetByUseridItems) (*GetAllItems, error)
	mustEmbedUnimplementedOrderItemServiceServer()
}

// UnimplementedOrderItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderItemServiceServer struct {
}

func (UnimplementedOrderItemServiceServer) AddItems(context.Context, *AddItemsRequest) (*AddItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItems not implemented")
}
func (UnimplementedOrderItemServiceServer) DeleteItems(context.Context, *DeleItemsRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItems not implemented")
}
func (UnimplementedOrderItemServiceServer) ListOrderItems(context.Context, *GetByUseridItems) (*GetAllItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrderItems not implemented")
}
func (UnimplementedOrderItemServiceServer) mustEmbedUnimplementedOrderItemServiceServer() {}

// UnsafeOrderItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderItemServiceServer will
// result in compilation errors.
type UnsafeOrderItemServiceServer interface {
	mustEmbedUnimplementedOrderItemServiceServer()
}

func RegisterOrderItemServiceServer(s grpc.ServiceRegistrar, srv OrderItemServiceServer) {
	s.RegisterService(&OrderItemService_ServiceDesc, srv)
}

func _OrderItemService_AddItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).AddItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OrderItemService/AddItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).AddItems(ctx, req.(*AddItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_DeleteItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).DeleteItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OrderItemService/DeleteItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).DeleteItems(ctx, req.(*DeleItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemService_ListOrderItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUseridItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemServiceServer).ListOrderItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OrderItemService/ListOrderItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemServiceServer).ListOrderItems(ctx, req.(*GetByUseridItems))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderItemService_ServiceDesc is the grpc.ServiceDesc for OrderItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.OrderItemService",
	HandlerType: (*OrderItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItems",
			Handler:    _OrderItemService_AddItems_Handler,
		},
		{
			MethodName: "DeleteItems",
			Handler:    _OrderItemService_DeleteItems_Handler,
		},
		{
			MethodName: "ListOrderItems",
			Handler:    _OrderItemService_ListOrderItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderitems.proto",
}