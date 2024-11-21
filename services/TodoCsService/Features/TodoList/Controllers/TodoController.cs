using Grpc.Core;
using TodoCsLib;
using TodoCsService.Features.TodoList.Domain.Services;
using TodoCsService.Features.TodoList.Domain.Values;

namespace TodoCsService.Features.TodoList.Controllers;

public class TodoController(ILogger<TodoController> logger, Domain.Services.TodoService todoService) : TodoCsLib.TodoService.TodoServiceBase
{
    private readonly ILogger<TodoController> _logger = logger;
    private readonly Domain.Services.TodoService _todoService = todoService;

    public override async Task<CompleteTodoResponse> CompleteTodo(CompleteTodoRequest request, ServerCallContext context)
    {
        await _todoService.CompleteTodo(
            new CompleteTodoCommand(TodoId.Parse(request.TodoId)),
            context.CancellationToken);
        return new CompleteTodoResponse { };
    }

    public override async Task<CreateTodoResponse> CreateTodo(CreateTodoRequest request, ServerCallContext context)
    {
        var id = await _todoService.CreateTodo(new CreateTodoCommand(
            new TodoNote(request.Note),
            new TodoDue(request.Due?.ToDateTime())),
            context.CancellationToken);
        return new CreateTodoResponse
        {
            TodoId = id.Value.ToString(),
        };
    }

    public override async Task<GetTodoResponse> GetTodo(GetTodoRequest request, ServerCallContext context)
    {
        var todo = await _todoService.GetTodo(
            new GetTodoQuery(TodoId.Parse(request.TodoId)),
            context.CancellationToken);
        return new GetTodoResponse
        {
            Todo = todo.ToDto(),
        };
    }

    public override async Task<ListTodosResponse> ListTodos(ListTodosRequest request, ServerCallContext context)
    {
        var list = await _todoService.ListTodos(
            new ListTodosQuery(request.Limit, request.Offset),
            context.CancellationToken);
        var response = new ListTodosResponse
        {
            Count = list.Count,
        };
        response.Data.AddRange(list.Select(t => t.ToDto()));
        return response;
    }

    public override async Task<ResetTodoResponse> ResetTodo(ResetTodoRequest request, ServerCallContext context)
    {
        await _todoService.ResetTodo(
            new ResetTodocommand(TodoId.Parse(request.TodoId)),
            context.CancellationToken);
        return new ResetTodoResponse { };
    }
}
