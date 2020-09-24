import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const acrAdminUserEnabledParam = config.getBoolean("acrAdminUserEnabledParam") || false;
const acrNameParam = config.require("acrNameParam");
const acrSkuParam = config.get("acrSkuParam") || "Basic";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_nextgen.resources.latest.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const registryResource = new azure_nextgen.containerregistry.v20191201preview.Registry("registryResource", {
    adminUserEnabled: acrAdminUserEnabledParam,
    location: locationParam,
    registryName: acrNameParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: acrSkuParam,
    },
});
export const acrLoginServerOut = registryResource.loginServer;
