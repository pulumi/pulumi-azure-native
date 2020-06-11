import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";
import { URL } from "url";

const resourceGroup = new azurerm.core.ResourceGroup("rg", {
    resourceGroupName: "azurerm-static-website",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

// Create a Storage Account for our static website
const storageAccount = new azurerm.storage.StorageAccount("websitesa", {
    resourceGroupName: resourceGroup.resourceGroupName,
    accountName: "pulumiswsa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
    properties: {
        staticWebsite: {
            indexDocument: "index.html",
        },
    }
});

// Upload the files
// TODO: Blobs are not part of ARM API.
// ["index.html", "404.html"].map(name =>
//     new azure.storage.Blob(name, {
//         name,
//         storageAccountName: storageAccount.name,
//         storageContainerName: "$web",
//         type: "Block",
//         source: new pulumi.asset.FileAsset(`./wwwroot/${name}`),
//         contentType: "text/html",
//     }),
// );

// Web endpoint to the website
export const staticEndpoint = storageAccount.properties.primaryEndpoints.web;
const staticHostname = staticEndpoint.apply(url => new URL(url).hostname);

// We can add a CDN in front of the website
const cdn =  new azurerm.cdn.Profile("website-cdn", {
    resourceGroupName: resourceGroup.resourceGroupName,
    profileName: "pulumi-static-website",
    location: "global",
    sku: {
          name: "Standard_Microsoft",
    },
});

const endpoint = new azurerm.cdn.ProfileEndpoint("website-cdn-ep", {
    resourceGroupName: resourceGroup.resourceGroupName,
    profileName: cdn.profileName,
    endpointName: "pulumi-static-website-ep",
    location: "westus",
    properties: {
        originHostHeader: staticHostname,
        origins: [{
            name: "blobstorage",
            properties: {
                hostName: staticHostname,
            },
        }],        
    },
});

// CDN endpoint to the website.
// Allow it some time after the deployment to get ready.
export const cdnEndpoint = pulumi.interpolate`https://${endpoint.properties.hostName}/`;
