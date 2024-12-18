// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.1
// source: todo.proto

package togo_go_lib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TodoStatus int32

const (
	TodoStatus_TODO_STATUS_UNSPECIFIED TodoStatus = 0
	TodoStatus_TODO_STATUS_INCOMPLETE  TodoStatus = 1
	TodoStatus_TODO_STATUS_COMPLETE    TodoStatus = 2
)

// Enum value maps for TodoStatus.
var (
	TodoStatus_name = map[int32]string{
		0: "TODO_STATUS_UNSPECIFIED",
		1: "TODO_STATUS_INCOMPLETE",
		2: "TODO_STATUS_COMPLETE",
	}
	TodoStatus_value = map[string]int32{
		"TODO_STATUS_UNSPECIFIED": 0,
		"TODO_STATUS_INCOMPLETE":  1,
		"TODO_STATUS_COMPLETE":    2,
	}
)

func (x TodoStatus) Enum() *TodoStatus {
	p := new(TodoStatus)
	*p = x
	return p
}

func (x TodoStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_proto_enumTypes[0].Descriptor()
}

func (TodoStatus) Type() protoreflect.EnumType {
	return &file_todo_proto_enumTypes[0]
}

func (x TodoStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoStatus.Descriptor instead.
func (TodoStatus) EnumDescriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

type ChangeNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
	Note   string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *ChangeNoteRequest) Reset() {
	*x = ChangeNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNoteRequest) ProtoMessage() {}

