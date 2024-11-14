// TODO: This is likely only needed when linking locally
import { credentials } from "todo-ts-lib/node_modules/@grpc/grpc-js";
import { TodoServiceClient } from "todo-ts-lib/src/generated/todo_grpc_pb";
import { ListTodosRequest } from "todo-ts-lib/src/generated/todo_pb";

import { unaryCallToPromise } from "./unaryCallToPromise";

function createGreeterClient(host: string) {
  const todoClient = new TodoServiceClient(host, credentials.createInsecure());
  return {
    listTodos: unaryCallToPromise(todoClient.listTodos.bind(todoClient)),
  };
}

export class TodoService {
  private todoClient: ReturnType<typeof createGreeterClient>;

  constructor() {
    this.todoClient = createGreeterClient("localhost:9001");
  }

  async listTodos(limit: number, offset: number) {
    const request = new ListTodosRequest();
    request.setLimit(limit);
    request.setOffset(offset);
    const response = await this.todoClient.listTodos(request);
    return {
      count: response.getCount(),
      data: response.getDataList().map((value) => {
        return {
          id: value.getTodoId(),
          version: value.getVersion(),
          due: value.getDue(),
          note: value.getNote(),
          status: value.getStatus(),
        };
      }),
    };
  }
}

export const todoService = new TodoService();
