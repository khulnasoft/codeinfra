using System.Collections.Generic;
using System.Linq;
using Codeinfra;

return await Deployment.RunAsync(() => 
{
    var stackRef = new Codeinfra.StackReference("stackRef", new()
    {
        Name = "foo/bar/dev",
    });

});

