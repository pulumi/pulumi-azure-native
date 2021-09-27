// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

import * as insights from "@pulumi/azure-native/insights";
import * as resources from "@pulumi/azure-native/resources";
import * as security from "@pulumi/azure-native/security";
import * as sql from "@pulumi/azure-native/sql";
import * as storage from "@pulumi/azure-native/storage";
import * as web from "@pulumi/azure-native/web";

const resourceGroup = new resources.ResourceGroup("rg");

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
});

new security.AdvancedThreatProtection("atp", {
    settingName: "current",
    resourceId: storageAccount.id,
    isEnabled: true,
});

const appServicePlan  = new web.AppServicePlan("asp", {
    resourceGroupName: resourceGroup.name,
    kind: "App",
    sku: {
        name: "B1",
        tier: "Basic",
    },
});

const storageContainer = new storage.BlobContainer("files", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    publicAccess: storage.PublicAccess.None,
});

const blob = new storage.Blob("wwwroot", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: storageContainer.name,
    type: "Block",
    source: new pulumi.asset.FileArchive("wwwroot"),
});

const appInsights = new insights.Component("ai", {
    resourceGroupName: resourceGroup.name,
    kind: "web",
    applicationType: insights.ApplicationType.Web,
});

new insights.ComponentCurrentBillingFeature("ccbf", {
    resourceGroupName: resourceGroup.name,
    resourceName: appInsights.name,
    currentBillingFeatures: ["Basic"],
    dataVolumeCap: {
        cap: 2,
    },
});

const username = "pulumi";
const pwd = "Not2S3cure!?";

const sqlServer = new sql.Server("sql", {
    resourceGroupName: resourceGroup.name,
    administratorLogin: username,
    administratorLoginPassword: pwd,
    version: "12.0",
});

const database = new sql.Database("db", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    sku: {
        name: "S0",
    },
});

new sql.DatabaseSecurityAlertPolicy("dsal", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    databaseName: database.name,
    securityAlertPolicyName: "default",
    state: "Enabled",
});

new sql.TransparentDataEncryption("current", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    databaseName: database.name,
    status: sql.TransparentDataEncryptionStatus.Enabled,
});

new sql.ServerAdvisor("ForceLastGoodPlan", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    autoExecuteStatus: sql.AutoExecuteStatus.Enabled,
});

new sql.EncryptionProtector("current", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    serverKeyType: "ServiceManaged",
});

const app = new web.WebApp("as", {
    resourceGroupName: resourceGroup.name,
    serverFarmId: appServicePlan.id,
});

new web.WebAppMetadata("meta", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    properties: {
        CURRENT_STACK: "dotnetcore",
    },
});

new web.WebAppApplicationSettings("settings", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    properties: {
        "ApplicationInsights:InstrumentationKey": appInsights.instrumentationKey?.apply(v => v!),
        "test": "Test"
    },
});

new web.WebAppConnectionStrings("conns", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    properties: {
        db: {
            value: pulumi.all([sqlServer.name, database.name]).apply(([server, db]) =>
                `Server=tcp:${server}.database.windows.net;initial catalog=${db};user ID=${username};password=${pwd};Min Pool Size=0;Max Pool Size=30;Persist Security Info=true;`),
            type: "SQLAzure",
        },
    },
});

new web.WebAppAuthSettingsV2("auth", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    httpSettings: {
        requireHttps: true,
    },
});

new web.WebAppSourceControl("sc", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    repoUrl: "https://github.com/octocat/Hello-World",
    isManualIntegration: true,
});

new web.WebAppDiagnosticLogsConfiguration("wadlc", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    applicationLogs: {
        fileSystem: { level: "Verbose" },
    },
});

export const endpoint = pulumi.interpolate `https://${app.defaultHostName}`;
