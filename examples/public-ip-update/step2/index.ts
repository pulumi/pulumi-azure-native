import * as resources from "@pulumi/azure-native/resources";
import * as network from "@pulumi/azure-native/network";

const resourceGroup = new resources.ResourceGroup("resource-group");

new network.PublicIPAddress("ip", {
  resourceGroupName: resourceGroup.name,
  publicIPAllocationMethod: "Static",
  // Adding the below lines causes a diff on a map
  sku: {
    name: "Standard",
    tier: "Regional",
  },
});
