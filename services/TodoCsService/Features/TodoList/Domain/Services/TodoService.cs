using TodoCsService.Features.TodoList.Domain.Entity;
using TodoCsService.Features.TodoList.Domain.Repositories;
using TodoCsService.Features.TodoList.Domain.Value;

namespace TodoCsService.Features.TodoList.Domain.Services;

public class TodoService(ITodoRepository todoRepository)
{
    private readonly ITodoRepository _todoRepository = todoRepository;

    public async Task CompleteTodo(CompleteTodoCommand cmd)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id);
        todo = todo.Complete();
        await _todoRepository.ReplaceTodo(todo);
    }

    public async Task ChangeNote(ChangeNoteCommand cmd)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id);
        todo = todo.ChangeNote(cmd.Note);
        await _todoRepository.ReplaceTodo(todo);
    }

    public async Task<TodoId> CreateTodo(CreateTodoCommand cmd)
    {
        var todo = new Todo(
            new TodoId(Guid.NewGuid()),
            new Base.AggregateVersion(0),
            cmd.Note,
            cmd.Due,
            TodoStatus.incomplete
        );
        var result = await _todoRepository.InsertTodo(todo);
        return result;
    }

    public async Task ResetTodo(ResetTodocommand cmd)
    {
        var todo = await _todoRepository.GetTodo(cmd.Id);
        todo = todo.Reset();
        await _todoRepository.ReplaceTodo(todo);
    }

    public async Task<Todo> GetTodo(GetTodoQuery qry)
    {
        var todo = await _todoRepository.GetTodo(qry.Id);
        return todo;
    }

    public async Task<ICollection<Todo>> ListTodos(ListTodosQuery qry)
    {
        var todos = await _todoRepository.ListTodos(qry.Limit, qry.Offset);
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
