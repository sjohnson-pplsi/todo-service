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

import 'google/protobuf/timestamp.pb.dart' as $1;
import 'todo.pbenum.dart';

export 'todo.pbenum.dart';

class ChangeNoteRequest extends $pb.GeneratedMessage {
  factory ChangeNoteRequest({
    $core.String? todoId,
    $core.String? note,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    if (note != null) {
      $result.note = note;
    }
    return $result;
  }
  ChangeNoteRequest._() : super();
  factory ChangeNoteRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ChangeNoteRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ChangeNoteRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..aOS(2, _omitFieldNames ? '' : 'note')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ChangeNoteRequest clone() => ChangeNoteRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ChangeNoteRequest copyWith(void Function(ChangeNoteRequest) updates) => super.copyWith((message) => updates(message as ChangeNoteRequest)) as ChangeNoteRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ChangeNoteRequest create() => ChangeNoteRequest._();
  ChangeNoteRequest createEmptyInstance() => create();
  static $pb.PbList<ChangeNoteRequest> createRepeated() => $pb.PbList<ChangeNoteRequest>();
  @$core.pragma('dart2js:noInline')
  static ChangeNoteRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ChangeNoteRequest>(create);
  static ChangeNoteRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get note => $_getSZ(1);
  @$pb.TagNumber(2)
  set note($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNote() => $_has(1);
  @$pb.TagNumber(2)
  void clearNote() => clearField(2);
}

class ChangeNoteResponse extends $pb.GeneratedMessage {
  factory ChangeNoteResponse() => create();
  ChangeNoteResponse._() : super();
  factory ChangeNoteResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ChangeNoteResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ChangeNoteResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ChangeNoteResponse clone() => ChangeNoteResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ChangeNoteResponse copyWith(void Function(ChangeNoteResponse) updates) => super.copyWith((message) => updates(message as ChangeNoteResponse)) as ChangeNoteResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ChangeNoteResponse create() => ChangeNoteResponse._();
  ChangeNoteResponse createEmptyInstance() => create();
  static $pb.PbList<ChangeNoteResponse> createRepeated() => $pb.PbList<ChangeNoteResponse>();
  @$core.pragma('dart2js:noInline')
  static ChangeNoteResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ChangeNoteResponse>(create);
  static ChangeNoteResponse? _defaultInstance;
}

class CompleteTodoRequest extends $pb.GeneratedMessage {
  factory CompleteTodoRequest({
    $core.String? todoId,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    return $result;
  }
  CompleteTodoRequest._() : super();
  factory CompleteTodoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CompleteTodoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CompleteTodoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CompleteTodoRequest clone() => CompleteTodoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CompleteTodoRequest copyWith(void Function(CompleteTodoRequest) updates) => super.copyWith((message) => updates(message as CompleteTodoRequest)) as CompleteTodoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CompleteTodoRequest create() => CompleteTodoRequest._();
  CompleteTodoRequest createEmptyInstance() => create();
  static $pb.PbList<CompleteTodoRequest> createRepeated() => $pb.PbList<CompleteTodoRequest>();
  @$core.pragma('dart2js:noInline')
  static CompleteTodoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CompleteTodoRequest>(create);
  static CompleteTodoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class CompleteTodoResponse extends $pb.GeneratedMessage {
  factory CompleteTodoResponse() => create();
  CompleteTodoResponse._() : super();
  factory CompleteTodoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CompleteTodoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CompleteTodoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CompleteTodoResponse clone() => CompleteTodoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CompleteTodoResponse copyWith(void Function(CompleteTodoResponse) updates) => super.copyWith((message) => updates(message as CompleteTodoResponse)) as CompleteTodoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CompleteTodoResponse create() => CompleteTodoResponse._();
  CompleteTodoResponse createEmptyInstance() => create();
  static $pb.PbList<CompleteTodoResponse> createRepeated() => $pb.PbList<CompleteTodoResponse>();
  @$core.pragma('dart2js:noInline')
  static CompleteTodoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CompleteTodoResponse>(create);
  static CompleteTodoResponse? _defaultInstance;
}

class CreateTodoRequest extends $pb.GeneratedMessage {
  factory CreateTodoRequest({
    $core.String? note,
    $1.Timestamp? due,
  }) {
    final $result = create();
    if (note != null) {
      $result.note = note;
    }
    if (due != null) {
      $result.due = due;
    }
    return $result;
  }
  CreateTodoRequest._() : super();
  factory CreateTodoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateTodoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTodoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'note')
    ..aOM<$1.Timestamp>(2, _omitFieldNames ? '' : 'due', subBuilder: $1.Timestamp.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateTodoRequest clone() => CreateTodoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateTodoRequest copyWith(void Function(CreateTodoRequest) updates) => super.copyWith((message) => updates(message as CreateTodoRequest)) as CreateTodoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTodoRequest create() => CreateTodoRequest._();
  CreateTodoRequest createEmptyInstance() => create();
  static $pb.PbList<CreateTodoRequest> createRepeated() => $pb.PbList<CreateTodoRequest>();
  @$core.pragma('dart2js:noInline')
  static CreateTodoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTodoRequest>(create);
  static CreateTodoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get note => $_getSZ(0);
  @$pb.TagNumber(1)
  set note($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasNote() => $_has(0);
  @$pb.TagNumber(1)
  void clearNote() => clearField(1);

  @$pb.TagNumber(2)
  $1.Timestamp get due => $_getN(1);
  @$pb.TagNumber(2)
  set due($1.Timestamp v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDue() => $_has(1);
  @$pb.TagNumber(2)
  void clearDue() => clearField(2);
  @$pb.TagNumber(2)
  $1.Timestamp ensureDue() => $_ensure(1);
}

class CreateTodoResponse extends $pb.GeneratedMessage {
  factory CreateTodoResponse({
    $core.String? todoId,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    return $result;
  }
  CreateTodoResponse._() : super();
  factory CreateTodoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateTodoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateTodoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateTodoResponse clone() => CreateTodoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateTodoResponse copyWith(void Function(CreateTodoResponse) updates) => super.copyWith((message) => updates(message as CreateTodoResponse)) as CreateTodoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTodoResponse create() => CreateTodoResponse._();
  CreateTodoResponse createEmptyInstance() => create();
  static $pb.PbList<CreateTodoResponse> createRepeated() => $pb.PbList<CreateTodoResponse>();
  @$core.pragma('dart2js:noInline')
  static CreateTodoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTodoResponse>(create);
  static CreateTodoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class GetTodoRequest extends $pb.GeneratedMessage {
  factory GetTodoRequest({
    $core.String? todoId,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    return $result;
  }
  GetTodoRequest._() : super();
  factory GetTodoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetTodoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTodoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetTodoRequest clone() => GetTodoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetTodoRequest copyWith(void Function(GetTodoRequest) updates) => super.copyWith((message) => updates(message as GetTodoRequest)) as GetTodoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTodoRequest create() => GetTodoRequest._();
  GetTodoRequest createEmptyInstance() => create();
  static $pb.PbList<GetTodoRequest> createRepeated() => $pb.PbList<GetTodoRequest>();
  @$core.pragma('dart2js:noInline')
  static GetTodoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTodoRequest>(create);
  static GetTodoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class GetTodoResponse extends $pb.GeneratedMessage {
  factory GetTodoResponse({
    Todo? todo,
  }) {
    final $result = create();
    if (todo != null) {
      $result.todo = todo;
    }
    return $result;
  }
  GetTodoResponse._() : super();
  factory GetTodoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetTodoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetTodoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOM<Todo>(1, _omitFieldNames ? '' : 'todo', subBuilder: Todo.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetTodoResponse clone() => GetTodoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetTodoResponse copyWith(void Function(GetTodoResponse) updates) => super.copyWith((message) => updates(message as GetTodoResponse)) as GetTodoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetTodoResponse create() => GetTodoResponse._();
  GetTodoResponse createEmptyInstance() => create();
  static $pb.PbList<GetTodoResponse> createRepeated() => $pb.PbList<GetTodoResponse>();
  @$core.pragma('dart2js:noInline')
  static GetTodoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTodoResponse>(create);
  static GetTodoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Todo get todo => $_getN(0);
  @$pb.TagNumber(1)
  set todo(Todo v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodo() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodo() => clearField(1);
  @$pb.TagNumber(1)
  Todo ensureTodo() => $_ensure(0);
}

class ListTodosRequest extends $pb.GeneratedMessage {
  factory ListTodosRequest({
    $core.int? limit,
    $core.int? offset,
  }) {
    final $result = create();
    if (limit != null) {
      $result.limit = limit;
    }
    if (offset != null) {
      $result.offset = offset;
    }
    return $result;
  }
  ListTodosRequest._() : super();
  factory ListTodosRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListTodosRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTodosRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'limit', $pb.PbFieldType.O3)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'offset', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListTodosRequest clone() => ListTodosRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListTodosRequest copyWith(void Function(ListTodosRequest) updates) => super.copyWith((message) => updates(message as ListTodosRequest)) as ListTodosRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTodosRequest create() => ListTodosRequest._();
  ListTodosRequest createEmptyInstance() => create();
  static $pb.PbList<ListTodosRequest> createRepeated() => $pb.PbList<ListTodosRequest>();
  @$core.pragma('dart2js:noInline')
  static ListTodosRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTodosRequest>(create);
  static ListTodosRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get limit => $_getIZ(0);
  @$pb.TagNumber(1)
  set limit($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLimit() => $_has(0);
  @$pb.TagNumber(1)
  void clearLimit() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get offset => $_getIZ(1);
  @$pb.TagNumber(2)
  set offset($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOffset() => $_has(1);
  @$pb.TagNumber(2)
  void clearOffset() => clearField(2);
}

class ListTodosResponse extends $pb.GeneratedMessage {
  factory ListTodosResponse({
    $core.int? count,
    $core.Iterable<Todo>? data,
  }) {
    final $result = create();
    if (count != null) {
      $result.count = count;
    }
    if (data != null) {
      $result.data.addAll(data);
    }
    return $result;
  }
  ListTodosResponse._() : super();
  factory ListTodosResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListTodosResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ListTodosResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'count', $pb.PbFieldType.O3)
    ..pc<Todo>(2, _omitFieldNames ? '' : 'data', $pb.PbFieldType.PM, subBuilder: Todo.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListTodosResponse clone() => ListTodosResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListTodosResponse copyWith(void Function(ListTodosResponse) updates) => super.copyWith((message) => updates(message as ListTodosResponse)) as ListTodosResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTodosResponse create() => ListTodosResponse._();
  ListTodosResponse createEmptyInstance() => create();
  static $pb.PbList<ListTodosResponse> createRepeated() => $pb.PbList<ListTodosResponse>();
  @$core.pragma('dart2js:noInline')
  static ListTodosResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListTodosResponse>(create);
  static ListTodosResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get count => $_getIZ(0);
  @$pb.TagNumber(1)
  set count($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCount() => $_has(0);
  @$pb.TagNumber(1)
  void clearCount() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<Todo> get data => $_getList(1);
}

class ResetTodoRequest extends $pb.GeneratedMessage {
  factory ResetTodoRequest({
    $core.String? todoId,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    return $result;
  }
  ResetTodoRequest._() : super();
  factory ResetTodoRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResetTodoRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ResetTodoRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResetTodoRequest clone() => ResetTodoRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResetTodoRequest copyWith(void Function(ResetTodoRequest) updates) => super.copyWith((message) => updates(message as ResetTodoRequest)) as ResetTodoRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ResetTodoRequest create() => ResetTodoRequest._();
  ResetTodoRequest createEmptyInstance() => create();
  static $pb.PbList<ResetTodoRequest> createRepeated() => $pb.PbList<ResetTodoRequest>();
  @$core.pragma('dart2js:noInline')
  static ResetTodoRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResetTodoRequest>(create);
  static ResetTodoRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class ResetTodoResponse extends $pb.GeneratedMessage {
  factory ResetTodoResponse() => create();
  ResetTodoResponse._() : super();
  factory ResetTodoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ResetTodoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ResetTodoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ResetTodoResponse clone() => ResetTodoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ResetTodoResponse copyWith(void Function(ResetTodoResponse) updates) => super.copyWith((message) => updates(message as ResetTodoResponse)) as ResetTodoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ResetTodoResponse create() => ResetTodoResponse._();
  ResetTodoResponse createEmptyInstance() => create();
  static $pb.PbList<ResetTodoResponse> createRepeated() => $pb.PbList<ResetTodoResponse>();
  @$core.pragma('dart2js:noInline')
  static ResetTodoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ResetTodoResponse>(create);
  static ResetTodoResponse? _defaultInstance;
}

class Todo extends $pb.GeneratedMessage {
  factory Todo({
    $core.String? todoId,
    $core.int? version,
    $core.String? note,
    $1.Timestamp? due,
    TodoStatus? status,
  }) {
    final $result = create();
    if (todoId != null) {
      $result.todoId = todoId;
    }
    if (version != null) {
      $result.version = version;
    }
    if (note != null) {
      $result.note = note;
    }
    if (due != null) {
      $result.due = due;
    }
    if (status != null) {
      $result.status = status;
    }
    return $result;
  }
  Todo._() : super();
  factory Todo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Todo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Todo', package: const $pb.PackageName(_omitMessageNames ? '' : 'greet'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'todoId')
    ..a<$core.int>(2, _omitFieldNames ? '' : 'version', $pb.PbFieldType.O3)
    ..aOS(3, _omitFieldNames ? '' : 'note')
    ..aOM<$1.Timestamp>(4, _omitFieldNames ? '' : 'due', subBuilder: $1.Timestamp.create)
    ..e<TodoStatus>(5, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: TodoStatus.TODO_STATUS_UNSPECIFIED, valueOf: TodoStatus.valueOf, enumValues: TodoStatus.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Todo clone() => Todo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Todo copyWith(void Function(Todo) updates) => super.copyWith((message) => updates(message as Todo)) as Todo;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Todo create() => Todo._();
  Todo createEmptyInstance() => create();
  static $pb.PbList<Todo> createRepeated() => $pb.PbList<Todo>();
  @$core.pragma('dart2js:noInline')
  static Todo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Todo>(create);
  static Todo? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get version => $_getIZ(1);
  @$pb.TagNumber(2)
  set version($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasVersion() => $_has(1);
  @$pb.TagNumber(2)
  void clearVersion() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get note => $_getSZ(2);
  @$pb.TagNumber(3)
  set note($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNote() => $_has(2);
  @$pb.TagNumber(3)
  void clearNote() => clearField(3);

  @$pb.TagNumber(4)
  $1.Timestamp get due => $_getN(3);
  @$pb.TagNumber(4)
  set due($1.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasDue() => $_has(3);
  @$pb.TagNumber(4)
  void clearDue() => clearField(4);
  @$pb.TagNumber(4)
  $1.Timestamp ensureDue() => $_ensure(3);

  @$pb.TagNumber(5)
  TodoStatus get status => $_getN(4);
  @$pb.TagNumber(5)
  set status(TodoStatus v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(4);
  @$pb.TagNumber(5)
  void clearStatus() => clearField(5);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
