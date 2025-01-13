using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Splat = Codeinfra.Splat;

return await Deployment.RunAsync(() => 
{
    var allKeys = Splat.GetSshKeys.Invoke();

    var main = new Splat.Server("main", new()
    {
        SshKeys = allKeys.Apply(getSshKeysResult => getSshKeysResult.SshKeys).Select(__item => __item.Name).ToList(),
    });

});

