import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";

const resourceGroup = new resources.ResourceGroup("rg");

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
    networkRuleSet: {
        defaultAction: "Deny",
        ipRules: [
            {
                action: "Allow",
                iPAddressOrRange: "78.23.36.13"
            }
        ]
    }
});

export const rg = resourceGroup.name;
export const sa = storageAccount.name;
