namespace TodoCsService.Features.Base;

public class BaseGuid
{
    public static Guid? Parse(string value) => Guid.TryParse(value, out Guid id) ? id : null;
}
