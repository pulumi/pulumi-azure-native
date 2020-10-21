import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const cDNSkuParam = "Standard_Microsoft";
const endpointNameParam = "updatedestroy-cdn-ep"
const resourceGroupNameParam = "cdnendpoint-rg"
const originUrlParam = "www.cdnendpoint.com"
const profileNameParam = "cdnendpoint-profile"
const locationParam = "westus"

const resourceGroup = new azure_nextgen.resources.latest.ResourceGroup('resourceGroup', {
    resourceGroupName: resourceGroupNameParam,
    location: locationParam
  });

const profileResource = new azure_nextgen.cdn.v20190415.Profile("profileResource", {
    location: resourceGroup.location,
    profileName: profileNameParam,
    resourceGroupName: resourceGroup.name,
    sku: {
        name: cDNSkuParam,
    },
});

const endpointResource = new azure_nextgen.cdn.v20190415.Endpoint("endpointResource", {
    profileName: profileResource.name,
    contentTypesToCompress: [
        "text/plain",
        "text/html",
        "text/css",
        "application/x-javascript",
        "text/javascript",
    ],
    deliveryPolicy: {
        description: "Rewrite and Redirect",
        rules: [
            {
                actions: [{
                    name: "UrlRewrite",
                    parameters: {
                        odataType: "#Microsoft.Azure.Cdn.Models.DeliveryRuleUrlRewriteActionParameters",
                        destination: "/mobile",
                        sourcePattern: "/standard",
                    },
                }],
                conditions: [{
                    name: "IsDevice",
                    parameters: {
                        odataType: "#Microsoft.Azure.Cdn.Models.DeliveryRuleIsDeviceConditionParameters",
                        matchValues: ["Mobile"],
                        operator: "Equal",
                    },
                }],
                name: "PathRewriteBasedOnDeviceMatchCondition",
                order: 1,
            },
            {
                actions: [{
                    name: "UrlRedirect",
                    parameters: {
                        odataType: "#Microsoft.Azure.Cdn.Models.DeliveryRuleUrlRedirectActionParameters",
                        destinationProtocol: "Https",
                        redirectType: "Found",
                    },
                }],
                conditions: [{
                    name: "RequestScheme",
                    parameters: {
                        odataType: "#Microsoft.Azure.Cdn.Models.DeliveryRuleRequestSchemeConditionParameters",
                        matchValues: ["HTTP"],
                        operator: "Equal",
                    },
                }],
                name: "HttpVersionBasedRedirect",
                order: 2,
            },
        ],
    },
    endpointName: endpointNameParam,
    isCompressionEnabled: true,
    isHttpAllowed: true,
    isHttpsAllowed: true,
    location: resourceGroup.location,
    originHostHeader: originUrlParam,
    origins: [{
        hostName: originUrlParam,
        name: "origin1",
    }],
    queryStringCachingBehavior: "IgnoreQueryString",
    resourceGroupName: resourceGroup.name,
});



