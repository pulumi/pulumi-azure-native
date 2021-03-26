// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const acrAdminUserEnabledParam = config.getBoolean("acrAdminUserEnabledParam") || false;
const acrNameParam = config.require("acrNameParam");
const acrSkuParam = config.get("acrSkuParam") || "Basic";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_native.resources.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const registryResource = new azure_native.containerregistry.v20191201preview.Registry("registryResource", {
    adminUserEnabled: acrAdminUserEnabledParam,
    location: locationParam,
    registryName: acrNameParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: acrSkuParam,
    },
    tags: {
        "container.registry": acrNameParam,
        displayName: "Container Registry",
    },
});
export const acrLoginServerOut = registryResource.loginServer;
