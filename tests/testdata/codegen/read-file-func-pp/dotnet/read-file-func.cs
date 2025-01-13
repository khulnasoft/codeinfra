using System.Collections.Generic;
using System.IO;
using System.Linq;
using Codeinfra;

return await Deployment.RunAsync(() => 
{
    var key = File.ReadAllText("key.pub");

    return new Dictionary<string, object?>
    {
        ["result"] = key,
    };
});

