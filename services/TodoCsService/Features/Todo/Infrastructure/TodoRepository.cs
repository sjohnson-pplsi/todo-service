using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;
using TodoCsService.Features.Base;
using TodoCsService.Features.Todo.Domain.Repositories;
using TodoCsService.Features.Todo.Domain.Value;

namespace TodoCsService.Features.Todo.Infrastructure;

public class TodoRepository(IMongoDatabase database) : ITodoRepository
{
    readonly IMongoCollection<TodoModel> collection =
        database.GetCollection<TodoModel>("todos");

    public static void Register()
    {
        BsonSerializer.RegisterSerializer(new GuidSerializer(GuidRepresentation.Standard));
        BsonClassMap.RegisterClassMap<TodoModel>(cm =>
        {
            cm.AutoMap();
            cm.MapIdMember(c => c.Id);
        });
    }

    public async Task<Domain.Entity.Todo> GetTodo(TodoId id)
    {
        var filter = Builders<TodoModel>.Filter.Eq("_id", id);
        var document = (await collection.FindAsync(filter)).First() ?? throw new Exception();
        return document.ToEntity();
    }

    public async Task<TodoId> InsertTodo(Domain.Entity.Todo todo)
    {
        await collection.InsertOneAsync(todo.ToModel());
        return todo.Id;
    }

    public async Task<ICollection<Domain.Entity.Todo>> ListTodos(int limit, int offset)
    {
        var filter = Builders<TodoModel>.Filter.Empty;
        var result = await collection.FindAsync(filter);
        return result.ToList().Select(t => t.ToEntity()).ToList();
    }

    public async Task ReplaceTodo(Domain.Entity.Todo todo)
    {
        var filter = Builders<TodoModel>.Filter.Eq("_id", todo.Id);
        await collection.ReplaceOneAsync(filter, todo.ToModel());
    }
}