func (x *ChangeNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNoteRequest.ProtoReflect.Descriptor instead.
func (*ChangeNoteRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeNoteRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

func (x *ChangeNoteRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type ChangeNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeNoteResponse) Reset() {
	*x = ChangeNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNoteResponse) ProtoMessage() {}

func (x *ChangeNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNoteResponse.ProtoReflect.Descriptor instead.
func (*ChangeNoteResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

type CompleteTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *CompleteTodoRequest) Reset() {
	*x = CompleteTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTodoRequest) ProtoMessage() {}

func (x *CompleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTodoRequest.ProtoReflect.Descriptor instead.
func (*CompleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteTodoRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type CompleteTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteTodoResponse) Reset() {
	*x = CompleteTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTodoResponse) ProtoMessage() {}

func (x *CompleteTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTodoResponse.ProtoReflect.Descriptor instead.
func (*CompleteTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note string                 `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Due  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=due,proto3" json:"due,omitempty"`
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTodoRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CreateTodoRequest) GetDue() *timestamppb.Timestamp {
	if x != nil {
		return x.Due
	}
	return nil
}

type CreateTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *CreateTodoResponse) Reset() {
	*x = CreateTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoResponse) ProtoMessage() {}

func (x *CreateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoResponse.ProtoReflect.Descriptor instead.
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTodoResponse) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type GetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{6}
}

func (x *GetTodoRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type GetTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetTodoResponse) Reset() {
	*x = GetTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoResponse) ProtoMessage() {}

func (x *GetTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoResponse.ProtoReflect.Descriptor instead.
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{7}
}

func (x *GetTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type ListTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTodosRequest) Reset() {
	*x = ListTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosRequest) ProtoMessage() {}

func (x *ListTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosRequest.ProtoReflect.Descriptor instead.
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{8}
}

func (x *ListTodosRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTodosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Data  []*Todo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListTodosResponse) Reset() {
	*x = ListTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosResponse) ProtoMessage() {}

func (x *ListTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosResponse.ProtoReflect.Descriptor instead.
func (*ListTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{9}
}

func (x *ListTodosResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListTodosResponse) GetData() []*Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResetTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
}

func (x *ResetTodoRequest) Reset() {
	*x = ResetTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetTodoRequest) ProtoMessage() {}

func (x *ResetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetTodoRequest.ProtoReflect.Descriptor instead.
func (*ResetTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{10}
}

func (x *ResetTodoRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type ResetTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetTodoResponse) Reset() {
	*x = ResetTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetTodoResponse) ProtoMessage() {}

func (x *ResetTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetTodoResponse.ProtoReflect.Descriptor instead.
func (*ResetTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{11}
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId  string                 `protobuf:"bytes,1,opt,name=todo_id,json=todoId,proto3" json:"todo_id,omitempty"`
	Version int32                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Note    string                 `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Due     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due,proto3" json:"due,omitempty"`
	Status  TodoStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=greet.TodoStatus" json:"status,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{12}
}

func (x *Todo) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

func (x *Todo) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Todo) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Todo) GetDue() *timestamppb.Timestamp {
	if x != nil {
		return x.Due
	}
	return nil
}

func (x *Todo) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_TODO_STATUS_UNSPECIFIED
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x13,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x03, 0x64, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x64, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x64, 0x6f, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4a, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x64, 0x6f, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x64, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x03, 0x64, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2a, 0x5f, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x4f, 0x44,
	0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x02, 0x32, 0x96, 0x03, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x18, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x15, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x42,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6a, 0x6f, 0x68, 0x6e, 0x73, 0x6f, 0x6e, 0x2d, 0x70, 0x70, 0x6c, 0x73,
	0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x6c,
	0x69, 0x62, 0xaa, 0x02, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x73, 0x4c, 0x69, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_todo_proto_goTypes = []interface{}{
	(TodoStatus)(0),               // 0: greet.TodoStatus
	(*ChangeNoteRequest)(nil),     // 1: greet.ChangeNoteRequest
	(*ChangeNoteResponse)(nil),    // 2: greet.ChangeNoteResponse
	(*CompleteTodoRequest)(nil),   // 3: greet.CompleteTodoRequest
	(*CompleteTodoResponse)(nil),  // 4: greet.CompleteTodoResponse
	(*CreateTodoRequest)(nil),     // 5: greet.CreateTodoRequest
	(*CreateTodoResponse)(nil),    // 6: greet.CreateTodoResponse
	(*GetTodoRequest)(nil),        // 7: greet.GetTodoRequest
	(*GetTodoResponse)(nil),       // 8: greet.GetTodoResponse
	(*ListTodosRequest)(nil),      // 9: greet.ListTodosRequest
	(*ListTodosResponse)(nil),     // 10: greet.ListTodosResponse
	(*ResetTodoRequest)(nil),      // 11: greet.ResetTodoRequest
	(*ResetTodoResponse)(nil),     // 12: greet.ResetTodoResponse
	(*Todo)(nil),                  // 13: greet.Todo
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_todo_proto_depIdxs = []int32{
	14, // 0: greet.CreateTodoRequest.due:type_name -> google.protobuf.Timestamp
	13, // 1: greet.GetTodoResponse.todo:type_name -> greet.Todo
	13, // 2: greet.ListTodosResponse.data:type_name -> greet.Todo
	14, // 3: greet.Todo.due:type_name -> google.protobuf.Timestamp
	0,  // 4: greet.Todo.status:type_name -> greet.TodoStatus
	1,  // 5: greet.TodoService.ChangeNote:input_type -> greet.ChangeNoteRequest
	3,  // 6: greet.TodoService.CompleteTodo:input_type -> greet.CompleteTodoRequest
	5,  // 7: greet.TodoService.CreateTodo:input_type -> greet.CreateTodoRequest
	7,  // 8: greet.TodoService.GetTodo:input_type -> greet.GetTodoRequest
	9,  // 9: greet.TodoService.ListTodos:input_type -> greet.ListTodosRequest
	11, // 10: greet.TodoService.ResetTodo:input_type -> greet.ResetTodoRequest
	2,  // 11: greet.TodoService.ChangeNote:output_type -> greet.ChangeNoteResponse
	4,  // 12: greet.TodoService.CompleteTodo:output_type -> greet.CompleteTodoResponse
	6,  // 13: greet.TodoService.CreateTodo:output_type -> greet.CreateTodoResponse
	8,  // 14: greet.TodoService.GetTodo:output_type -> greet.GetTodoResponse
	10, // 15: greet.TodoService.ListTodos:output_type -> greet.ListTodosResponse
	12, // 16: greet.TodoService.ResetTodo:output_type -> greet.ResetTodoResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNoteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNoteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetTodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_todo_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		EnumInfos:         file_todo_proto_enumTypes,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
