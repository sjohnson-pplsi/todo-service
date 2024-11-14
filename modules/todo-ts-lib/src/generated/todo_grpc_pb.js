// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var todo_pb = require('./todo_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_greet_CompleteTodoRequest(arg) {
  if (!(arg instanceof todo_pb.CompleteTodoRequest)) {
    throw new Error('Expected argument of type greet.CompleteTodoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_CompleteTodoRequest(buffer_arg) {
  return todo_pb.CompleteTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_CompleteTodoResponse(arg) {
  if (!(arg instanceof todo_pb.CompleteTodoResponse)) {
    throw new Error('Expected argument of type greet.CompleteTodoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_CompleteTodoResponse(buffer_arg) {
  return todo_pb.CompleteTodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_CreateTodoRequest(arg) {
  if (!(arg instanceof todo_pb.CreateTodoRequest)) {
    throw new Error('Expected argument of type greet.CreateTodoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_CreateTodoRequest(buffer_arg) {
  return todo_pb.CreateTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_CreateTodoResponse(arg) {
  if (!(arg instanceof todo_pb.CreateTodoResponse)) {
    throw new Error('Expected argument of type greet.CreateTodoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_CreateTodoResponse(buffer_arg) {
  return todo_pb.CreateTodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_GetTodoRequest(arg) {
  if (!(arg instanceof todo_pb.GetTodoRequest)) {
    throw new Error('Expected argument of type greet.GetTodoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_GetTodoRequest(buffer_arg) {
  return todo_pb.GetTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_GetTodoResponse(arg) {
  if (!(arg instanceof todo_pb.GetTodoResponse)) {
    throw new Error('Expected argument of type greet.GetTodoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_GetTodoResponse(buffer_arg) {
  return todo_pb.GetTodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_ListTodosRequest(arg) {
  if (!(arg instanceof todo_pb.ListTodosRequest)) {
    throw new Error('Expected argument of type greet.ListTodosRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_ListTodosRequest(buffer_arg) {
  return todo_pb.ListTodosRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_ListTodosResponse(arg) {
  if (!(arg instanceof todo_pb.ListTodosResponse)) {
    throw new Error('Expected argument of type greet.ListTodosResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_ListTodosResponse(buffer_arg) {
  return todo_pb.ListTodosResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_ResetTodoRequest(arg) {
  if (!(arg instanceof todo_pb.ResetTodoRequest)) {
    throw new Error('Expected argument of type greet.ResetTodoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_ResetTodoRequest(buffer_arg) {
  return todo_pb.ResetTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_greet_ResetTodoResponse(arg) {
  if (!(arg instanceof todo_pb.ResetTodoResponse)) {
    throw new Error('Expected argument of type greet.ResetTodoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_greet_ResetTodoResponse(buffer_arg) {
  return todo_pb.ResetTodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The greeting service definition.
var TodoServiceService = exports.TodoServiceService = {
  completeTodo: {
    path: '/greet.TodoService/CompleteTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.CompleteTodoRequest,
    responseType: todo_pb.CompleteTodoResponse,
    requestSerialize: serialize_greet_CompleteTodoRequest,
    requestDeserialize: deserialize_greet_CompleteTodoRequest,
    responseSerialize: serialize_greet_CompleteTodoResponse,
    responseDeserialize: deserialize_greet_CompleteTodoResponse,
  },
  createTodo: {
    path: '/greet.TodoService/CreateTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.CreateTodoRequest,
    responseType: todo_pb.CreateTodoResponse,
    requestSerialize: serialize_greet_CreateTodoRequest,
    requestDeserialize: deserialize_greet_CreateTodoRequest,
    responseSerialize: serialize_greet_CreateTodoResponse,
    responseDeserialize: deserialize_greet_CreateTodoResponse,
  },
  getTodo: {
    path: '/greet.TodoService/GetTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.GetTodoRequest,
    responseType: todo_pb.GetTodoResponse,
    requestSerialize: serialize_greet_GetTodoRequest,
    requestDeserialize: deserialize_greet_GetTodoRequest,
    responseSerialize: serialize_greet_GetTodoResponse,
    responseDeserialize: deserialize_greet_GetTodoResponse,
  },
  listTodos: {
    path: '/greet.TodoService/ListTodos',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.ListTodosRequest,
    responseType: todo_pb.ListTodosResponse,
    requestSerialize: serialize_greet_ListTodosRequest,
    requestDeserialize: deserialize_greet_ListTodosRequest,
    responseSerialize: serialize_greet_ListTodosResponse,
    responseDeserialize: deserialize_greet_ListTodosResponse,
  },
  resetTodo: {
    path: '/greet.TodoService/ResetTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.ResetTodoRequest,
    responseType: todo_pb.ResetTodoResponse,
    requestSerialize: serialize_greet_ResetTodoRequest,
    requestDeserialize: deserialize_greet_ResetTodoRequest,
    responseSerialize: serialize_greet_ResetTodoResponse,
    responseDeserialize: deserialize_greet_ResetTodoResponse,
  },
};

exports.TodoServiceClient = grpc.makeGenericClientConstructor(TodoServiceService);
