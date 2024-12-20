// Copyright 2024, Pulumi Corporation.  All rights reserved.

import * as recoveryservices from '@pulumi/azure-native/recoveryservices';
import * as recoveryservicesTypes from '@pulumi/azure-native/types';
import * as resources from '@pulumi/azure-native/resources';
import * as storage from '@pulumi/azure-native/storage';
import * as pulumi from '@pulumi/pulumi';

const resourceGroup = new resources.ResourceGroup('resourceGroup', {
    location: 'westeurope',
});

const vault = new recoveryservices.Vault('vault', {
    identity: {
        type: 'SystemAssigned',
    },
    location: resourceGroup.location,
    properties: {
        publicNetworkAccess: 'Enabled',
        securitySettings: {
            softDeleteSettings: {
                // Deleting the vault is not allowed with soft delete enabled.
                softDeleteState: recoveryservices.SoftDeleteState.Disabled,
            },
        },
    },
    resourceGroupName: resourceGroup.name,
    sku: {
        name: 'Standard',
    },
});

const storageAccount = new storage.StorageAccount('sa', {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    kind: 'StorageV2',
    sku: {
        name: 'Standard_LRS',
    },
    publicNetworkAccess: 'Enabled',
}, {
    dependsOn: [vault],
});

const protectionPolicy = new recoveryservices.ProtectionPolicy('protectionPolicy', {
    location: resourceGroup.location,
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    properties: {
        backupManagementType: recoveryservices.BackupManagementType.AzureStorage,
        schedulePolicy: {
            schedulePolicyType: 'SimpleSchedulePolicy',
            scheduleRunFrequency: 'Daily',
            scheduleRunTimes: ['2024-10-08T19:30:00Z'],
        },
    retentionPolicy: {
        retentionPolicyType: 'LongTermRetentionPolicy',
        dailySchedule: {
                retentionDuration: {
                    count: 30,
                    durationType: 'Days',
                },
                retentionTimes: ['2024-10-08T19:30:00Z'],
            },
        },
        timeZone: 'GTB Standard Time',
        workLoadType: recoveryservices.WorkloadType.AzureFileShare,
    } // as recoveryservicesTypes.input.recoveryservices.AzureFileShareProtectionPolicyArgs,
}, {
    dependsOn: [storageAccount],
});

const protectionContainer = new recoveryservices.ProtectionContainer("exampleProtectionContainer", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    containerName: pulumi.interpolate`storagecontainer;storage;${resourceGroup.name};${storageAccount.name}`,
    fabricName: "Azure",
    properties: {
        friendlyName: "exampleContainer",
        containerType: "StorageContainer",
        backupManagementType: recoveryservices.BackupManagementType.AzureStorage,
        sourceResourceId: storageAccount.id,
    },
});

const fileShare = new storage.FileShare('fileshare', {
    accountName: storageAccount.name,
    accessTier: 'TransactionOptimized',
    resourceGroupName: resourceGroup.name,
}, {
    deletedWith: storageAccount,
});

// Enable backup for Azure File Share (Protected Item)
const protectedItem = new recoveryservices.ProtectedItem('protectedItem', {
    containerName: protectionContainer.name,
    fabricName: 'Azure',
    properties: {
        policyId: protectionPolicy.id,
        protectedItemType: 'AzureFileShareProtectedItem',
        sourceResourceId: storageAccount.id,
    }, // as recoveryservicesTypes.input.recoveryservices.AzureFileshareProtectedItemArgs,
    protectedItemName: fileShare.name,
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
});