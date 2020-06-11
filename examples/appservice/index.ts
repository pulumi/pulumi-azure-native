import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";

const resourceGroup = new azurerm.core.ResourceGroup("rg", {
    resourceGroupName: "azurerm-appservice",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

const storageAccount = new azurerm.storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.resourceGroupName,
    accountName: "pulumiassa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
});

const appServicePlan  = new azurerm.web.AppServicePlan("asp", {
    resourceGroupName: resourceGroup.resourceGroupName,
    name: "appservice-plan",
    location: "westus2",
    kind: "App",
    sku: {
        name: "B1",
        tier: "Basic",
    },
});

const storageContainer = new azurerm.storage.StorageAccountBlobServiceContainer("c", {
    resourceGroupName: resourceGroup.resourceGroupName,
    accountName: storageAccount.accountName,
    containerName: "files",
    properies: {
        publicAccess: "none",
    },
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

const appInsights = new azurerm.insights.Component("ai", {
    resourceGroupName: resourceGroup.resourceGroupName,
    location: "westus2",
    resourceName: "pulumi-as-ai",
    kind: "web",
    properties: {
        Application_Type: "web",
    },
});

const username = "pulumi";
const pwd = "Not2S3cure!?";

const sqlServer = new azurerm.sql.Server("sql", {
    resourceGroupName: resourceGroup.resourceGroupName,
    location: "westus2",
    serverName: "pulumi-as-sql",
    properties: {
        administratorLogin: username,
        administratorLoginPassword: pwd,
        version: "12.0",
    },
});

const database = new azurerm.sql.ServerDatabase("db", {
    resourceGroupName: resourceGroup.resourceGroupName,
    location: "westus2",
    serverName: sqlServer.serverName,
    databaseName: "db",
    properties: {
        requestedServiceObjectiveName: "S0",
    }
});

const app = new azurerm.web.AppService("as", {
    resourceGroupName: resourceGroup.resourceGroupName,
    location: "westus2",
    name: "pulumi-rm-as",

    properties: {
        appServicePlanId: appServicePlan.id,
        siteConfig: {
            appSettings: {
                // TODO: "WEBSITE_RUN_FROM_ZIP": codeBlobUrl,
                "ApplicationInsights:InstrumentationKey": appInsights.properties.InstrumentationKey,
            },
            connectionStrings: [{
                name: "db",
                value:
                    pulumi.all([sqlServer.serverName, database.databaseName]).apply(([server, db]) =>
                        `Server=tcp:${server}.database.windows.net;initial catalog=${db};user ID=${username};password=${pwd};Min Pool Size=0;Max Pool Size=30;Persist Security Info=true;`),
                type: "SQLAzure",
            }],    
        },
    }
});

export const endpoint = pulumi.interpolate `https://${app.properties.defaultHostName}`;
