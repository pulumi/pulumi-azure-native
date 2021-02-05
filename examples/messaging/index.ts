import * as random from "@pulumi/random";
import * as eventhub from "@pulumi/azure-nextgen/eventhub/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
    number: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
});

const eventHubNamespace = new eventhub.Namespace("ehns", {
    resourceGroupName: resourceGroup.name,
    namespaceName: randomString.result,
    location: resourceGroup.location,
    sku: {
        name: eventhub.SkuName.Standard,
    },
});

const eventHub = new eventhub.EventHub("eh", {
    resourceGroupName: resourceGroup.name,
    namespaceName: eventHubNamespace.name,
    eventHubName: "testeh",
    messageRetentionInDays: 1,
    partitionCount: 4,
});

new eventhub.NamespaceNetworkRuleSet("rs", {
    resourceGroupName: resourceGroup.name,
    namespaceName: eventHubNamespace.name,
    defaultAction: eventhub.DefaultAction.Allow,
});
