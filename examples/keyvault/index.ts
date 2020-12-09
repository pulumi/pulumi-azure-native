import * as random from "@pulumi/random";
import * as keyvault from "@pulumi/azure-nextgen/keyvault/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
    number: false,
}).result;

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString,
    location: "westus2",
});

// TODO: read IDs with an invoke when it's available
// https://github.com/pulumi/pulumi-azure-nextgen/issues/107
const tenantId = "706143bc-e1d4-4593-aee2-c9dc60ab9be7";
const currentPrincipalId = "e7d1b8ea-596b-44be-890f-8355a351244c";

const vault = new keyvault.Vault("vault", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    vaultName: randomString,
    properties: {
        accessPolicies: [{
            objectId: currentPrincipalId,
            permissions: {
                keys: [
                    keyvault.KeyPermissions.Get,
                    keyvault.KeyPermissions.Create,
                    keyvault.KeyPermissions.Delete,
                    keyvault.KeyPermissions.List,
                    keyvault.KeyPermissions.Recover,
                    keyvault.KeyPermissions.Purge,
                ],
                secrets: [
                    keyvault.SecretPermissions.Get,
                    keyvault.SecretPermissions.List,
                    keyvault.SecretPermissions.Set,
                    keyvault.SecretPermissions.Delete,
                    keyvault.SecretPermissions.Recover,
                    keyvault.SecretPermissions.Purge,
                ],
            },
            tenantId,
        }],
        sku: {
            family: keyvault.SkuFamily.A,
            name: keyvault.SkuName.Standard,
        },
        tenantId,
    },
});

const secret = new keyvault.Secret("mysecret", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    secretName: "mysecret",
    properties: {
        value: "myvalue",
    },
});

const key = new keyvault.Key("mykey", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    keyName: "mykey",
    properties: {
        kty: "RSA",
        keySize: 2048,
        keyOps: [
            "decrypt",
            "encrypt",
            "sign",
            "unwrapKey",
            "verify",
            "wrapKey",
        ],
    },
});
