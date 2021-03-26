// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const farmerpulumi2 = new azure_native.storage.v20190401.StorageAccount("farmerpulumi2", {
    accountName: "farmerpulumi2",
    kind: "StorageV2",
    location: "northeurope",
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: "Standard_LRS",
    },
    tags: {},
});
const farmerpulumi3Ai = new azure_native.insights.v20150501.Component("farmerpulumi3Ai", {
    applicationType: "web",
    disableIpMasking: false,
    kind: "web",
    location: "northeurope",
    resourceGroupName: resourceGroupNameParam,
    resourceName: "farmerpulumi3-ai",
    samplingPercentage: 100,
    tags: {
        "[concat('hidden-link:', resourceGroup().id, '/providers/Microsoft.Web/sites/', 'farmerpulumi3')]": "Resource",
    },
});
const farmerpulumi3Farm = new azure_native.web.v20180201.AppServicePlan("farmerpulumi3Farm", {
    location: "northeurope",
    name: "farmerpulumi3-farm",
    perSiteScaling: false,
    reserved: false,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        capacity: 1,
        name: "F1",
        size: "0",
        tier: "Free",
    },
    tags: {},
});
const farmerpulumi3 = new azure_native.web.v20160801.WebApp("farmerpulumi3", {
    httpsOnly: false,
    kind: "app",
    location: "northeurope",
    name: "farmerpulumi3",
    resourceGroupName: resourceGroupNameParam,
    serverFarmId: "farmerpulumi3-farm",
    siteConfig: {
        alwaysOn: false,
        appSettings: [
            {
                name: "APPINSIGHTS_INSTRUMENTATIONKEY",
                value: "[reference(resourceId('Microsoft.Insights/components', 'farmerpulumi3-ai'), '2014-04-01').InstrumentationKey]",
            },
            {
                name: "APPINSIGHTS_PROFILERFEATURE_VERSION",
                value: "1.0.0",
            },
            {
                name: "APPINSIGHTS_SNAPSHOTFEATURE_VERSION",
                value: "1.0.0",
            },
            {
                name: "ApplicationInsightsAgent_EXTENSION_VERSION",
                value: "~2",
            },
            {
                name: "DiagnosticServices_EXTENSION_VERSION",
                value: "~3",
            },
            {
                name: "InstrumentationEngine_EXTENSION_VERSION",
                value: "~1",
            },
            {
                name: "SnapshotDebugger_EXTENSION_VERSION",
                value: "~1",
            },
            {
                name: "XDT_MicrosoftApplicationInsights_BaseExtensions",
                value: "~1",
            },
            {
                name: "XDT_MicrosoftApplicationInsights_Mode",
                value: "recommended",
            },
            {
                name: "storageKey",
                value: "[concat('DefaultEndpointsProtocol=https;AccountName=farmerpulumi2;AccountKey=', listKeys(resourceId('Microsoft.Storage/storageAccounts', 'farmerpulumi2'), '2017-10-01').keys[0].value)]",
            },
        ],
        connectionStrings: [],
    },
    tags: {},
}, {
    dependsOn: [
        farmerpulumi3Farm,
        farmerpulumi2,
        farmerpulumi3Ai,
    ],
});
