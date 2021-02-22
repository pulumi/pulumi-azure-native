import * as eventhub from "@pulumi/azure-native/eventhub";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("rg");

const eventHubNamespace = new eventhub.Namespace("ehns", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    sku: {
        name: eventhub.SkuName.Standard,
    },
});

const eventHub = new eventhub.EventHub("eh", {
    resourceGroupName: resourceGroup.name,
    namespaceName: eventHubNamespace.name,
    messageRetentionInDays: 1,
    partitionCount: 4,
});

new eventhub.NamespaceNetworkRuleSet("rs", {
    resourceGroupName: resourceGroup.name,
    namespaceName: eventHubNamespace.name,
    defaultAction: eventhub.DefaultAction.Allow,
});
