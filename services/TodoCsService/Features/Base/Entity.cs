namespace TodoCsService.Features.Base;

public class Entity<T>(T id, AggregateVersion version)
{
    public readonly T Id = id;
    public readonly AggregateVersion Version = version;

    public override bool Equals(object? obj)
    {
        return obj switch
        {
            null => false,
            Entity<T> e => Id?.Equals(e.Id) ?? false,
            _ => false,
        };
    }

    public override int GetHashCode()
    {
        return Id?.GetHashCode() ?? string.Empty.GetHashCode();
    }
}
