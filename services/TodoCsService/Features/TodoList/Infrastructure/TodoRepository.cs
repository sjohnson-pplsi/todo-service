using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Conventions;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;
using TodoCsService.Features.TodoList.Domain.Entities;
using TodoCsService.Features.TodoList.Domain.Repositories;
using TodoCsService.Features.TodoList.Domain.Values;

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

    public async Task<Todo> GetTodo(TodoId id, CancellationToken token = default)
    {
        var filter = Builders<TodoModel>.Filter.Eq("_id", id.Value.ToString());
        var document = (await collection.FindAsync(filter, null, token)).FirstOrDefault(cancellationToken: token) ?? throw new Exception();
        return document.ToEntity();
    }

    public async Task<TodoId> InsertTodo(Todo todo, CancellationToken token = default)
    {
        await collection.InsertOneAsync(todo.ToModel(), new InsertOneOptions(), token);
        return todo.Id;
    }

    public async Task<ICollection<Todo>> ListTodos(int limit, int offset, CancellationToken token = default)
    {
        var filter = Builders<TodoModel>.Filter.Empty;
        var result = await collection.FindAsync(filter, null, token);
        return result.ToList(token).Select(t => t.ToEntity()).ToList();
    }

    public async Task ReplaceTodo(Todo todo, CancellationToken token = default)
    {
        var builder = Builders<TodoModel>.Filter;
        var filter = builder.Eq("_id", todo.Id.Value.ToString()) & builder.Eq("version", todo.Version.Value);
        var m = todo.ToModel().Increment();
        await collection.ReplaceOneAsync(filter, m, new ReplaceOptions(), token);
    }
}
