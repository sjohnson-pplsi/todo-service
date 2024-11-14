using MongoDB.Bson.Serialization;
using TodoCsService.Features.Base;
using TodoCsService.Features.TodoList.Domain.Entity;
using TodoCsService.Features.TodoList.Domain.Value;

namespace TodoCsService.Features.TodoList.Infrastructure;

public record TodoModel(
    string Id,
    int Version,
    string Note,
    DateTime? Due,
    string Status
)
{
    public Todo ToEntity()
    {
        return new Todo(
            TodoId.Parse(Id),
            new AggregateVersion(Version),
            new TodoNote(Note),
            new TodoDue(Due),
            ParseStatus
        );
    }

    public TodoModel Increment() => this with { Version = Version + 1 };

    private TodoStatus ParseStatus => BaseEnum.Parse(Status, TodoStatus.incomplete);

    public static void Register()
    {
        BsonClassMap.RegisterClassMap<TodoModel>(cm =>
        {
            cm.AutoMap();
            cm.MapIdMember(c => c.Id);
        });
    }
}


public static class TodoModelExtensions
{
    public static TodoModel ToModel(this Todo todo)
    {
        return new TodoModel(
            todo.Id.ToString(),
            todo.Version.Value,
            todo.Note.Value,
            todo.Due.Value,
            todo.Status.ToString()
        );
    }
}