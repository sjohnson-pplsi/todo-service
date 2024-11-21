using TodoCsService.Features.TodoList.Domain.Entities;
using TodoCsService.Features.TodoList.Domain.Repositories;
using TodoCsService.Features.TodoList.Domain.Values;

namespace TodoCsService.Features.TodoList.Domain.Services;

public class TodoService(ITodoRepository todoRepository)
{
    private readonly ITodoRepository _todoRepository = todoRepository;

    public async Task CompleteTodo(CompleteTodoCommand cmd, CancellationToken token = default)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id, token);
        todo.Complete();
        await _todoRepository.ReplaceTodo(todo, token);
    }

    public async Task ChangeNote(ChangeNoteCommand cmd, CancellationToken token = default)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id, token);
        todo.ChangeNote(cmd.Note);
        await _todoRepository.ReplaceTodo(todo, token);
    }

    public async Task<TodoId> CreateTodo(CreateTodoCommand cmd, CancellationToken token = default)
    {
        var todo = new Todo(
            new TodoId(Guid.NewGuid()),
            new Base.AggregateVersion(0),
            cmd.Note,
            cmd.Due,
            TodoStatus.incomplete
        );
        var result = await _todoRepository.InsertTodo(todo, token);
        return result;
    }

    public async Task ResetTodo(ResetTodocommand cmd, CancellationToken token = default)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id, token);
        todo.Reset();
        await _todoRepository.ReplaceTodo(todo, token);
    }

    public async Task<Todo> GetTodo(GetTodoQuery qry, CancellationToken token = default)
    {
        var todo = await _todoRepository.GetTodo(qry.Id, token);
        return todo;
    }

    public async Task<ICollection<Todo>> ListTodos(ListTodosQuery qry, CancellationToken token = default)
    {
        var todos = await _todoRepository.ListTodos(qry.Limit, qry.Offset, token);
        return todos;
    }
}

public record CompleteTodoCommand(
    TodoId Id
);

public record CreateTodoCommand(
    TodoNote Note,
    TodoDue Due
);

public record ChangeNoteCommand(
    TodoId Id,
    TodoNote Note
);

public record ResetTodocommand(
    TodoId Id
);

public record GetTodoQuery(
    TodoId Id
);

public record ListTodosQuery(
    int Limit,
    int Offset
);
