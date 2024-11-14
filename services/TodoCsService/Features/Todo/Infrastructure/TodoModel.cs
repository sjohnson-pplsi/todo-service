using TodoCsService.Features.Base;
using TodoCsService.Features.Todo.Domain.Value;

namespace TodoCsService.Features.Todo.Infrastructure;

public record TodoModel(
    string Id,
    int Version,
    string Note,
    DateOnly Due,
    string Status
)
{
    public Domain.Entity.Todo ToEntity()
    {
        return new Domain.Entity.Todo(
            new TodoId(Guid.Parse(Id)),
            new AggregateVersion(Version),
            new TodoNote(Note),
            new TodoDue(Due),
            Enum.Parse<TodoStatus>(Status)
        );
    }
}


public static class TodoModelExtensions
{
    public static TodoModel ToModel(this Domain.Entity.Todo todo)
    {
        return new TodoModel(
            todo.Id.Value.ToString(),
            todo.Version.Value,
            todo.Note.Value,
            todo.Due.Value,
            todo.Status.ToString()
        );
    }
}