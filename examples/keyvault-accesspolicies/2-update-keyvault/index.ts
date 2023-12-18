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

const ap1 = new keyvault.AccessPolicy("ap1", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    accessPolicy: {
        objectId: config.objectId,
        permissions: {
            keys: [
                keyvault.KeyPermissions.Get,
                keyvault.KeyPermissions.Create,
                keyvault.KeyPermissions.Delete,
                keyvault.KeyPermissions.List,
            ],
            secrets: [
                keyvault.SecretPermissions.Get,
                keyvault.SecretPermissions.List,
                keyvault.SecretPermissions.Set,
                keyvault.SecretPermissions.Delete,
            ],
        },
        tenantId: config.tenantId,
    }
});

const ap2 = new keyvault.AccessPolicy("ap2", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    accessPolicy: {
        objectId: "00000000-0000-0000-0000-000000000000",
        permissions: {
            keys: [
                keyvault.KeyPermissions.Get,
            ]
        },
        tenantId: config.tenantId,
    }
});

export const rgName = resourceGroup.name;
export const kvName = vault.name;