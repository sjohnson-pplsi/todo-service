namespace TodoCsService.Features.Base;

public class BaseEnum
{
    public static TEnum Parse<TEnum>(string value, TEnum defaultValue) where TEnum : struct
    {
        return Enum.TryParse(value, out TEnum status)
                ? status
                : defaultValue;
    }
}
