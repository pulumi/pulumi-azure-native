import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

import * as insights from "@pulumi/azure-nextgen/insights/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as sql from "@pulumi/azure-nextgen/sql/latest";
import * as storage from "@pulumi/azure-nextgen/storage/latest";
import * as web from "@pulumi/azure-nextgen/web/v20200601";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
});

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    accountName: randomString.result,
    location: "westus2",
    sku: {
        name: "Standard_LRS",
    },
    kind: "StorageV2",
});

const appServicePlan  = new web.AppServicePlan("asp", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    location: "westus2",
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
    publicAccess: "None",
});

// TODO: blobs are not supported
// const blob = new azure.storage.Blob(`${prefix}-b`, {
//     storageAccountName: storageAccount.name,
//     storageContainerName: storageContainer.name,
//     type: "Block",

//     source: new pulumi.asset.FileArchive("wwwroot"),
// });

// TODO: invokes are not supported yet
// const codeBlobUrl = azure.storage.signedBlobReadUrl(blob, storageAccount);

const appInsights = new insights.Component("ai", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    resourceName: randomString.result,
    kind: "web",
    applicationType: "web",
});

const username = "pulumi";
const pwd = "Not2S3cure!?";

const sqlServer = new sql.Server("sql", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    serverName: randomString.result,
    administratorLogin: username,
    administratorLoginPassword: pwd,
    version: "12.0",
});

const database = new sql.Database("db", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    serverName: sqlServer.name,
    databaseName: "db",
    requestedServiceObjectiveName: "S0",
});

new sql.TransparentDataEncryption("tde", {
    resourceGroupName: resourceGroup.name,
    transparentDataEncryptionName: "current",
    serverName: sqlServer.name,
    databaseName: database.name,
    status: "Enabled",
});

const app = new web.WebApp("as", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    name: randomString.result,

    serverFarmId: appServicePlan.id,
    siteConfig: {
        appSettings: [
            // TODO: "WEBSITE_RUN_FROM_ZIP": codeBlobUrl,
            {
                name: "ApplicationInsights:InstrumentationKey",
                value: appInsights.instrumentationKey?.apply(v => v!),
            },
        ],
        connectionStrings: [{
            name: "db",
            connectionString:
                pulumi.all([sqlServer.name, database.name]).apply(([server, db]) =>
                    `Server=tcp:${server}.database.windows.net;initial catalog=${db};user ID=${username};password=${pwd};Min Pool Size=0;Max Pool Size=30;Persist Security Info=true;`),
            type: "SQLAzure",
        }],    
    },
});

new web.WebAppSourceControl("sc", {
    resourceGroupName: resourceGroup.name,
    name: app.name,
    repoUrl: "https://github.com/octocat/Hello-World",
    isManualIntegration: true,
});

export const endpoint = pulumi.interpolate `https://${app.defaultHostName}`;
