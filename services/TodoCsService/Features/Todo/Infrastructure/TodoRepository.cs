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
    readonly IMongoCollection<Domain.Entity.Todo> collection =
        database.GetCollection<Domain.Entity.Todo>("todos");

    public static void Register()
    {
        BsonSerializer.RegisterSerializer(new GuidSerializer(GuidRepresentation.Standard));
        BsonClassMap.RegisterClassMap<Entity<TodoId>>(cm =>
        {
            cm.AutoMap();
            cm.MapIdMember(c => c.Id);
        });
    }

    public async Task<Domain.Entity.Todo> GetTodo(TodoId id)
    {
        var filter = Builders<Domain.Entity.Todo>.Filter.Eq("_id", id);
        var document = (await collection.FindAsync(filter)).First() ?? throw new Exception();
        return document;
    }

    public async Task<TodoId> InsertTodo(Domain.Entity.Todo todo)
    {
        await collection.InsertOneAsync(todo);
        return todo.Id;
    }

    public async Task<ICollection<Domain.Entity.Todo>> ListTodos(int limit, int offset)
    {
        var filter = Builders<Domain.Entity.Todo>.Filter.Empty;
        var result = await collection.FindAsync(filter);
        return result.ToList();
    }

    public async Task ReplaceTodo(Domain.Entity.Todo todo)
    {
        var filter = Builders<Domain.Entity.Todo>.Filter.Eq("_id", todo.Id);
        await collection.ReplaceOneAsync(filter, todo);
    }
}