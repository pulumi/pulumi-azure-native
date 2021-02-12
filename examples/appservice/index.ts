import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

import * as insights from "@pulumi/azure-nextgen/insights";
import * as resources from "@pulumi/azure-nextgen/resources";
import * as sql from "@pulumi/azure-nextgen/sql/v20200801preview";
import * as sqltde from "@pulumi/azure-nextgen/sql";
import * as storage from "@pulumi/azure-nextgen/storage";
import * as web from "@pulumi/azure-nextgen/web";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
});

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    accountName: randomString.result,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
});

const appServicePlan  = new web.AppServicePlan("asp", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    kind: "App",
    sku: {
        name: "B1",
        tier: "Basic",
    },
});

const storageContainer = new storage.BlobContainer("c", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: "files",
    publicAccess: storage.PublicAccess.None,
});

const blob = new storage.Blob("b", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: storageContainer.name,
    blobName: "wwwroot",
    type: "Block",

    source: new pulumi.asset.FileArchive("wwwroot"),
});

// TODO: invokes are not supported yet
// const codeBlobUrl = azure.storage.signedBlobReadUrl(blob, storageAccount);

const appInsights = new insights.Component("ai", {
    resourceGroupName: resourceGroup.name,
    resourceName: randomString.result,
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
    serverName: randomString.result,
    administratorLogin: username,
    administratorLoginPassword: pwd,
    version: "12.0",
});

const database = new sql.Database("db", {
    resourceGroupName: resourceGroup.name,
    serverName: sqlServer.name,
    databaseName: "db",
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

new sqltde.TransparentDataEncryption("tde", {
    resourceGroupName: resourceGroup.name,
    transparentDataEncryptionName: "current",
    serverName: sqlServer.name,
    databaseName: database.name,
    status: sqltde.TransparentDataEncryptionStatus.Enabled,
});

const app = new web.WebApp("as", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
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
