using TodoCsService.Features.TodoList.Domain.Entities;

namespace TodoCsService.Features.TodoList.Domain.Repositories;

public interface ITodoRepository
{
    Task<Todo> GetTodo(Values.TodoId id, CancellationToken token = default);
    Task<Values.TodoId> InsertTodo(Todo todo, CancellationToken token = default);
    Task<ICollection<Todo>> ListTodos(int limit, int offset, CancellationToken token = default);
    Task ReplaceTodo(Todo todo, CancellationToken token = default);
}
