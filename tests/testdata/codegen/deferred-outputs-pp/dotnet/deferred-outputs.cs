using System.Collections.Generic;
using System.Linq;
using Codeinfra;

return await Deployment.RunAsync(() => 
{
    var secondPasswordLength = new Codeinfra.DeferredOutput<int>();
    var first = new Components.First("first", new()
    {
        PasswordLength = secondPasswordLength.Output,
    });

    var second = new Components.Second("second", new()
    {
        PetName = first.PetName,
    });

    secondPasswordLength.Resolve(second.PasswordLength);
    var loopingOverMany = new Codeinfra.DeferredOutput<List<int>>();
    var another = new Components.First("another", new()
    {
        PasswordLength = loopingOverMany.Output.Apply(loopingOverMany => loopingOverMany.Length),
    });

    var many = new List<Components.Second>();
    for (var rangeIndex = 0; rangeIndex < 10; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        many.Add(new Components.Second($"many-{range.Value}", new()
        {
            PetName = another.PetName,
        }));
    }
    loopingOverMany.Resolve(Output.Create(many.Select((value, i) => new { Key = i.ToString(), Value = pair.Value }).Select(v => 
    {
        return v.PasswordLength;
    }).ToList()));
});

