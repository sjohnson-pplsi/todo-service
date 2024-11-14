//
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class TodoStatus extends $pb.ProtobufEnum {
  static const TodoStatus TODO_STATUS_UNSPECIFIED = TodoStatus._(0, _omitEnumNames ? '' : 'TODO_STATUS_UNSPECIFIED');
  static const TodoStatus TODO_STATUS_INCOMPLETE = TodoStatus._(1, _omitEnumNames ? '' : 'TODO_STATUS_INCOMPLETE');
  static const TodoStatus TODO_STATUS_COMPLETE = TodoStatus._(2, _omitEnumNames ? '' : 'TODO_STATUS_COMPLETE');

  static const $core.List<TodoStatus> values = <TodoStatus> [
    TODO_STATUS_UNSPECIFIED,
    TODO_STATUS_INCOMPLETE,
    TODO_STATUS_COMPLETE,
  ];

  static final $core.Map<$core.int, TodoStatus> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TodoStatus? valueOf($core.int value) => _byValue[value];

  const TodoStatus._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
