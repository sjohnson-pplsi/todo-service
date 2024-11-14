using Google.Protobuf.WellKnownTypes;
using TodoCsService.Features.Todo.Domain.Value;

namespace TodoCsService.Features.Todo.Controllers;

public static class TodoExtensions
{
    public static TodoCsLib.Todo ToDto(this Domain.Entity.Todo todo)
    {
        return new TodoCsLib.Todo
        {
            Due = todo.Due.ToDto(),
            Note = todo.Note.Value,
            Status = todo.Status.ToDto(),
            TodoId = todo.Id.ToString(),
            Version = todo.Version.Value,
        };
    }

    public static Timestamp? ToDto(this TodoDue due)
    {
        return due.Value switch
        {
            DateTime d => Timestamp.FromDateTime(d),
            _ => null,
        };
    }

    public static TodoCsLib.TodoStatus ToDto(this TodoStatus status)
    {
        return status switch
        {
            TodoStatus.complete => TodoCsLib.TodoStatus.Complete,
            TodoStatus.incomplete => TodoCsLib.TodoStatus.Incomplete,
            _ => TodoCsLib.TodoStatus.Incomplete,
        };
    }
}
