// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_native.resources.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const storageAccountNameParam = config.require("storageAccountNameParam");
const storageAccountTypeParam = config.get("storageAccountTypeParam") || "Standard_LRS";
const storageAccountResource = new azure_native.storage.v20190401.StorageAccount("storageAccountResource", {
    accountName: storageAccountNameParam,
    kind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: storageAccountTypeParam,
    },
});
export const storageAccountNameOut = storageAccountResource.name;
