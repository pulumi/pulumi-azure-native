// Copyright 2023, Pulumi Corporation.  All rights reserved.

import * as postgresql from "@pulumi/azure-native/dbforpostgresql";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg", {
    // // westus2 was not available for Postgres Flexible Server at the time of writing, see
    // // https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/overview#azure-regions
    // location: "uksouth",
});

const flexibleServer = new postgresql.Server("server", {
    location: resourceGroup.location,
    resourceGroupName: resourceGroup.name,
    sku: {
        tier: "GeneralPurpose",
        name: "Standard_D2s_v3",
    },
    highAvailability: {
        mode: postgresql.HighAvailabilityMode.SameZone,
    },
    administratorLogin: "cloudsa",
    administratorLoginPassword: `pa$$w0rd`,
    version: "14",
    storage: {
        storageSizeGB: 64,
    },
});

new postgresql.Configuration("backslash_quote", {
    resourceGroupName: resourceGroup.name,
    serverName: flexibleServer.name,
    source: "user-override",
    value: "on",
});
