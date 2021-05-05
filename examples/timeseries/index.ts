// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as resources from "@pulumi/azure-native/resources";
import * as timeseries from "@pulumi/azure-native/timeseriesinsights";

const resourceGroup = new resources.ResourceGroup("rg");

new timeseries.Gen1Environment("environment", {
    resourceGroupName: resourceGroup.name,
    kind: "Gen1",
    sku: {
        name: "S1",
        capacity: 1
    },
    dataRetentionTime: "P31D",
    partitionKeyProperties: [{
        name: "DeviceId1",
        type: "String"
    }],
});
