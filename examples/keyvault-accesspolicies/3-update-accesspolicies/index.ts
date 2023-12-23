// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as authorization from "@pulumi/azure-native/authorization";
import * as keyvault from "@pulumi/azure-native/keyvault";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const config = pulumi.output(authorization.getClientConfig());

// enabledForDeployment changed back to false
const vault = new keyvault.Vault("vault", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    properties: {
        sku: {
            family: keyvault.SkuFamily.A,
            name: keyvault.SkuName.Standard,
        },
        tenantId: config.tenantId,
        enabledForDeployment: false,
    },
});

// ap1 has different permissions
const ap1 = new keyvault.AccessPolicy("ap1", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    policy: {
        objectId: config.objectId,
        permissions: {
            keys: [
                keyvault.KeyPermissions.Get,
            ],
            secrets: [
                keyvault.SecretPermissions.Get,
            ]
        },
        tenantId: config.tenantId,
    }
});

// ap2 is deleted

export const rgName = resourceGroup.name;
export const kvName = vault.name;

// Read the Vault's state directly from Azure. We export it so the test can use it to make
// additional assertions that Pulumi's view of the world matches Azure's.
// We use `apply` here as a barrier, ensuring the invoke runs after the vault is updated.
const newVaultState = vault.properties?.apply(_ => {
    return keyvault.getVaultOutput({
        resourceGroupName: resourceGroup.name,
        vaultName: vault.name,
    });
})

export const numberOfAPs = newVaultState.properties.accessPolicies?.apply(aps => aps?.length || 0 );
export const aps = newVaultState.properties.accessPolicies;