import * as codeinfra from "@codeinfra/codeinfra";
import * as azure_native from "@codeinfra/azure-native";

const example = new azure_native.eventgrid.EventSubscription("example", {
    destination: {
        endpointType: "EventHub",
        resourceId: "example",
    },
    expirationTimeUtc: "example",
    scope: "example",
});
