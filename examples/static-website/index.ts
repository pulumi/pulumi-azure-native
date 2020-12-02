import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import { URL } from "url";

import * as cdn from "@pulumi/azure-nextgen/cdn/v20200415";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as storage from "@pulumi/azure-nextgen/storage/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
});

// Create a Storage Account for our static website
const storageAccount = new storage.StorageAccount("websitesa", {
    resourceGroupName: resourceGroup.name,
    accountName: randomString.result,
    location: "westus2",
    sku: {
        name: "Standard_LRS",
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
const profile =  new cdn.Profile("website-cdn", {
    resourceGroupName: resourceGroup.name,
    profileName: randomString.result,
    location: "global",
    sku: {
          name: "Standard_Microsoft",
    },
});

const endpoint = new cdn.Endpoint("website-cdn-ep", {
    resourceGroupName: resourceGroup.name,
    profileName: profile.name,
    endpointName: randomString.result,
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
