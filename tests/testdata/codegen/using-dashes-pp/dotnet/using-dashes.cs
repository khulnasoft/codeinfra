using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using UsingDashes = Codeinfra.UsingDashes;

return await Deployment.RunAsync(() => 
{
    var main = new UsingDashes.Dash("main", new()
    {
        Stack = "dev",
    });

});

