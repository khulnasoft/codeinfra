using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Random = Codeinfra.Random;

return await Deployment.RunAsync(() => 
{
    var random_pet = new Random.RandomPet("random-pet", new()
    {
        Prefix = "doggo",
    });

});

