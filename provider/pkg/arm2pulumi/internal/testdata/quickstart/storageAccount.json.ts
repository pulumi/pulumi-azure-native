import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_nextgen.resources.latest.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountNameParam = config.require("storageAccountNameParam");
const storageAccountTypeParam = config.get("storageAccountTypeParam") || "Standard_LRS";
const storageAccountResource = new azure_nextgen.storage.v20190401.StorageAccount("storageAccountResource", {
    accountName: storageAccountNameParam,
    kind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: storageAccountTypeParam,
    },
});
export const storageAccountNameOut = storageAccountResource.name;
