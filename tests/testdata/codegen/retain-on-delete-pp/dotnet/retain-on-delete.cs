using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Random = Codeinfra.Random;

return await Deployment.RunAsync(() => 
{
    var foo = new Random.RandomPet("foo", new()
    {
    }, new CustomResourceOptions
    {
        RetainOnDelete = true,
    });

});

