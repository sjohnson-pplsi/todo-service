// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
// source: todo.proto

package togo_go_lib

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CompleteTodo(ctx context.Context, in *CompleteTodoRequest, opts ...grpc.CallOption) (*CompleteTodoResponse, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
	ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error)
	ResetTodo(ctx context.Context, in *ResetTodoRequest, opts ...grpc.CallOption) (*ResetTodoResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CompleteTodo(ctx context.Context, in *CompleteTodoRequest, opts ...grpc.CallOption) (*CompleteTodoResponse, error) {
	out := new(CompleteTodoResponse)
	err := c.cc.Invoke(ctx, "/greet.TodoService/CompleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/greet.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/greet.TodoService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error) {
	out := new(ListTodosResponse)
	err := c.cc.Invoke(ctx, "/greet.TodoService/ListTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ResetTodo(ctx context.Context, in *ResetTodoRequest, opts ...grpc.CallOption) (*ResetTodoResponse, error) {
	out := new(ResetTodoResponse)
	err := c.cc.Invoke(ctx, "/greet.TodoService/ResetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CompleteTodo(context.Context, *CompleteTodoRequest) (*CompleteTodoResponse, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
	ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error)
	ResetTodo(context.Context, *ResetTodoRequest) (*ResetTodoResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CompleteTodo(context.Context, *CompleteTodoRequest) (*CompleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTodo not implemented")
}
func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodoServiceServer) ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
func (UnimplementedTodoServiceServer) ResetTodo(context.Context, *ResetTodoRequest) (*ResetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CompleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CompleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.TodoService/CompleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CompleteTodo(ctx, req.(*CompleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.TodoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.TodoService/ListTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodos(ctx, req.(*ListTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ResetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ResetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.TodoService/ResetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ResetTodo(ctx, req.(*ResetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteTodo",
			Handler:    _TodoService_CompleteTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
		{
			MethodName: "ListTodos",
			Handler:    _TodoService_ListTodos_Handler,
		},
		{
			MethodName: "ResetTodo",
			Handler:    _TodoService_ResetTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}