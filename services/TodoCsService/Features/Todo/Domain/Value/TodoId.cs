namespace TodoCsService.Features.Todo.Domain.Value;

public record TodoId(Guid Value)
{
    public override string ToString() => Value.ToString();

    public static TodoId Parse(string value) => new TodoId(Guid.Parse(value));
}
