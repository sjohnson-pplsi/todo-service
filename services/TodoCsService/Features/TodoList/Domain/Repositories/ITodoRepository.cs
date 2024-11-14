using TodoCsService.Features.TodoList.Domain.Entity;

namespace TodoCsService.Features.TodoList.Domain.Repositories;

public interface ITodoRepository
{
    Task<Todo> GetTodo(Value.TodoId id);
    Task<Value.TodoId> InsertTodo(Todo todo);
    Task<ICollection<Todo>> ListTodos(int limit, int offset);
    Task ReplaceTodo(Todo todo);
}
