// TODO: This is likely only needed when linking locally
import { credentials } from "todo-ts-lib/node_modules/@grpc/grpc-js";
import { GreeterClient } from "todo-ts-lib/src/generated/greet_grpc_pb";
import { HelloRequest } from "todo-ts-lib/src/generated/greet_pb";

import { unaryCallToPromise } from "./unaryCallToPromise";

function createGreeterClient(host: string) {
  const greeterClient = new GreeterClient(host, credentials.createInsecure());
  return {
    sayHello: unaryCallToPromise(greeterClient.sayHello.bind(greeterClient)),
  };
}

export class TodoService {
  greeterClient: ReturnType<typeof createGreeterClient>;

  constructor() {
    this.greeterClient = createGreeterClient("localhost:5196");
  }

  async sayHello(name: string) {
    const request = new HelloRequest();
    request.setName(name);
    const response = await this.greeterClient.sayHello(request);
    return response.getMessage();
  }
}

export const todoService = new TodoService();
