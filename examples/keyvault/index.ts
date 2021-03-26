// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as authorization from "@pulumi/azure-native/authorization";
import * as containerinstance from "@pulumi/azure-native/containerinstance";
import * as keyvault from "@pulumi/azure-native/keyvault";
import * as managedidentity from "@pulumi/azure-native/managedidentity";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const userIdentity = new managedidentity.UserAssignedIdentity("uai", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
});

const container = new containerinstance.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
    osType: containerinstance.OperatingSystemTypes.Linux,
    containers: [{
        name: "foo",
        image: "nginx",
        resources: {
            requests: {
                memoryInGB: 1,
                cpu: 1,
            },
        },
    }],
    identity: {
        type: containerinstance.ResourceIdentityType.UserAssigned,
        userAssignedIdentities: userIdentity.id.apply(id => {
            const dict: { [key: string] : object } = {};
            dict[id] = {};
            return dict;
        }),
    },
});

const config = pulumi.output(authorization.getClientConfig());

const vault = new keyvault.Vault("vault", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    properties: {
        accessPolicies: [{
            objectId: config.objectId,
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
            tenantId: config.tenantId,
        }, {
            objectId: container.identity.apply(i => i?.userAssignedIdentities!).apply(uai => Object.values(uai)[0].principalId),
            permissions: {
                secrets: [keyvault.SecretPermissions.Get],
            },
            tenantId: config.tenantId,
        }],
        sku: {
            family: keyvault.SkuFamily.A,
            name: keyvault.SkuName.Standard,
        },
        tenantId: config.tenantId,
    },
});

const secret = new keyvault.Secret("mysecret", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    properties: {
        value: "myvalue",
    },
});

const key = new keyvault.Key("mykey", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    properties: {
        kty: "RSA",
    },
});

const scope = pulumi.interpolate`/subscriptions/${config.subscriptionId}`;
const roleDefinition = new authorization.RoleDefinition("custom-role", {
    roleName: "my-custom-role",
    assignableScopes: [scope],
    permissions: [{
        actions: ["*"],
        notActions: [],
    }],
    scope,
});

export const validatedRoleDefinitionId = pulumi.all([scope, roleDefinition.id]).apply(([scope, id]) => {
    if (id.indexOf(scope) < 0)
        throw new Error(`Expected '${scope}' in ${id} but not found`);
    return id;
});
