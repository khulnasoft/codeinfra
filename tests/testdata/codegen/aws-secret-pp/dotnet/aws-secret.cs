using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Aws = Codeinfra.Aws;

return await Deployment.RunAsync(() => 
{
    var dbCluster = new Aws.Rds.Cluster("dbCluster", new()
    {
        MasterPassword = Output.CreateSecret("foobar"),
    });

});

