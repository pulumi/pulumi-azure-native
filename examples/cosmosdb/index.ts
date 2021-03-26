// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as resources from "@pulumi/azure-native/resources";
import * as cosmosdb from "./component";

const resourceGroup = new resources.ResourceGroup("rg", {
    location: "westeurope",
});

export const cosmosdbAccount = new cosmosdb.DatabaseAccount("da", {
    resourceGroup,
    api: "Sql",
    consisencyPolicy: { defaultConsistencyLevel: "Session" },
    location: { 
        type: "single",
        region: "westeurope",
    },    
});
