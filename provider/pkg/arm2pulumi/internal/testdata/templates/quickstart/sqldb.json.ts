import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const administratorLoginParam = config.require("administratorLoginParam");
const administratorLoginPasswordParam = config.require("administratorLoginPasswordParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_native.resources.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const serverNameParam = config.get("serverNameParam") || "[uniqueString('sql', resourceGroup().id)]";
const sqlDBNameParam = config.get("sqlDBNameParam") || "SampleDB";
const serverResource = new azure_native.sql.v20190601preview.Server("serverResource", {
    administratorLogin: administratorLoginParam,
    administratorLoginPassword: administratorLoginPasswordParam,
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    serverName: serverNameParam,
});
const databaseResource = new azure_native.sql.v20190601preview.Database("databaseResource", {
    databaseName: `${serverNameParam}/${sqlDBNameParam}`,
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    serverName: serverResource.name,
    sku: {
        name: "Standard",
        tier: "Standard",
    },
});
