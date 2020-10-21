import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const cDNSkuParam = config.get("cDNSkuParam") || "Standard_Microsoft";
const endpointNameParam = config.require("endpointNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const resourceGroupVar = azure_nextgen.resources.latest.getResourceGroup({
    resourceGroupName: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);
const originUrlParam = config.require("originUrlParam");
const profileNameParam = config.require("profileNameParam");
const profileResource = new azure_nextgen.cdn.v20190415.Profile("profileResource", {
    location: locationParam,
    profileName: profileNameParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: cDNSkuParam,
    },
});
const endpointResource = new azure_nextgen.cdn.v20190415.Endpoint("endpointResource", {
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
                        "@odata.type": "#Microsoft.Azure.Cdn.Models.DeliveryRuleUrlRewriteActionParameters",
                        destination: "/mobile",
                        sourcePattern: "/standard",
                    },
                }],
                conditions: [{
                    name: "IsDevice",
                    parameters: {
                        "@odata.type": "#Microsoft.Azure.Cdn.Models.DeliveryRuleIsDeviceConditionParameters",
                        matchValues: ["Mobile"],
                        operator: "Equal",
                    },
                }],
                name: "PathRewriteBasedOnDeviceMatchCondition",
                order: "1",
            },
            {
                actions: [{
                    name: "UrlRedirect",
                    parameters: {
                        "@odata.type": "#Microsoft.Azure.Cdn.Models.DeliveryRuleUrlRedirectActionParameters",
                        destinationProtocol: "Https",
                        redirectType: "Found",
                    },
                }],
                conditions: [{
                    name: "RequestScheme",
                    parameters: {
                        "@odata.type": "#Microsoft.Azure.Cdn.Models.DeliveryRuleRequestSchemeConditionParameters",
                        matchValues: ["HTTP"],
                        operator: "Equal",
                    },
                }],
                name: "HttpVersionBasedRedirect",
                order: "2",
            },
        ],
    },
    endpointName: `${profileNameParam}/${endpointNameParam}`,
    isCompressionEnabled: true,
    isHttpAllowed: true,
    isHttpsAllowed: true,
    location: locationParam,
    originHostHeader: originUrlParam,
    origins: [{
        hostName: originUrlParam,
        name: "origin1",
    }],
    profileName: profileResource.name,
    queryStringCachingBehavior: "IgnoreQueryString",
    resourceGroupName: resourceGroupNameParam,
});
