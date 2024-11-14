// package: greet
// file: todo.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class ChangeNoteRequest extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): ChangeNoteRequest;
    getNote(): string;
    setNote(value: string): ChangeNoteRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChangeNoteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ChangeNoteRequest): ChangeNoteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChangeNoteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChangeNoteRequest;
    static deserializeBinaryFromReader(message: ChangeNoteRequest, reader: jspb.BinaryReader): ChangeNoteRequest;
}

export namespace ChangeNoteRequest {
    export type AsObject = {
        todoId: string,
        note: string,
    }
}

export class ChangeNoteResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ChangeNoteResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ChangeNoteResponse): ChangeNoteResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ChangeNoteResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ChangeNoteResponse;
    static deserializeBinaryFromReader(message: ChangeNoteResponse, reader: jspb.BinaryReader): ChangeNoteResponse;
}

export namespace ChangeNoteResponse {
    export type AsObject = {
    }
}

export class CompleteTodoRequest extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): CompleteTodoRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CompleteTodoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CompleteTodoRequest): CompleteTodoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CompleteTodoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CompleteTodoRequest;
    static deserializeBinaryFromReader(message: CompleteTodoRequest, reader: jspb.BinaryReader): CompleteTodoRequest;
}

export namespace CompleteTodoRequest {
    export type AsObject = {
        todoId: string,
    }
}

export class CompleteTodoResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CompleteTodoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CompleteTodoResponse): CompleteTodoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CompleteTodoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CompleteTodoResponse;
    static deserializeBinaryFromReader(message: CompleteTodoResponse, reader: jspb.BinaryReader): CompleteTodoResponse;
}

export namespace CompleteTodoResponse {
    export type AsObject = {
    }
}

export class CreateTodoRequest extends jspb.Message { 
    getNote(): string;
    setNote(value: string): CreateTodoRequest;

    hasDue(): boolean;
    clearDue(): void;
    getDue(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setDue(value?: google_protobuf_timestamp_pb.Timestamp): CreateTodoRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateTodoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateTodoRequest): CreateTodoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateTodoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateTodoRequest;
    static deserializeBinaryFromReader(message: CreateTodoRequest, reader: jspb.BinaryReader): CreateTodoRequest;
}

export namespace CreateTodoRequest {
    export type AsObject = {
        note: string,
        due?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class CreateTodoResponse extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): CreateTodoResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateTodoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateTodoResponse): CreateTodoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateTodoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateTodoResponse;
    static deserializeBinaryFromReader(message: CreateTodoResponse, reader: jspb.BinaryReader): CreateTodoResponse;
}

export namespace CreateTodoResponse {
    export type AsObject = {
        todoId: string,
    }
}

export class GetTodoRequest extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): GetTodoRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTodoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetTodoRequest): GetTodoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTodoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTodoRequest;
    static deserializeBinaryFromReader(message: GetTodoRequest, reader: jspb.BinaryReader): GetTodoRequest;
}

export namespace GetTodoRequest {
    export type AsObject = {
        todoId: string,
    }
}

export class GetTodoResponse extends jspb.Message { 

    hasTodo(): boolean;
    clearTodo(): void;
    getTodo(): Todo | undefined;
    setTodo(value?: Todo): GetTodoResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTodoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetTodoResponse): GetTodoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTodoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTodoResponse;
    static deserializeBinaryFromReader(message: GetTodoResponse, reader: jspb.BinaryReader): GetTodoResponse;
}

export namespace GetTodoResponse {
    export type AsObject = {
        todo?: Todo.AsObject,
    }
}

export class ListTodosRequest extends jspb.Message { 
    getLimit(): number;
    setLimit(value: number): ListTodosRequest;
    getOffset(): number;
    setOffset(value: number): ListTodosRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListTodosRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListTodosRequest): ListTodosRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListTodosRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListTodosRequest;
    static deserializeBinaryFromReader(message: ListTodosRequest, reader: jspb.BinaryReader): ListTodosRequest;
}

export namespace ListTodosRequest {
    export type AsObject = {
        limit: number,
        offset: number,
    }
}

export class ListTodosResponse extends jspb.Message { 
    getCount(): number;
    setCount(value: number): ListTodosResponse;
    clearDataList(): void;
    getDataList(): Array<Todo>;
    setDataList(value: Array<Todo>): ListTodosResponse;
    addData(value?: Todo, index?: number): Todo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListTodosResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListTodosResponse): ListTodosResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListTodosResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListTodosResponse;
    static deserializeBinaryFromReader(message: ListTodosResponse, reader: jspb.BinaryReader): ListTodosResponse;
}

export namespace ListTodosResponse {
    export type AsObject = {
        count: number,
        dataList: Array<Todo.AsObject>,
    }
}

export class ResetTodoRequest extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): ResetTodoRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResetTodoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ResetTodoRequest): ResetTodoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResetTodoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResetTodoRequest;
    static deserializeBinaryFromReader(message: ResetTodoRequest, reader: jspb.BinaryReader): ResetTodoRequest;
}

export namespace ResetTodoRequest {
    export type AsObject = {
        todoId: string,
    }
}

export class ResetTodoResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResetTodoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ResetTodoResponse): ResetTodoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResetTodoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResetTodoResponse;
    static deserializeBinaryFromReader(message: ResetTodoResponse, reader: jspb.BinaryReader): ResetTodoResponse;
}

export namespace ResetTodoResponse {
    export type AsObject = {
    }
}

export class Todo extends jspb.Message { 
    getTodoId(): string;
    setTodoId(value: string): Todo;
    getVersion(): number;
    setVersion(value: number): Todo;
    getNote(): string;
    setNote(value: string): Todo;

    hasDue(): boolean;
    clearDue(): void;
    getDue(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setDue(value?: google_protobuf_timestamp_pb.Timestamp): Todo;
    getStatus(): TodoStatus;
    setStatus(value: TodoStatus): Todo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Todo.AsObject;
    static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Todo;
    static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
    export type AsObject = {
        todoId: string,
        version: number,
        note: string,
        due?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        status: TodoStatus,
    }
}

export enum TodoStatus {
    TODO_STATUS_UNSPECIFIED = 0,
    TODO_STATUS_INCOMPLETE = 1,
    TODO_STATUS_COMPLETE = 2,
}
