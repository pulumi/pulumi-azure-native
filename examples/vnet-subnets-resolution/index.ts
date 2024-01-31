import * as resources from "@pulumi/azure-native/resources";
import * as network from "@pulumi/azure-native/network";

const resourceGroup = new resources.ResourceGroup("vnet-subnets-rg");

// This network has subnets defined inline.
const inlineVnet = new network.VirtualNetwork("inline", {
    resourceGroupName: resourceGroup.name,
    addressSpace: { addressPrefixes: ["10.4.0.0/16"] },
    tags: {
        "step": "1",
    },
    subnets: [
      { name: "default", addressPrefix: "10.4.1.0/24" },
      { name: "second", addressPrefix: "10.4.2.0/24" },
      { name: "third", addressPrefix: "10.4.3.0/24" },
    ]
});

// Allocate a resource in a subnet so that the subnet can't be destroyed accidentally.
// If provider tries to destroy the subnet, the test will fail.
new network.NetworkInterface("inline-nic", {
  resourceGroupName: resourceGroup.name,
  ipConfigurations: [{
      name: "cfg",
      subnet: { id: inlineVnet.subnets.apply(subnet => subnet![0].id!) },
      privateIPAllocationMethod: network.IPAllocationMethod.Dynamic,
  }],
});

const externalVnet = new network.VirtualNetwork("external", {
  resourceGroupName: resourceGroup.name,
  addressSpace: { addressPrefixes: ["10.5.0.0/16"] },
  tags: {
      "step": "1",
  },
// Can be removed after https://github.com/pulumi/pulumi-azure-native/issues/3049 is fixed.
}, { ignoreChanges: ["subnets"] });

const externalSubnet = new network.Subnet("default", {
    resourceGroupName: resourceGroup.name,
    virtualNetworkName: externalVnet.name,
    addressPrefix: "10.5.1.0/24",
});

const second = new network.Subnet("second", {
  resourceGroupName: resourceGroup.name,
  virtualNetworkName: externalVnet.name,
  addressPrefix: "10.5.2.0/24",
}, {dependsOn: externalSubnet});

new network.Subnet("third", {
  resourceGroupName: resourceGroup.name,
  virtualNetworkName: externalVnet.name,
  addressPrefix: "10.5.3.0/24",
}, {dependsOn: second});

new network.NetworkInterface("external-nic", {
  resourceGroupName: resourceGroup.name,
  ipConfigurations: [{
      name: "cfg",
      subnet: {
        id: externalSubnet.id,
      },
      privateIPAllocationMethod: network.IPAllocationMethod.Dynamic,
  }],
});
