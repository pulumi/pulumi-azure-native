import * as resources from "@pulumi/azure-native/resources";
import * as storage_v20240101 from "@pulumi/azure-native_storage_v20240101";

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("resourceGroup");

// Create an Azure resource (Storage Account)
const storageAccount = new storage_v20240101.storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage_v20240101.storage.SkuName.Standard_LRS,
    },
    kind: storage_v20240101.storage.Kind.StorageV2,
});

export const storageAccountName = storageAccount.name;
