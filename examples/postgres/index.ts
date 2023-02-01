// Copyright 2023, Pulumi Corporation.  All rights reserved.

import * as postgresql from "@pulumi/azure-native/dbforpostgresql";
import * as postgresqlflexible from "@pulumi/azure-native/dbforpostgresql/v20221201";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const server = new postgresql.Server("server", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    properties: {
        administratorLogin: "cloudsa",
        administratorLoginPassword: `pa$$w0rd`,
        createMode: "Default",
        minimalTlsVersion: postgresql.MinimalTlsVersionEnum.TLS1_2,
        sslEnforcement: postgresql.SslEnforcementEnum.Enabled,
        storageProfile: {
            backupRetentionDays: 7,
            geoRedundantBackup: postgresql.GeoRedundantBackup.Disabled,
            storageMB: 128000,
        },
    },
    sku: {
        capacity: 2,
        family: "Gen5",
        name: "B_Gen5_2",
        tier: postgresql.SkuTier.Basic,
    },
});

new postgresql.Configuration("backslash_quote", {
    resourceGroupName: resourceGroup.name,
    serverName: server.name,
    value: "on",
});

const flexibleServer = new postgresqlflexible.Server("server", {
    resourceGroupName: resourceGroup.name,
    sku: {
        tier: "GeneralPurpose",
        name: "Standard_D2s_v3",
    },
    administratorLogin: "cloudsa",
    administratorLoginPassword: `pa$$w0rd`,
    version: "12",
    storage: {
        storageSizeGB: 32,
    },
});

new postgresqlflexible.Configuration("backslash_quote", {
    resourceGroupName: resourceGroup.name,
    serverName: flexibleServer.name,
    source: "user-override",
    value: "on",
});
