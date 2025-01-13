import * as codeinfra from "@codeinfra/codeinfra";
import * as azure from "@codeinfra/azure";

const config = new codeinfra.Config();
// The name of the storage account
const storageAccountNameParam = config.require("storageAccountNameParam");
// The name of the resource group
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure.core.getResourceGroup({
    name: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});
export const storageAccountNameOut = storageAccountResource.name;
