# Example gRPC Service

This is an example gRPC implementation.  We have provided an example gRPC definition, and generated client and server stubs for several languages: **C#**, **Dart**, **Go**, and **TypeScript**.  We have also provided an two full service implementations, in **C#** and **Go**, and one client implementation, in **Next.js**.

## gRPC Definitions

The gRPC definitions are made up of `.proto` files in the `/definitions` directory.  These are used to generate client and server stubs, which are stored in shared libraries.

### VS Code Workspace

To open in VSCode, the definitions can be loaded as a whole by opening `/definitions.code-workspace`.

### Build

To generate the client and server stubs, run:

``` bash
cd definitions
make
```

This will generate code in the `/modules` folder.

| Language   | Folder                   | Notes                                     |
| ---------- | ------------------------ | ----------------------------------------- |
| C#         | `/modules/TodoCsLib`     | Generated by dotnet tools, not `makefile` |
| Dart       | `/modules/todo-dart-lib` |                                           |
| Go         | `/modules/todo-go-lib`   |                                           |
| TypeScript | `/modules/todo-ts-lib`   |                                           |


## C# Service

The C# implementation is located in the `/services/TodoCsService` directory.  The solution file is `/todo-service.sln`.

The project is designed using Domain Driven Design.  We have one bounded context `TodoCsService.Features.TodoList`.  This relies on a set of base classes in `TodoCsService.Feature.Base`.

Inside our context, we have our domain namespace, made up of four namespaces:

* `TodoCsService.Features.TodoList.Entities` - Entities
* `TodoCsService.Features.TodoList.Repositories` - Repository interfaces
* `TodoCsService.Features.TodoList.Services` - Application Services
* `TodoCsService.Features.TodoList.Values` - Value Objects

We also have a Controller layer in `TodoCsService.Features.TodoList.Controllers`, which implements the gRPC interface.  This class calls the Application Service directly.

We then implement our Repositories in `TodoCsService.Features.TodoList.Repositories`.  In our case, the Repository is simple a MongoDb wrapper, and a mapper to a Data Model for storage.

### VS Code Workspace

To open in VSCode, the service and definitions can be loaded by opening `/todo-cs-service.code-workspace`.

### Build and Run

Run the service at `http://localhost:9001` with the following command:

``` bash
cd services/TodoCsService
dotnet run
```

## Go Service

The Go implementation is located in the `/services/todo-go-service` directory.

The project is designed using Domain Driven Design.  We have one bounded context `".../todo-go-service/feature/todo"`.  This relies on a set of base classes in `".../todo-go-service/feature/base`.

Inside our context, we have our domain namespace, made up of three packages:

* `".../todo-go-service/feature/todo/domain/entity"` - Entities
* `".../todo-go-service/feature/todo/domain/service"` - Application Services
* `".../todo-go-service/feature/todo/domain/value"` - Value Objects

We also have a Controller layer in `.../todo-go-service/feature/todo/controller`, which implements the gRPC interface.  This class calls the Application Service directly.

We then implement our Repositories in `.../todo-go-service/feature/todo/repository`.  In our case, the Repository is simple a MongoDb wrapper, and a mapper to a Data Model for storage.

### VS Code Workspace

To open in VSCode, the service and definitions can be loaded by opening `/todo-go-service.code-workspace`.

### Build and Run

Run the service at `http://localhost:9001` with the following command:

``` bash
cd services/todo-go-service
go run ./cmd
```

## Next.js Client

The Next.js Client implementation is located in the `/clients/todo-app` directory.

### VS Code Workspace

To open in VSCode, the client and definitions can be loaded by opening `/todo-app.code-workspace`.

### Build and Run

Run the client at `http://localhost:3000` with the following command:

``` bash
cd clients/todo-app
npm run dev
```
