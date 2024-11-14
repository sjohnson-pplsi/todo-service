using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Conventions;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;
using TodoCsService.Features.TodoList.Domain.Entity;
using TodoCsService.Features.TodoList.Domain.Repositories;
using TodoCsService.Features.TodoList.Domain.Value;

namespace TodoCsService.Features.TodoList.Infrastructure;

public class TodoRepository(IMongoDatabase database) : ITodoRepository
{
    readonly IMongoCollection<TodoModel> collection =
        database.GetCollection<TodoModel>("todos");

    public static void Register()
    {
        BsonSerializer.RegisterSerializer(new GuidSerializer(GuidRepresentation.Standard));
        var conventionPack = new ConventionPack { new CamelCaseElementNameConvention() };
        ConventionRegistry.Register("camelCase", conventionPack, t => true);
        TodoModel.Register();
    }

    public async Task<Todo> GetTodo(TodoId id)
    {
        var filter = Builders<TodoModel>.Filter.Eq("_id", id.Value.ToString());
        var document = (await collection.FindAsync(filter)).FirstOrDefault() ?? throw new Exception();
        return document.ToEntity();
    }

    public async Task<TodoId> InsertTodo(Todo todo)
    {
        await collection.InsertOneAsync(todo.ToModel());
        return todo.Id;
    }

    public async Task<ICollection<Todo>> ListTodos(int limit, int offset)
    {
        var filter = Builders<TodoModel>.Filter.Empty;
        var result = await collection.FindAsync(filter);
        return result.ToList().Select(t => t.ToEntity()).ToList();
    }

    public async Task ReplaceTodo(Todo todo)
    {
        var builder = Builders<TodoModel>.Filter;
        var filter = builder.Eq("_id", todo.Id.Value.ToString()) & builder.Eq("version", todo.Version.Value);
        var m = todo.ToModel().Increment();
        await collection.ReplaceOneAsync(filter, m);
    }
}
