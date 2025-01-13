using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using AzureNative = Codeinfra.AzureNative;

return await Deployment.RunAsync(() => 
{
    var example = new AzureNative.EventGrid.EventSubscription("example", new()
    {
        Destination = new AzureNative.EventGrid.Inputs.EventHubEventSubscriptionDestinationArgs
        {
            EndpointType = "EventHub",
            ResourceId = "example",
        },
        ExpirationTimeUtc = "example",
        Scope = "example",
    });

});

