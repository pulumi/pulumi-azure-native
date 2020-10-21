import * as pulumi from "@pulumi/pulumi";
import * as azure_nextgen from "@pulumi/azure-nextgen";

const config = new pulumi.Config();
const availabilitySetsPulumirancherAvsetNameParam = config.get("availabilitySetsPulumirancherAvsetNameParam") || "pulumirancher-avset";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const virtualMachinesPulumirancherNameParam = config.get("virtualMachinesPulumirancherNameParam") || "pulumirancher";
const availabilitySetResource = new azure_nextgen.compute.v20190701.AvailabilitySet("availabilitySetResource", {
    availabilitySetName: availabilitySetsPulumirancherAvsetNameParam,
    location: "westus2",
    platformFaultDomainCount: 1,
    platformUpdateDomainCount: 1,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: "Aligned",
    },
});
const extensionsDockerExtensionCaParam = config.require("extensionsDockerExtensionCaParam");
const extensionsDockerExtensionCertParam = config.require("extensionsDockerExtensionCertParam");
const extensionsDockerExtensionKeyParam = config.require("extensionsDockerExtensionKeyParam");
const networkInterfacesPulumirancherNicNameParam = config.get("networkInterfacesPulumirancherNicNameParam") || "pulumirancher-nic";
const publicIPAddressesPulumirancherPip1NameParam = config.get("publicIPAddressesPulumirancherPip1NameParam") || "pulumirancher-pip1";
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
    publicIpAddressName: publicIPAddressesPulumirancherPip1NameParam,
    resourceGroupName: resourceGroupNameParam,
    sku: {
        name: "Basic",
    },
});
const networkSecurityGroupsPulumirancherNsgNameParam = config.get("networkSecurityGroupsPulumirancherNsgNameParam") || "pulumirancher-nsg";
const networkSecurityGroupResource = new azure_nextgen.network.v20200501.NetworkSecurityGroup("networkSecurityGroupResource", {
    location: "westus2",
    networkSecurityGroupName: networkSecurityGroupsPulumirancherNsgNameParam,
    resourceGroupName: resourceGroupNameParam,
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
const virtualNetworksPulumirancherVnetNameParam = config.get("virtualNetworksPulumirancherVnetNameParam") || "pulumirancher-vnet";
const virtualNetworkResource = new azure_nextgen.network.v20200501.VirtualNetwork("virtualNetworkResource", {
    addressSpace: {
        addressPrefixes: ["192.168.254.0/24"],
    },
    enableDdosProtection: false,
    enableVmProtection: false,
    location: "westus2",
    resourceGroupName: resourceGroupNameParam,
    subnets: [{
        addressPrefix: "192.168.254.0/24",
        delegations: [],
        name: "pulumirancher-subnet",
        networkSecurityGroup: {
            id: networkSecurityGroupResource.id,
        },
        privateEndpointNetworkPolicies: "Enabled",
        privateLinkServiceNetworkPolicies: "Enabled",
    }],
    virtualNetworkName: virtualNetworksPulumirancherVnetNameParam,
    virtualNetworkPeerings: [],
});
const subnetResource = new azure_nextgen.network.v20200501.Subnet("subnetResource", {
    addressPrefix: "192.168.254.0/24",
    delegations: [],
    networkSecurityGroup: {
        id: networkSecurityGroupResource.id,
    },
    privateEndpointNetworkPolicies: "Enabled",
    privateLinkServiceNetworkPolicies: "Enabled",
    resourceGroupName: resourceGroupNameParam,
    subnetName: `${virtualNetworksPulumirancherVnetNameParam}/pulumirancher-subnet`,
    virtualNetworkName: virtualNetworkResource.name,
});
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
            id: publicIPAddressResource.id,
        },
        subnet: {
            id: subnetResource.id,
        },
    }],
    location: "westus2",
    networkInterfaceName: networkInterfacesPulumirancherNicNameParam,
    resourceGroupName: resourceGroupNameParam,
});
const virtualMachineResource = new azure_nextgen.compute.v20190701.VirtualMachine("virtualMachineResource", {
    availabilitySet: {
        id: availabilitySetResource.id,
    },
    hardwareProfile: {
        vmSize: "Standard_B2s",
    },
    location: "westus2",
    networkProfile: {
        networkInterfaces: [{
            id: networkInterfaceResource.id,
        }],
    },
    osProfile: {
        adminUsername: "pulumibot",
        computerName: virtualMachinesPulumirancherNameParam,
        linuxConfiguration: {
            disablePasswordAuthentication: false,
        },
        secrets: [],
    },
    resourceGroupName: resourceGroupNameParam,
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
                id: "[resourceId('Microsoft.Compute/disks',concat(parameters('virtualMachines_pulumirancher_name'),'_Os_Disk_1'))]",
                storageAccountType: "Standard_LRS",
            },
            name: `${virtualMachinesPulumirancherNameParam}_Os_Disk_1`,
            osType: "Linux",
        },
    },
    vmName: virtualMachinesPulumirancherNameParam,
});
const extensionResource = new azure_nextgen.compute.v20190701.VirtualMachineExtension("extensionResource", {
    autoUpgradeMinorVersion: true,
    location: "westus2",
    protectedSettings: {
        certs: {
            ca: extensionsDockerExtensionCaParam,
            cert: extensionsDockerExtensionCertParam,
            key: extensionsDockerExtensionKeyParam,
        },
    },
    publisher: "Microsoft.Azure.Extensions",
    resourceGroupName: resourceGroupNameParam,
    settings: {
        docker: {
            port: "2375",
        },
    },
    type: "DockerExtension",
    typeHandlerVersion: "1.0",
    vmExtensionName: `${virtualMachinesPulumirancherNameParam}/DockerExtension`,
    vmName: virtualMachineResource.name,
});
const extensionsEnablevmaccessExpirationParam = config.require("extensionsEnablevmaccessExpirationParam");
const extensionsEnablevmaccessPasswordParam = config.require("extensionsEnablevmaccessPasswordParam");
const extensionsEnablevmaccessRemoveUserParam = config.require("extensionsEnablevmaccessRemoveUserParam");
const extensionsEnablevmaccessResetSshParam = config.require("extensionsEnablevmaccessResetSshParam");
const extensionsEnablevmaccessSshKeyParam = config.require("extensionsEnablevmaccessSshKeyParam");
const extensionsEnablevmaccessUsernameParam = config.require("extensionsEnablevmaccessUsernameParam");
const extensionResource0 = new azure_nextgen.compute.v20190701.VirtualMachineExtension("extensionResource0", {
    autoUpgradeMinorVersion: true,
    location: "westus2",
    protectedSettings: {
        expiration: extensionsEnablevmaccessExpirationParam,
        password: extensionsEnablevmaccessPasswordParam,
        remove_user: extensionsEnablevmaccessRemoveUserParam,
        reset_ssh: extensionsEnablevmaccessResetSshParam,
        ssh_key: extensionsEnablevmaccessSshKeyParam,
        username: extensionsEnablevmaccessUsernameParam,
    },
    publisher: "Microsoft.OSTCExtensions",
    resourceGroupName: resourceGroupNameParam,
    settings: {},
    type: "VMAccessForLinux",
    typeHandlerVersion: "1.4",
    vmExtensionName: `${virtualMachinesPulumirancherNameParam}/enablevmaccess`,
    vmName: virtualMachineResource.name,
});
const securityRuleResource = new azure_nextgen.network.v20200501.SecurityRule("securityRuleResource", {
    access: "Allow",
    description: "Docker",
    destinationAddressPrefix: "*",
    destinationAddressPrefixes: [],
    destinationPortRange: "2375",
    destinationPortRanges: [],
    direction: "Inbound",
    networkSecurityGroupName: networkSecurityGroupResource.name,
    priority: 300,
    protocol: "Tcp",
    resourceGroupName: resourceGroupNameParam,
    securityRuleName: `${networkSecurityGroupsPulumirancherNsgNameParam}/Docker`,
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
    networkSecurityGroupName: networkSecurityGroupResource.name,
    priority: 200,
    protocol: "Tcp",
    resourceGroupName: resourceGroupNameParam,
    securityRuleName: `${networkSecurityGroupsPulumirancherNsgNameParam}/Rancher`,
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
    networkSecurityGroupName: networkSecurityGroupResource.name,
    priority: 100,
    protocol: "*",
    resourceGroupName: resourceGroupNameParam,
    securityRuleName: `${networkSecurityGroupsPulumirancherNsgNameParam}/SSH`,
    sourceAddressPrefix: "*",
    sourceAddressPrefixes: [],
    sourcePortRange: "*",
    sourcePortRanges: [],
});
