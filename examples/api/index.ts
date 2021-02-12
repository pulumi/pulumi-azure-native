import * as random from "@pulumi/random";
import * as apimanagement from "@pulumi/azure-nextgen/apimanagement";
import * as resources from "@pulumi/azure-nextgen/resources";
import * as web from "@pulumi/azure-nextgen/web";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
    number: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
});

new web.Connection("conn", {
    resourceGroupName: resourceGroup.name,
    connectionName: "office365",
    location: resourceGroup.location,
    properties: {
        displayName: "office365",
        api: {
            id: "/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/providers/Microsoft.Web/locations/westus2/managedApis/office365",
        }
    },
});

const apiManagement = new apimanagement.ApiManagementService("apim", {
    resourceGroupName: resourceGroup.name,
    serviceName: randomString.result,
    sku: {
        name: "Consumption",
        capacity: 0,
    },
    publisherEmail: "drones@contoso.com",
    publisherName: "contoso",
});

const versionSet = new apimanagement.ApiVersionSet("dronestatusversionset", {
    resourceGroupName: resourceGroup.name,
    serviceName: apiManagement.name,
    versionSetId: "dronestatusversionset",
    displayName: "Drone Delivery API",
    versioningScheme: "Segment",
});

const api = new apimanagement.Api("dronedeliveryapiv1", {
    resourceGroupName: resourceGroup.name,
    serviceName: apiManagement.name,
    apiId: "dronedeliveryapiv1",
    displayName: "Drone Delivery API",
    description: "Drone Delivery API",
    path: "api",
    apiVersion: "v1",
    apiRevision: "1",
    apiVersionSetId: versionSet.id,
    protocols: ["https"],
});

const product = new apimanagement.Product("dronedeliveryprodapi", {
    resourceGroupName: resourceGroup.name,
    serviceName: apiManagement.name,
    productId: "dronedeliveryprodapi",
    displayName: "drone delivery product api",
    description: "drone delivery product api",
    terms: "terms for example product",
    subscriptionRequired: false,    
    state: "published",
});

const productApi = new apimanagement.ProductApi("dronedeliveryapiv1", {
    resourceGroupName: resourceGroup.name,
    serviceName: apiManagement.name,
    apiId: api.name,
    productId: product.name,
});
