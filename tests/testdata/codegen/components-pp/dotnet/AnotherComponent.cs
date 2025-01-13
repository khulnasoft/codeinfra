using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Random = Codeinfra.Random;

namespace Components
{
    public class AnotherComponent : global::Codeinfra.ComponentResource
    {
        public AnotherComponent(string name, ComponentResourceOptions? opts = null)
            : base("components:index:AnotherComponent", name, ResourceArgs.Empty, opts)
        {
            var firstPassword = new Random.RandomPassword($"{name}-firstPassword", new()
            {
                Length = 16,
                Special = true,
            }, new CustomResourceOptions
            {
                Parent = this,
            });

            this.RegisterOutputs();
        }
    }
}
