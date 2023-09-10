import * as resources from "@pulumi/azure-native/resources";
import * as network from "@pulumi/azure-native/network";

const resourceGroup = new resources.ResourceGroup("vnet-subnets-rg");

const inlineVnet = new network.VirtualNetwork("inline", {
    resourceGroupName: resourceGroup.name,
    addressSpace: { addressPrefixes: ["10.4.0.0/16"] },
    tags: {
        "step": "2",
    },
    subnets: [
      { name: "default", addressPrefix: "10.4.1.0/24" },
      { name: "second", addressPrefix: "10.4.2.0/24" },
      { name: "third", addressPrefix: "10.4.3.0/24" },
      { name: "fourth", addressPrefix: "10.4.4.0/24" },
    ]
});

new network.NetworkInterface("inline-nic", {
  resourceGroupName: resourceGroup.name,
  ipConfigurations: [{
      name: "cfg",
      subnet: inlineVnet.subnets.apply(subnet => subnet![0]),
      privateIPAllocationMethod: network.IPAllocationMethod.Dynamic,
  }],
});

const externalVnet = new network.VirtualNetwork("external", {
  resourceGroupName: resourceGroup.name,
  addressSpace: { addressPrefixes: ["10.5.0.0/16"] },
  tags: {
      // Updating tags causes a resource update that should NOT try to delete subnets.
      "step": "2",
  },
});

const externalSubnet = new network.Subnet("default", {
    resourceGroupName: resourceGroup.name,
    virtualNetworkName: externalVnet.name,
    addressPrefix: "10.5.1.0/24",
});

new network.Subnet("second", {
  resourceGroupName: resourceGroup.name,
  virtualNetworkName: externalVnet.name,
  addressPrefix: "10.5.2.0/24",
});

new network.Subnet("third", {
  resourceGroupName: resourceGroup.name,
  virtualNetworkName: externalVnet.name,
  addressPrefix: "10.5.3.0/24",
});

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
