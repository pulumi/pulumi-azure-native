import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const availabilitySets_pulumirancher_avset_nameParam = config.get("availabilitySets_pulumirancher_avset_nameParam") || "pulumirancher-avset";
const availabilitySetResource = new azure_nextgen.compute.v20190701.AvailabilitySet("availabilitySetResource", {
    availabilitySetName: availabilitySets_pulumirancher_avset_nameParam,
    location: "westus2",
    platformFaultDomainCount: 1,
    platformUpdateDomainCount: 1,
    virtualMachines: [{}],
});
const extensions_DockerExtension_caParam = config.require("extensions_DockerExtension_caParam");
const extensions_DockerExtension_certParam = config.require("extensions_DockerExtension_certParam");
const extensions_DockerExtension_keyParam = config.require("extensions_DockerExtension_keyParam");
const virtualMachines_pulumirancher_nameParam = config.get("virtualMachines_pulumirancher_nameParam") || "pulumirancher";
const extensionResource = undefined;
const extensions_enablevmaccess_expirationParam = config.require("extensions_enablevmaccess_expirationParam");
const extensions_enablevmaccess_passwordParam = config.require("extensions_enablevmaccess_passwordParam");
const extensions_enablevmaccess_remove_userParam = config.require("extensions_enablevmaccess_remove_userParam");
const extensions_enablevmaccess_reset_sshParam = config.require("extensions_enablevmaccess_reset_sshParam");
const extensions_enablevmaccess_ssh_keyParam = config.require("extensions_enablevmaccess_ssh_keyParam");
const extensions_enablevmaccess_usernameParam = config.require("extensions_enablevmaccess_usernameParam");
const extensionResource0 = undefined;
const networkInterfaces_pulumirancher_nic_nameParam = config.get("networkInterfaces_pulumirancher_nic_nameParam") || "pulumirancher-nic";
const networkInterfaceResource = new azure_nextgen.network.v20200501.NetworkInterface("networkInterfaceResource", {
    dnsSettings: {
        dnsServers: [],
    },
    enableAcceleratedNetworking: false,
    enableIPForwarding: false,
    ipConfigurations: [{
        name: "ipconfig1",
        primary: true,
        privateIPAddress: "192.168.254.4",
        privateIPAddressVersion: "IPv4",
        privateIPAllocationMethod: "Dynamic",
        publicIPAddress: {
            id: undefined,
        },
        subnet: {
            id: undefined,
        },
    }],
    location: "westus2",
    networkInterfaceName: networkInterfaces_pulumirancher_nic_nameParam,
});
const networkSecurityGroups_pulumirancher_nsg_nameParam = config.get("networkSecurityGroups_pulumirancher_nsg_nameParam") || "pulumirancher-nsg";
const networkSecurityGroupResource = new azure_nextgen.network.v20200501.NetworkSecurityGroup("networkSecurityGroupResource", {
    location: "westus2",
    networkSecurityGroupName: networkSecurityGroups_pulumirancher_nsg_nameParam,
    securityRules: [
        {
            access: "Allow",
            description: "SSH",
            destinationAddressPrefix: "*",
            destinationAddressPrefixes: [],
            destinationPortRange: "22",
            destinationPortRanges: [],
            direction: "Inbound",
            name: "SSH",
            priority: 100,
            protocol: "*",
            sourceAddressPrefix: "*",
            sourceAddressPrefixes: [],
            sourcePortRange: "*",
            sourcePortRanges: [],
        },
        {
            access: "Allow",
            description: "Rancher-HTTPS",
            destinationAddressPrefix: "*",
            destinationAddressPrefixes: [],
            destinationPortRange: "8443",
            destinationPortRanges: [],
            direction: "Inbound",
            name: "Rancher",
            priority: 200,
            protocol: "Tcp",
            sourceAddressPrefix: "*",
            sourceAddressPrefixes: [],
            sourcePortRange: "*",
            sourcePortRanges: [],
        },
        {
            access: "Allow",
            description: "Docker",
            destinationAddressPrefix: "*",
            destinationAddressPrefixes: [],
            destinationPortRange: "2375",
            destinationPortRanges: [],
            direction: "Inbound",
            name: "Docker",
            priority: 300,
            protocol: "Tcp",
            sourceAddressPrefix: "*",
            sourceAddressPrefixes: [],
            sourcePortRange: "*",
            sourcePortRanges: [],
        },
    ],
});
const publicIPAddresses_pulumirancher_pip1_nameParam = config.get("publicIPAddresses_pulumirancher_pip1_nameParam") || "pulumirancher-pip1";
const publicIPAddressResource = new azure_nextgen.network.v20200501.PublicIPAddress("publicIPAddressResource", {
    dnsSettings: {
        domainNameLabel: "pulumirancher",
        fqdn: "pulumirancher.westus2.cloudapp.azure.com",
    },
    idleTimeoutInMinutes: 4,
    ipAddress: "52.175.229.132",
    ipTags: [],
    location: "westus2",
    publicIPAddressVersion: "IPv4",
    publicIPAllocationMethod: "Static",
    publicIpAddressName: publicIPAddresses_pulumirancher_pip1_nameParam,
    resourceGroupName: resource
});
const securityRuleResource = new azure_nextgen.network.v20200501.SecurityRule("securityRuleResource", {
    access: "Allow",
    description: "Docker",
    destinationAddressPrefix: "*",
    destinationAddressPrefixes: [],
    destinationPortRange: "2375",
    destinationPortRanges: [],
    direction: "Inbound",
    priority: 300,
    protocol: "Tcp",
    securityRuleName: "networkSecurityGroups_pulumirancher_nsg_nameParam}/Docker",
    sourceAddressPrefix: "*",
    sourceAddressPrefixes: [],
    sourcePortRange: "*",
    sourcePortRanges: [],
});
const securityRuleResource0 = new azure_nextgen.network.v20200501.SecurityRule("securityRuleResource0", {
    access: "Allow",
    description: "Rancher-HTTPS",
    destinationAddressPrefix: "*",
    destinationAddressPrefixes: [],
    destinationPortRange: "8443",
    destinationPortRanges: [],
    direction: "Inbound",
    priority: 200,
    protocol: "Tcp",
    securityRuleName: "networkSecurityGroups_pulumirancher_nsg_nameParam}/Rancher",
    sourceAddressPrefix: "*",
    sourceAddressPrefixes: [],
    sourcePortRange: "*",
    sourcePortRanges: [],
});
const securityRuleResource1 = new azure_nextgen.network.v20200501.SecurityRule("securityRuleResource1", {
    access: "Allow",
    description: "SSH",
    destinationAddressPrefix: "*",
    destinationAddressPrefixes: [],
    destinationPortRange: "22",
    destinationPortRanges: [],
    direction: "Inbound",
    priority: 100,
    protocol: "*",
    securityRuleName: "networkSecurityGroups_pulumirancher_nsg_nameParam}/SSH",
    sourceAddressPrefix: "*",
    sourceAddressPrefixes: [],
    sourcePortRange: "*",
    sourcePortRanges: [],
});
const virtualNetworks_pulumirancher_vnet_nameParam = config.get("virtualNetworks_pulumirancher_vnet_nameParam") || "pulumirancher-vnet";
const subnetResource = new azure_nextgen.network.v20200501.Subnet("subnetResource", {
    addressPrefix: "192.168.254.0/24",
    delegations: [],
    networkSecurityGroup: {},
    privateEndpointNetworkPolicies: "Enabled",
    privateLinkServiceNetworkPolicies: "Enabled",
    subnetName: "virtualNetworks_pulumirancher_vnet_nameParam}/pulumirancher-subnet",
});
const virtualMachineResource = new azure_nextgen.compute.v20190701.VirtualMachine("virtualMachineResource", {
    availabilitySet: {},
    hardwareProfile: {
        vmSize: "Standard_B2s",
    },
    location: "westus2",
    networkProfile: {
        networkInterfaces: [{}],
    },
    osProfile: {
        adminUsername: "pulumibot",
        computerName: virtualMachines_pulumirancher_nameParam,
        linuxConfiguration: {
            disablePasswordAuthentication: false,
        },
        secrets: [],
    },
    storageProfile: {
        dataDisks: [],
        imageReference: {
            offer: "UbuntuServer",
            publisher: "Canonical",
            sku: "18.04-LTS",
            version: "latest",
        },
        osDisk: {
            caching: "ReadWrite",
            createOption: "FromImage",
            diskSizeGB: 30,
            managedDisk: {
                storageAccountType: "Standard_LRS",
            },
            name: "virtualMachines_pulumirancher_nameParam}_Os_Disk_1",
            osType: "Linux",
        },
    },
    vmName: virtualMachines_pulumirancher_nameParam,
});
const virtualNetworkResource = new azure_nextgen.network.v20200501.VirtualNetwork("virtualNetworkResource", {
    addressSpace: {
        addressPrefixes: ["192.168.254.0/24"],
    },
    enableDdosProtection: false,
    enableVmProtection: false,
    location: "westus2",
    subnets: [{
        addressPrefix: "192.168.254.0/24",
        delegations: [],
        name: "pulumirancher-subnet",
        networkSecurityGroup: {
            id: undefined,
        },
        privateEndpointNetworkPolicies: "Enabled",
        privateLinkServiceNetworkPolicies: "Enabled",
    }],
    virtualNetworkName: virtualNetworks_pulumirancher_vnet_nameParam,
    virtualNetworkPeerings: [],
});
