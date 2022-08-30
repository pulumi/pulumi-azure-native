// All of these examples work under released version of Pulumi, but
// the introduction of lazy loading breaks test1 and test2.

import * as pulumi from "@pulumi/pulumi";

// Import pattern 1: we currently do not recommend this pattern
// but we do support it.
import * as azure from "@pulumi/azure-native";

export function test1(resourceGroupName: string): azure.storage.StorageAccount {
    const args: azure.storage.StorageAccountArgs = {
        resourceGroupName: resourceGroupName,
        sku: {
            name:azure.storage.SkuName.Standard_LRS,
        },
        kind: azure.storage.Kind.StorageV2,
    };

    return new azure.storage.StorageAccount("sa", args);
}


// Import pattern 2: not recommended but possible.
import { storage as storage2 } from "@pulumi/azure-native";

export function test2(resourceGroupName: string): storage2.StorageAccount {
    const args: storage2.StorageAccountArgs = {
        resourceGroupName: resourceGroupName,
        sku: {
            name:storage2.SkuName.Standard_LRS,
        },
        kind: storage2.Kind.StorageV2,
    };

    return new storage2.StorageAccount("sa", args);
}


// Import pattern 3: recommended.
import * as storage from "@pulumi/azure-native/storage";

export function test3(resourceGroupName: string): storage.StorageAccount {
    const args: storage.StorageAccountArgs = {
        resourceGroupName: resourceGroupName,
        sku: {
            name:storage.SkuName.Standard_LRS,
        },
        kind: storage.Kind.StorageV2,
    };

    return new storage.StorageAccount("sa", args);
}
