//
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'todo.pb.dart' as $0;

export 'todo.pb.dart';

@$pb.GrpcServiceName('greet.TodoService')
class TodoServiceClient extends $grpc.Client {
  static final _$changeNote = $grpc.ClientMethod<$0.ChangeNoteRequest, $0.ChangeNoteResponse>(
      '/greet.TodoService/ChangeNote',
      ($0.ChangeNoteRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ChangeNoteResponse.fromBuffer(value));
  static final _$completeTodo = $grpc.ClientMethod<$0.CompleteTodoRequest, $0.CompleteTodoResponse>(
      '/greet.TodoService/CompleteTodo',
      ($0.CompleteTodoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CompleteTodoResponse.fromBuffer(value));
  static final _$createTodo = $grpc.ClientMethod<$0.CreateTodoRequest, $0.CreateTodoResponse>(
      '/greet.TodoService/CreateTodo',
      ($0.CreateTodoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateTodoResponse.fromBuffer(value));
  static final _$getTodo = $grpc.ClientMethod<$0.GetTodoRequest, $0.GetTodoResponse>(
      '/greet.TodoService/GetTodo',
      ($0.GetTodoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GetTodoResponse.fromBuffer(value));
  static final _$listTodos = $grpc.ClientMethod<$0.ListTodosRequest, $0.ListTodosResponse>(
      '/greet.TodoService/ListTodos',
      ($0.ListTodosRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ListTodosResponse.fromBuffer(value));
  static final _$resetTodo = $grpc.ClientMethod<$0.ResetTodoRequest, $0.ResetTodoResponse>(
      '/greet.TodoService/ResetTodo',
      ($0.ResetTodoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ResetTodoResponse.fromBuffer(value));

  TodoServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.ChangeNoteResponse> changeNote($0.ChangeNoteRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$changeNote, request, options: options);
  }

  $grpc.ResponseFuture<$0.CompleteTodoResponse> completeTodo($0.CompleteTodoRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$completeTodo, request, options: options);
  }

  $grpc.ResponseFuture<$0.CreateTodoResponse> createTodo($0.CreateTodoRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createTodo, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetTodoResponse> getTodo($0.GetTodoRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getTodo, request, options: options);
  }

  $grpc.ResponseFuture<$0.ListTodosResponse> listTodos($0.ListTodosRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listTodos, request, options: options);
  }

  $grpc.ResponseFuture<$0.ResetTodoResponse> resetTodo($0.ResetTodoRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$resetTodo, request, options: options);
  }
}

@$pb.GrpcServiceName('greet.TodoService')
abstract class TodoServiceBase extends $grpc.Service {
  $core.String get $name => 'greet.TodoService';

  TodoServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ChangeNoteRequest, $0.ChangeNoteResponse>(
        'ChangeNote',
        changeNote_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ChangeNoteRequest.fromBuffer(value),
        ($0.ChangeNoteResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CompleteTodoRequest, $0.CompleteTodoResponse>(
        'CompleteTodo',
        completeTodo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CompleteTodoRequest.fromBuffer(value),
        ($0.CompleteTodoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CreateTodoRequest, $0.CreateTodoResponse>(
        'CreateTodo',
        createTodo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateTodoRequest.fromBuffer(value),
        ($0.CreateTodoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetTodoRequest, $0.GetTodoResponse>(
        'GetTodo',
        getTodo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetTodoRequest.fromBuffer(value),
        ($0.GetTodoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ListTodosRequest, $0.ListTodosResponse>(
        'ListTodos',
        listTodos_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ListTodosRequest.fromBuffer(value),
        ($0.ListTodosResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ResetTodoRequest, $0.ResetTodoResponse>(
        'ResetTodo',
        resetTodo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ResetTodoRequest.fromBuffer(value),
        ($0.ResetTodoResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.ChangeNoteResponse> changeNote_Pre($grpc.ServiceCall call, $async.Future<$0.ChangeNoteRequest> request) async {
    return changeNote(call, await request);
  }

  $async.Future<$0.CompleteTodoResponse> completeTodo_Pre($grpc.ServiceCall call, $async.Future<$0.CompleteTodoRequest> request) async {
    return completeTodo(call, await request);
  }

  $async.Future<$0.CreateTodoResponse> createTodo_Pre($grpc.ServiceCall call, $async.Future<$0.CreateTodoRequest> request) async {
    return createTodo(call, await request);
  }

  $async.Future<$0.GetTodoResponse> getTodo_Pre($grpc.ServiceCall call, $async.Future<$0.GetTodoRequest> request) async {
    return getTodo(call, await request);
  }

  $async.Future<$0.ListTodosResponse> listTodos_Pre($grpc.ServiceCall call, $async.Future<$0.ListTodosRequest> request) async {
    return listTodos(call, await request);
  }

  $async.Future<$0.ResetTodoResponse> resetTodo_Pre($grpc.ServiceCall call, $async.Future<$0.ResetTodoRequest> request) async {
    return resetTodo(call, await request);
  }

  $async.Future<$0.ChangeNoteResponse> changeNote($grpc.ServiceCall call, $0.ChangeNoteRequest request);
  $async.Future<$0.CompleteTodoResponse> completeTodo($grpc.ServiceCall call, $0.CompleteTodoRequest request);
  $async.Future<$0.CreateTodoResponse> createTodo($grpc.ServiceCall call, $0.CreateTodoRequest request);
  $async.Future<$0.GetTodoResponse> getTodo($grpc.ServiceCall call, $0.GetTodoRequest request);
  $async.Future<$0.ListTodosResponse> listTodos($grpc.ServiceCall call, $0.ListTodosRequest request);
  $async.Future<$0.ResetTodoResponse> resetTodo($grpc.ServiceCall call, $0.ResetTodoRequest request);
}
