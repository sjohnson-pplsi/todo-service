namespace TodoCsService.Features.Todo.Domain.Repositories;

public interface ITodoRepository
{
    Task<Entity.Todo> GetTodo(Value.TodoId id);
    Task<Value.TodoId> InsertTodo(Entity.Todo todo);
    Task<ICollection<Entity.Todo>> ListTodos(int limit, int offset);
    Task ReplaceTodo(Entity.Todo todo);
}
