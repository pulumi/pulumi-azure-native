// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as authorization from "@pulumi/azure-native/authorization";
import * as keyvault from "@pulumi/azure-native/keyvault";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const config = pulumi.output(authorization.getClientConfig());

const vault = new keyvault.Vault("vault", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    properties: {
        sku: {
            family: keyvault.SkuFamily.A,
            name: keyvault.SkuName.Standard,
        },
        tenantId: config.tenantId,
        enabledForDeployment: true,
    },
});

// ap1 has different permissions in step 3
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

// ap2 is deleted in step 3

export const rgName = resourceGroup.name;
export const kvName = vault.name;