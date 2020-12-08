import * as random from "@pulumi/random";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as web from "@pulumi/azure-nextgen/web/v20200601";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
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
