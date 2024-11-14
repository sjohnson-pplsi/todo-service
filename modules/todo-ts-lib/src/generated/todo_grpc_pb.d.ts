// package: greet
// file: todo.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as todo_pb from "./todo_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

interface ITodoServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    completeTodo: ITodoServiceService_ICompleteTodo;
    createTodo: ITodoServiceService_ICreateTodo;
    getTodo: ITodoServiceService_IGetTodo;
    listTodos: ITodoServiceService_IListTodos;
    resetTodo: ITodoServiceService_IResetTodo;
}

interface ITodoServiceService_ICompleteTodo extends grpc.MethodDefinition<todo_pb.CompleteTodoRequest, todo_pb.CompleteTodoResponse> {
    path: "/greet.TodoService/CompleteTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.CompleteTodoRequest>;
    requestDeserialize: grpc.deserialize<todo_pb.CompleteTodoRequest>;
    responseSerialize: grpc.serialize<todo_pb.CompleteTodoResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.CompleteTodoResponse>;
}
interface ITodoServiceService_ICreateTodo extends grpc.MethodDefinition<todo_pb.CreateTodoRequest, todo_pb.CreateTodoResponse> {
    path: "/greet.TodoService/CreateTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.CreateTodoRequest>;
    requestDeserialize: grpc.deserialize<todo_pb.CreateTodoRequest>;
    responseSerialize: grpc.serialize<todo_pb.CreateTodoResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.CreateTodoResponse>;
}
interface ITodoServiceService_IGetTodo extends grpc.MethodDefinition<todo_pb.GetTodoRequest, todo_pb.GetTodoResponse> {
    path: "/greet.TodoService/GetTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.GetTodoRequest>;
    requestDeserialize: grpc.deserialize<todo_pb.GetTodoRequest>;
    responseSerialize: grpc.serialize<todo_pb.GetTodoResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.GetTodoResponse>;
}
interface ITodoServiceService_IListTodos extends grpc.MethodDefinition<todo_pb.ListTodosRequest, todo_pb.ListTodosResponse> {
    path: "/greet.TodoService/ListTodos";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.ListTodosRequest>;
    requestDeserialize: grpc.deserialize<todo_pb.ListTodosRequest>;
    responseSerialize: grpc.serialize<todo_pb.ListTodosResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.ListTodosResponse>;
}
interface ITodoServiceService_IResetTodo extends grpc.MethodDefinition<todo_pb.ResetTodoRequest, todo_pb.ResetTodoResponse> {
    path: "/greet.TodoService/ResetTodo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<todo_pb.ResetTodoRequest>;
    requestDeserialize: grpc.deserialize<todo_pb.ResetTodoRequest>;
    responseSerialize: grpc.serialize<todo_pb.ResetTodoResponse>;
    responseDeserialize: grpc.deserialize<todo_pb.ResetTodoResponse>;
}

export const TodoServiceService: ITodoServiceService;

export interface ITodoServiceServer extends grpc.UntypedServiceImplementation {
    completeTodo: grpc.handleUnaryCall<todo_pb.CompleteTodoRequest, todo_pb.CompleteTodoResponse>;
    createTodo: grpc.handleUnaryCall<todo_pb.CreateTodoRequest, todo_pb.CreateTodoResponse>;
    getTodo: grpc.handleUnaryCall<todo_pb.GetTodoRequest, todo_pb.GetTodoResponse>;
    listTodos: grpc.handleUnaryCall<todo_pb.ListTodosRequest, todo_pb.ListTodosResponse>;
    resetTodo: grpc.handleUnaryCall<todo_pb.ResetTodoRequest, todo_pb.ResetTodoResponse>;
}

export interface ITodoServiceClient {
    completeTodo(request: todo_pb.CompleteTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    completeTodo(request: todo_pb.CompleteTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    completeTodo(request: todo_pb.CompleteTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    createTodo(request: todo_pb.CreateTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    createTodo(request: todo_pb.CreateTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    createTodo(request: todo_pb.CreateTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    getTodo(request: todo_pb.GetTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    getTodo(request: todo_pb.GetTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    getTodo(request: todo_pb.GetTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    listTodos(request: todo_pb.ListTodosRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    listTodos(request: todo_pb.ListTodosRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    listTodos(request: todo_pb.ListTodosRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    resetTodo(request: todo_pb.ResetTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
    resetTodo(request: todo_pb.ResetTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
    resetTodo(request: todo_pb.ResetTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
}

export class TodoServiceClient extends grpc.Client implements ITodoServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public completeTodo(request: todo_pb.CompleteTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    public completeTodo(request: todo_pb.CompleteTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    public completeTodo(request: todo_pb.CompleteTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.CompleteTodoResponse) => void): grpc.ClientUnaryCall;
    public createTodo(request: todo_pb.CreateTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    public createTodo(request: todo_pb.CreateTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    public createTodo(request: todo_pb.CreateTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.CreateTodoResponse) => void): grpc.ClientUnaryCall;
    public getTodo(request: todo_pb.GetTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    public getTodo(request: todo_pb.GetTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    public getTodo(request: todo_pb.GetTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.GetTodoResponse) => void): grpc.ClientUnaryCall;
    public listTodos(request: todo_pb.ListTodosRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    public listTodos(request: todo_pb.ListTodosRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    public listTodos(request: todo_pb.ListTodosRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.ListTodosResponse) => void): grpc.ClientUnaryCall;
    public resetTodo(request: todo_pb.ResetTodoRequest, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
    public resetTodo(request: todo_pb.ResetTodoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
    public resetTodo(request: todo_pb.ResetTodoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: todo_pb.ResetTodoResponse) => void): grpc.ClientUnaryCall;
}
