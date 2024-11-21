using TodoCsService.Features.TodoList.Domain.Entities;

namespace TodoCsService.Features.TodoList.Domain.Repositories;

public interface ITodoRepository
{
    Task<Todo> GetTodo(Values.TodoId id);
    Task<Values.TodoId> InsertTodo(Todo todo);
    Task<ICollection<Todo>> ListTodos(int limit, int offset);
    Task ReplaceTodo(Todo todo);
}
