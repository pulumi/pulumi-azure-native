import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("resourceGroup");

// Create an Azure resource (Storage Account)
const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
});

const container = new storage.BlobContainer("container", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: "lhcontainer",
});

const legalHold = new storage.BlobContainerLegalHold("legalHold", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: container.name,
    tags: ["tag1", "tag2"],
    allowProtectedAppendWritesAll: false,
});

export const containerName = container.name;
export const accountName = storageAccount.name;
export const resourceGroupName = resourceGroup.name;
