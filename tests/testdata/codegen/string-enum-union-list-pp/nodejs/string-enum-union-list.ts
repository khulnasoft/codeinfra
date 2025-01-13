import * as codeinfra from "@codeinfra/codeinfra";
import * as azure_native from "@codeinfra/azure-native";

const namespaceAuthorizationRule = new azure_native.servicebus.NamespaceAuthorizationRule("namespaceAuthorizationRule", {
    authorizationRuleName: "sdk-AuthRules-1788",
    namespaceName: "sdk-Namespace-6914",
    resourceGroupName: "ArunMonocle",
    rights: [
        azure_native.servicebus.AccessRights.Listen,
        azure_native.servicebus.AccessRights.Send,
    ],
});
