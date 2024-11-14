using TodoCsService.Features.Base;
using TodoCsService.Features.TodoList.Domain.Value;

namespace TodoCsService.Features.TodoList.Domain.Entity;

public class Todo(
    TodoId id,
    AggregateVersion version,
    TodoNote note,
    TodoDue due,
    TodoStatus status
) : Entity<TodoId>(id, version)
{
    public TodoNote Note { get; private set; } = note;
    public TodoDue Due { get; private set; } = due;
    public TodoStatus Status { get; private set; } = status;

    public void Complete()
    {
        Status = TodoStatus.complete;
    }

    public void Reset()
    {
        Status = TodoStatus.incomplete;
    }

    public void ChangeNote(TodoNote note)
    {
        Note = note;
    }
}
