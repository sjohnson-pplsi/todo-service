using Grpc.Core;
using TodoCsLib;

namespace TodoCsService.Services;

public class TodoService(ILogger<TodoService> logger) : TodoCsLib.TodoService.TodoServiceBase
{
    private readonly ILogger<TodoService> _logger = logger;

    public override Task<CompleteTodoResponse> CompleteTodo(CompleteTodoRequest request, ServerCallContext context)
    {
        return base.CompleteTodo(request, context);
    }

    public override Task<CreateTodoResponse> CreateTodo(CreateTodoRequest request, ServerCallContext context)
    {
        return Task.FromResult(new CreateTodoResponse
        {
            TodoId = ""
        });
    }

    public override Task<GetTodoResponse> GetTodo(GetTodoRequest request, ServerCallContext context)
    {
        return base.GetTodo(request, context);
    }

    public override Task<ListTodosResponse> ListTodos(ListTodosRequest request, ServerCallContext context)
    {
        return base.ListTodos(request, context);
    }

    public override Task<ResetTodoResponse> ResetTodo(ResetTodoRequest request, ServerCallContext context)
    {
        return base.ResetTodo(request, context);
    }
}
