using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Unknown = Codeinfra.Unknown;

return await Deployment.RunAsync(() => 
{
    var data = Unknown.Index.GetData.Invoke(new()
    {
        Input = "hello",
    });

    var values = Unknown.Eks.ModuleValues.Invoke();

    return new Dictionary<string, object?>
    {
        ["content"] = data.Content,
    };
});

