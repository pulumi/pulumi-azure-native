// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as mysqlpreview from "@pulumi/azure-native/dbformysql/v20210501";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const flexibleServer = new mysqlpreview.Server("server", {
    resourceGroupName: resourceGroup.name,
    sku: {
        tier: "Burstable",
        name: "Standard_B1s",
    },
    administratorLogin: "cloudsa",
    administratorLoginPassword: `pa$$w0rd`,
    version: "5.7",
});

new mysqlpreview.Configuration("innodb_strict_mode", {
    resourceGroupName: resourceGroup.name,
    serverName: flexibleServer.name,
    source: "user-override",
    value: "off",
});
