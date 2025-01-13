using System.Collections.Generic;
using System.Linq;
using Codeinfra;

return await Deployment.RunAsync(() => 
{
    return new Dictionary<string, object?>
    {
        ["output_true"] = true,
        ["output_false"] = false,
        ["output_number"] = 4,
        ["output_string"] = "hello",
    };
});

