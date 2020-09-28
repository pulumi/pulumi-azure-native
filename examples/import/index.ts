import * as pulumi from "@pulumi/pulumi";

import * as network from "@pulumi/azure-nextgen/network/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as manual from "./manualResource";

function hashCode(s: string) {
    var hash = 0;
    for (var i = 0; i < s.length; i++) {
        var character = s.charCodeAt(i);
        hash = ((hash<<5)-hash)+character;
        hash = hash & hash; // Convert to 32bit integer
    }
    return Math.abs(hash);
}

async function main() {
    const name = pulumi.getStack().toLowerCase();
    const resourceGroupName = `${name}-rg`;
    const virtualNetworkName = `${name}-vnet`;
    const location = "westus2";
    const virtualNetworkRange = `10.${hashCode(name)%250 + 4}.0.0`;

    // Create resources with Azure SDKs, so that we had something to import.
    const subscriptionID = await manual.createResources({ resourceGroupName, location, virtualNetworkName, virtualNetworkRange });

    // Now, import the same resources to Pulumi. pulumi destroy will clean them up.
    new resources.ResourceGroup(`${name}-rg`, {
        resourceGroupName,
        location,
    }, { import: `/subscriptions/${subscriptionID}/resourceGroups/${resourceGroupName}` });

    const vnetId = `/subscriptions/${subscriptionID}/resourceGroups/${resourceGroupName}/providers/Microsoft.Network/virtualNetworks/${virtualNetworkName}`;
    const subnetId = `${vnetId}/subnets/default`;
    new network.VirtualNetwork("vnet", {
        resourceGroupName,
        virtualNetworkName,
        location,
        addressSpace: {
            addressPrefixes: [`${virtualNetworkRange}/16`],
        },
        subnets: [{
            id: subnetId,
            name: "default",
            addressPrefix: `${virtualNetworkRange}/24`,
            privateEndpointNetworkPolicies: "Enabled",
            privateLinkServiceNetworkPolicies: "Enabled",
        }],
    }, { import: vnetId });
}

module.exports = main();
