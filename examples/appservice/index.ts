import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";

const resourceGroup = new azurerm.resources.latest.ResourceGroup("rg", {
    resourceGroupName: "azurerm-appservice",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

const storageAccount = new azurerm.storage.latest.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    accountName: "pulumiassa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
});

const appServicePlan  = new azurerm.web.latest.AppServicePlan("asp", {
    resourceGroupName: resourceGroup.name,
    name: "appservice-plan",
    location: "westus2",
    kind: "App",
    sku: {
        name: "B1",
        tier: "Basic",
    },
});

const storageContainer = new azurerm.storage.latest.BlobContainer("c", {
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

const appInsights = new azurerm.insights.latest.Component("ai", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    resourceName: "pulumi-as-ai",
    kind: "web",
    applicationType: "web",
});

const username = "pulumi";
const pwd = "Not2S3cure!?";

const sqlServer = new azurerm.sql.latest.Server("sql", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    serverName: "pulumi-as-sql",
    administratorLogin: username,
    administratorLoginPassword: pwd,
    version: "12.0",
});

const database = new azurerm.sql.latest.Database("db", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    serverName: sqlServer.name,
    databaseName: "db",
    requestedServiceObjectiveName: "S0",
});

const app = new azurerm.web.latest.WebApp("as", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    name: "pulumi-rm-as",

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

export const endpoint = pulumi.interpolate `https://${app.defaultHostName}`;
