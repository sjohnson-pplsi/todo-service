//
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use todoStatusDescriptor instead')
const TodoStatus$json = {
  '1': 'TodoStatus',
  '2': [
    {'1': 'TODO_STATUS_UNSPECIFIED', '2': 0},
    {'1': 'TODO_STATUS_INCOMPLETE', '2': 1},
    {'1': 'TODO_STATUS_COMPLETE', '2': 2},
  ],
};

/// Descriptor for `TodoStatus`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List todoStatusDescriptor = $convert.base64Decode(
    'CgpUb2RvU3RhdHVzEhsKF1RPRE9fU1RBVFVTX1VOU1BFQ0lGSUVEEAASGgoWVE9ET19TVEFUVV'
    'NfSU5DT01QTEVURRABEhgKFFRPRE9fU1RBVFVTX0NPTVBMRVRFEAI=');

@$core.Deprecated('Use completeTodoRequestDescriptor instead')
const CompleteTodoRequest$json = {
  '1': 'CompleteTodoRequest',
  '2': [
    {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '10': 'todoId'},
  ],
};

/// Descriptor for `CompleteTodoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List completeTodoRequestDescriptor = $convert.base64Decode(
    'ChNDb21wbGV0ZVRvZG9SZXF1ZXN0EhcKB3RvZG9faWQYASABKAlSBnRvZG9JZA==');

@$core.Deprecated('Use completeTodoResponseDescriptor instead')
const CompleteTodoResponse$json = {
  '1': 'CompleteTodoResponse',
};

/// Descriptor for `CompleteTodoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List completeTodoResponseDescriptor = $convert.base64Decode(
    'ChRDb21wbGV0ZVRvZG9SZXNwb25zZQ==');

@$core.Deprecated('Use createTodoRequestDescriptor instead')
const CreateTodoRequest$json = {
  '1': 'CreateTodoRequest',
  '2': [
    {'1': 'note', '3': 1, '4': 1, '5': 9, '10': 'note'},
    {'1': 'due', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'due'},
  ],
};

/// Descriptor for `CreateTodoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTodoRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVUb2RvUmVxdWVzdBISCgRub3RlGAEgASgJUgRub3RlEiwKA2R1ZRgCIAEoCzIaLm'
    'dvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSA2R1ZQ==');

@$core.Deprecated('Use createTodoResponseDescriptor instead')
const CreateTodoResponse$json = {
  '1': 'CreateTodoResponse',
  '2': [
    {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '10': 'todoId'},
  ],
};

/// Descriptor for `CreateTodoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTodoResponseDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVUb2RvUmVzcG9uc2USFwoHdG9kb19pZBgBIAEoCVIGdG9kb0lk');

@$core.Deprecated('Use getTodoRequestDescriptor instead')
const GetTodoRequest$json = {
  '1': 'GetTodoRequest',
  '2': [
    {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '10': 'todoId'},
  ],
};

/// Descriptor for `GetTodoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTodoRequestDescriptor = $convert.base64Decode(
    'Cg5HZXRUb2RvUmVxdWVzdBIXCgd0b2RvX2lkGAEgASgJUgZ0b2RvSWQ=');

@$core.Deprecated('Use getTodoResponseDescriptor instead')
const GetTodoResponse$json = {
  '1': 'GetTodoResponse',
  '2': [
    {'1': 'todo', '3': 1, '4': 1, '5': 11, '6': '.greet.Todo', '10': 'todo'},
  ],
};

/// Descriptor for `GetTodoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTodoResponseDescriptor = $convert.base64Decode(
    'Cg9HZXRUb2RvUmVzcG9uc2USHwoEdG9kbxgBIAEoCzILLmdyZWV0LlRvZG9SBHRvZG8=');

@$core.Deprecated('Use listTodosRequestDescriptor instead')
const ListTodosRequest$json = {
  '1': 'ListTodosRequest',
  '2': [
    {'1': 'limit', '3': 1, '4': 1, '5': 5, '10': 'limit'},
    {'1': 'offset', '3': 2, '4': 1, '5': 5, '10': 'offset'},
  ],
};

/// Descriptor for `ListTodosRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTodosRequestDescriptor = $convert.base64Decode(
    'ChBMaXN0VG9kb3NSZXF1ZXN0EhQKBWxpbWl0GAEgASgFUgVsaW1pdBIWCgZvZmZzZXQYAiABKA'
    'VSBm9mZnNldA==');

@$core.Deprecated('Use listTodosResponseDescriptor instead')
const ListTodosResponse$json = {
  '1': 'ListTodosResponse',
  '2': [
    {'1': 'count', '3': 1, '4': 1, '5': 5, '10': 'count'},
    {'1': 'data', '3': 2, '4': 3, '5': 11, '6': '.greet.Todo', '10': 'data'},
  ],
};

/// Descriptor for `ListTodosResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTodosResponseDescriptor = $convert.base64Decode(
    'ChFMaXN0VG9kb3NSZXNwb25zZRIUCgVjb3VudBgBIAEoBVIFY291bnQSHwoEZGF0YRgCIAMoCz'
    'ILLmdyZWV0LlRvZG9SBGRhdGE=');

@$core.Deprecated('Use resetTodoRequestDescriptor instead')
const ResetTodoRequest$json = {
  '1': 'ResetTodoRequest',
  '2': [
    {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '10': 'todoId'},
  ],
};

/// Descriptor for `ResetTodoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resetTodoRequestDescriptor = $convert.base64Decode(
    'ChBSZXNldFRvZG9SZXF1ZXN0EhcKB3RvZG9faWQYASABKAlSBnRvZG9JZA==');

@$core.Deprecated('Use resetTodoResponseDescriptor instead')
const ResetTodoResponse$json = {
  '1': 'ResetTodoResponse',
};

/// Descriptor for `ResetTodoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List resetTodoResponseDescriptor = $convert.base64Decode(
    'ChFSZXNldFRvZG9SZXNwb25zZQ==');

@$core.Deprecated('Use todoDescriptor instead')
const Todo$json = {
  '1': 'Todo',
  '2': [
    {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '10': 'todoId'},
    {'1': 'version', '3': 2, '4': 1, '5': 5, '10': 'version'},
    {'1': 'note', '3': 3, '4': 1, '5': 9, '10': 'note'},
    {'1': 'due', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'due'},
    {'1': 'status', '3': 5, '4': 1, '5': 14, '6': '.greet.TodoStatus', '10': 'status'},
  ],
};

/// Descriptor for `Todo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List todoDescriptor = $convert.base64Decode(
    'CgRUb2RvEhcKB3RvZG9faWQYASABKAlSBnRvZG9JZBIYCgd2ZXJzaW9uGAIgASgFUgd2ZXJzaW'
    '9uEhIKBG5vdGUYAyABKAlSBG5vdGUSLAoDZHVlGAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRp'
    'bWVzdGFtcFIDZHVlEikKBnN0YXR1cxgFIAEoDjIRLmdyZWV0LlRvZG9TdGF0dXNSBnN0YXR1cw'
    '==');

