import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as cdn from "@pulumi/azure-native/cdn";

// Create a resource group
const resourceGroup = new resources.ResourceGroup("resourceGroup");

// Create a CDN profile
const cdnProfile = new cdn.Profile("cdnProfile", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    sku: {
        name: "Standard_Microsoft",
    },
});

// Create a CDN endpoint
const cdnEndpoint = new cdn.Endpoint("cdnEndpoint", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    location: resourceGroup.location,
    originHostHeader: "www.contoso.com",
    origins: [{
        name: "contoso",
        hostName: "www.contoso.com",
    }],
});

// Create a custom domain
const customDomain = new cdn.CustomDomain("customDomain", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    endpointName: cdnEndpoint.name,
    hostName: "cdn.contoso.com",
});

// Enable HTTPS on the custom domain using CDN-managed certificate
const customDomainHttps = new cdn.CustomDomainHttps("customDomainHttps", {
    resourceGroupName: resourceGroup.name,
    profileName: cdnProfile.name,
    endpointName: cdnEndpoint.name,
    customDomainName: customDomain.name,
    httpsEnabled: true,
    certificateSource: "Cdn",
    protocolType: "ServerNameIndication",
    certificateType: "Shared",
    minimumTlsVersion: "TLS12",
});

// Export the custom domain name and HTTPS status
export const customDomainName = customDomain.name;
export const httpsEnabled = customDomainHttps.httpsEnabled;