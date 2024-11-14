using Grpc.Core;
using TodoCsLib;
using TodoCsService.Features.Todo.Domain.Entity;
using TodoCsService.Features.Todo.Domain.Repositories;

namespace TodoCsService.Services;

public class TodoService(ILogger<TodoService> logger, ITodoRepository todoRepository) : TodoCsLib.TodoService.TodoServiceBase
{
    private readonly ILogger<TodoService> _logger = logger;
    private readonly ITodoRepository _todoRepository = todoRepository;

    public override Task<CompleteTodoResponse> CompleteTodo(CompleteTodoRequest request, ServerCallContext context)
    {
        return base.CompleteTodo(request, context);
    }

    public override async Task<CreateTodoResponse> CreateTodo(CreateTodoRequest request, ServerCallContext context)
    {
        var id = await _todoRepository.InsertTodo(new Features.Todo.Domain.Entity.Todo(
            new Features.Todo.Domain.Value.TodoId(Guid.NewGuid()),
            new Features.Base.AggregateVersion(0),
            new Features.Todo.Domain.Value.TodoNote(""),
            new Features.Todo.Domain.Value.TodoDue(DateOnly.FromDateTime(DateTime.Now)),
            Features.Todo.Domain.Value.TodoStatus.incomplete
        ));
        return new CreateTodoResponse
        {
            TodoId = id.Value.ToString(),
        };
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
