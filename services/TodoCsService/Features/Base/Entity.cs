namespace TodoCsService.Features.Base;

public record Entity<T>(
    T Id,
    AggregateVersion Version
);
