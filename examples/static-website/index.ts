import * as pulumi from "@pulumi/pulumi";
import { URL } from "url";

import * as cdn from "@pulumi/azure-nextgen/cdn";
import * as resources from "@pulumi/azure-nextgen/resources";
import * as storage from "@pulumi/azure-nextgen/storage";

const resourceGroup = new resources.ResourceGroup("rg");

// Create a Storage Account for our static website
const storageAccount = new storage.StorageAccount("websitesa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
});

// Create a public endpoint for the storage account
const staticWebsite = new storage.StorageAccountStaticWebsite("staticsite", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    indexDocument: "index.html",
    error404Document: "404.html",
});

// Upload the files
["index.html", "404.html"].map(name =>
    new storage.Blob(name, {
        resourceGroupName: resourceGroup.name,
        accountName: storageAccount.name,
        containerName: staticWebsite.containerName,
        source: new pulumi.asset.FileAsset(`./wwwroot/${name}`),
        contentType: "text/html",
    }),
);

new storage.Blob("test.html", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: staticWebsite.containerName,
    source: new pulumi.asset.StringAsset("<h1>This is deployed as-is</h1>"),
    contentMd5: "+vC+ix/2YjgiJLA4GimBqg==",
    contentType: "text/html",
});

new storage.Blob("backup.zip", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: staticWebsite.containerName,
    source: new pulumi.asset.FileArchive("./wwwroot"),
    contentType: "text/html",
});

// A test for an append blob.
new storage.Blob("append-blob.dat", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    containerName: staticWebsite.containerName,
    type: storage.BlobType.Append,
});

// Web endpoint to the website
export const staticEndpoint = storageAccount.primaryEndpoints.web;
const staticHostname = staticEndpoint.apply(url => new URL(url).hostname);

// We can add a CDN in front of the website
const profile =  new cdn.Profile("website-cdn", {
    resourceGroupName: resourceGroup.name,
    location: "global",
    sku: {
          name: cdn.SkuName.Standard_Microsoft,
    },
});

const endpoint = new cdn.Endpoint("website-cdn-ep", {
    resourceGroupName: resourceGroup.name,
    location: "westus",
    profileName: profile.name,
    originHostHeader: staticHostname,
    origins: [{
        name: "blobstorage",
        hostName: staticHostname,
    }],        
});

// CDN endpoint to the website.
// Allow it some time after the deployment to get ready.
export const cdnEndpoint = pulumi.interpolate`https://${endpoint.hostName}/`;
