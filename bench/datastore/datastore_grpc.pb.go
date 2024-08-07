// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: datastore.proto

package datastore

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

const (
	Datastore_Get_FullMethodName           = "/datastore.Datastore/Get"
	Datastore_Set_FullMethodName           = "/datastore.Datastore/Set"
	Datastore_Query_FullMethodName         = "/datastore.Datastore/Query"
	Datastore_Delete_FullMethodName        = "/datastore.Datastore/Delete"
	Datastore_DeleteAtIndex_FullMethodName = "/datastore.Datastore/DeleteAtIndex"
)

// DatastoreClient is the client API for Datastore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatastoreClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteAtIndex(ctx context.Context, in *DeleteAtIndexRequest, opts ...grpc.CallOption) (*DeleteAtIndexResponse, error)
}

type datastoreClient struct {
	cc grpc.ClientConnInterface
}

func NewDatastoreClient(cc grpc.ClientConnInterface) DatastoreClient {
	return &datastoreClient{cc}
}

func (c *datastoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Datastore_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, Datastore_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, Datastore_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Datastore_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datastoreClient) DeleteAtIndex(ctx context.Context, in *DeleteAtIndexRequest, opts ...grpc.CallOption) (*DeleteAtIndexResponse, error) {
	out := new(DeleteAtIndexResponse)
	err := c.cc.Invoke(ctx, Datastore_DeleteAtIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatastoreServer is the server API for Datastore service.
// All implementations must embed UnimplementedDatastoreServer
// for forward compatibility
type DatastoreServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteAtIndex(context.Context, *DeleteAtIndexRequest) (*DeleteAtIndexResponse, error)
	mustEmbedUnimplementedDatastoreServer()
}

// UnimplementedDatastoreServer must be embedded to have forward compatible implementations.
type UnimplementedDatastoreServer struct {
}

func (UnimplementedDatastoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDatastoreServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedDatastoreServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDatastoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDatastoreServer) DeleteAtIndex(context.Context, *DeleteAtIndexRequest) (*DeleteAtIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAtIndex not implemented")
}
func (UnimplementedDatastoreServer) mustEmbedUnimplementedDatastoreServer() {}

// UnsafeDatastoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatastoreServer will
// result in compilation errors.
type UnsafeDatastoreServer interface {
	mustEmbedUnimplementedDatastoreServer()
}

func RegisterDatastoreServer(s grpc.ServiceRegistrar, srv DatastoreServer) {
	s.RegisterService(&Datastore_ServiceDesc, srv)
}

func _Datastore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datastore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datastore_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datastore_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datastore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datastore_DeleteAtIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAtIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatastoreServer).DeleteAtIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Datastore_DeleteAtIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatastoreServer).DeleteAtIndex(ctx, req.(*DeleteAtIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Datastore_ServiceDesc is the grpc.ServiceDesc for Datastore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datastore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datastore.Datastore",
	HandlerType: (*DatastoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Datastore_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Datastore_Set_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Datastore_Query_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Datastore_Delete_Handler,
		},
		{
			MethodName: "DeleteAtIndex",
			Handler:    _Datastore_DeleteAtIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datastore.proto",
}
