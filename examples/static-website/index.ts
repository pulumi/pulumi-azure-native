import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";
import { URL } from "url";

const resourceGroup = new azurerm.resources.latest.ResourceGroup("rg", {
    resourceGroupName: "azurerm-static-website",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

// Create a Storage Account for our static website
const storageAccount = new azurerm.storage.latest.StorageAccount("websitesa", {
    resourceGroupName: resourceGroup.name,
    accountName: "pulumiswsa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
    // Apparently, this is not supported via ARM.
    // staticWebsite: {
    //     indexDocument: "index.html",
    // },
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
export const staticEndpoint = storageAccount.primaryEndpoints.web;
// TODO: figure out why this output is resolved to undefined during the preview, which breaks the URL constructor call.
const staticHostname = staticEndpoint.apply(url => url ? new URL(url).hostname : "<preview>");

// We can add a CDN in front of the website
const cdn =  new azurerm.cdn.latest.Profile("website-cdn", {
    resourceGroupName: resourceGroup.name,
    profileName: "pulumi-static-website",
    location: "global",
    sku: {
          name: "Standard_Microsoft",
    },
});

const endpoint = new azurerm.cdn.latest.Endpoint("website-cdn-ep", {
    resourceGroupName: resourceGroup.name,
    profileName: cdn.name,
    endpointName: "pulumi-static-website-ep",
    location: "westus",
    originHostHeader: staticHostname,
    origins: [{
        name: "blobstorage",
        hostName: staticHostname,
    }],        
});

// CDN endpoint to the website.
// Allow it some time after the deployment to get ready.
export const cdnEndpoint = pulumi.interpolate`https://${endpoint.hostName}/`;
