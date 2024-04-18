// Copyright 2023, Pulumi Corporation.  All rights reserved.

import * as postgresql from "@pulumi/azure-native/dbforpostgresql";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const flexibleServer = new postgresql.Server("server", {
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

new postgresql.Configuration("backslash_quote", {
    resourceGroupName: resourceGroup.name,
    serverName: flexibleServer.name,
    source: "user-override",
    value: "on",
});
