using TodoCsService.Features.Base;
using TodoCsService.Features.Todo.Domain.Value;

namespace TodoCsService.Features.Todo.Domain.Entity;

public record Todo(
    TodoId Id,
    AggregateVersion Version,
    TodoNote Note,
    TodoDue Due,
    TodoStatus Status
) : Entity<TodoId>(Id, Version)
{
    public Todo Complete() => this with { Status = TodoStatus.complete };

    public Todo Reset() => this with { Status = TodoStatus.incomplete };

    public Todo ChangeNote(TodoNote note) => this with { Note = note };
}
