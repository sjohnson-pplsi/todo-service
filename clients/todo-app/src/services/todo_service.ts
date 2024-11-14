"use server";

// TODO: This is likely only needed when linking locally
import { credentials } from "todo-ts-lib/node_modules/@grpc/grpc-js";
import { TodoServiceClient } from "todo-ts-lib/src/generated/todo_grpc_pb";
import {
  ChangeNoteRequest,
  CompleteTodoRequest,
  CreateTodoRequest,
  ListTodosRequest,
  ResetTodoRequest,
  Todo as TodoDto,
  TodoStatus,
} from "todo-ts-lib/src/generated/todo_pb";

import { unaryCallToPromise } from "./unaryCallToPromise";

function createGreeterClient(host: string) {
  const todoClient = new TodoServiceClient(host, credentials.createInsecure());
  return {
    changeNote: unaryCallToPromise(todoClient.changeNote.bind(todoClient)),
    createTodo: unaryCallToPromise(todoClient.createTodo.bind(todoClient)),
    completeTodo: unaryCallToPromise(todoClient.completeTodo.bind(todoClient)),
    getTodo: unaryCallToPromise(todoClient.getTodo.bind(todoClient)),
    listTodos: unaryCallToPromise(todoClient.listTodos.bind(todoClient)),
    resetTodo: unaryCallToPromise(todoClient.resetTodo.bind(todoClient)),
  };
}

const todoClient = createGreeterClient("localhost:9001");

export async function listTodos(limit: number, offset: number) {
  const request = new ListTodosRequest();
  request.setLimit(limit);
  request.setOffset(offset);
  const response = await todoClient.listTodos(request);
  return {
    count: response.getCount(),
    data: response.getDataList().map(todoFromDto),
  };
}

export async function changeNote(id: string, note: string) {
  const request = new ChangeNoteRequest();
  request.setTodoId(id);
  request.setNote(note);
  await todoClient.changeNote(request);
}

export async function createTodo(todo: { due?: Date; note: string }) {
  const request = new CreateTodoRequest();
  request.setNote(todo.note);
  const response = await todoClient.createTodo(request);
  return response.getTodoId();
}

export async function completeTodo(id: string) {
  const request = new CompleteTodoRequest();
  request.setTodoId(id);
  await todoClient.completeTodo(request);
}

export async function resetTodo(id: string) {
  const request = new ResetTodoRequest();
  request.setTodoId(id);
  await todoClient.resetTodo(request);
}

function todoFromDto(dto: TodoDto) {
  return {
    id: dto.getTodoId(),
    version: dto.getVersion(),
    due: dto.getDue()?.toDate(),
    note: dto.getNote(),
    status: todoStatusFromDto(dto.getStatus()),
  };
}

export type Todo = ReturnType<typeof todoFromDto>;

function todoStatusFromDto(status: TodoStatus) {
  switch (status) {
    case TodoStatus.TODO_STATUS_COMPLETE:
      return true;
    case TodoStatus.TODO_STATUS_INCOMPLETE:
    case TodoStatus.TODO_STATUS_UNSPECIFIED:
    default:
      return false;
  }
}
