using MongoDB.Driver;
using TodoCsService.Features.Base;
using TodoCsService.Features.TodoList.Controllers;
using TodoCsService.Features.TodoList.Domain.Repositories;
using TodoCsService.Features.TodoList.Domain.Services;
using TodoCsService.Features.TodoList.Infrastructure;

var connectionString = Environment.GetEnvironmentVariable("MONGODB_URI") ?? "mongodb://localhost:27017/?maxPoolSize=20&w=majority";
var client = new MongoClient(connectionString);
var database = client.GetDatabase("todo-service");

BaseRepository.Register();
TodoRepository.Register();

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddSingleton(database);
builder.Services.AddScoped<TodoService>();
builder.Services.AddScoped<ITodoRepository, TodoRepository>();

// Add services to the container.
builder.Services.AddGrpc();
builder.Services.AddGrpcReflection();

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<TodoController>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

IWebHostEnvironment env = app.Environment;

if (env.IsDevelopment())
{
    app.MapGrpcReflectionService();
}

app.Run();
