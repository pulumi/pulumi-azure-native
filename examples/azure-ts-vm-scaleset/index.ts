// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

import * as azurerm from "../../sdk/nodejs";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

const config = new pulumi.Config();
const adminUser = config.get("adminUser") || "azureuser";
const adminPassword = config.getSecret("adminPassword") || new random.RandomPassword("pwd", {
    length: 20,
    special: true,
}).result;
const domain = config.get("domain") || new random.RandomString("domain", {
    length: 10,
    number: false,
    special: false,
    upper: false,
}).result;
const applicationPort = config.getNumber("applicationPort") || 80;

const location = config.get("location") || "westus2";

const resourceGroup = new azurerm.resources.v20200601.ResourceGroup("azurerm-rg", {
    name: "azurerg",
    location: location,
    tags: {
        Owner: "azurerm-test",
    },
});

const publicIp = new azurerm.network.v20200501.PublicIPAddress("public-ip", {
    name: "publicip",
    location: location, // Required by service
    resourceGroupName: resourceGroup.name,
    sku: {
        name: "Standard",
    },
    publicIPAllocationMethod: "Static",
    dnsSettings: {
        domainNameLabel: domain,
    },
});

const vnet = new azurerm.network.v20200501.VirtualNetwork("vnet", {
    resourceGroupName: resourceGroup.name,
    name: "vnet",
    location: location, // Required by service
    addressSpace: {
        addressPrefixes: ["10.0.0.0/16"],
    },
});

const subnet = new azurerm.network.v20200501.Subnet("subnet", {
    resourceGroupName: resourceGroup.name,
    name: "subnet",
    addressPrefix: "10.0.2.0/24",
    virtualNetworkName: vnet.name,
});

const loadBalancer = new azurerm.network.v20181201.LoadBalancer("lb", {
    resourceGroupName: resourceGroup.name,
    name: "lb",
    location: location, // Required by service
    sku: {
        name: "Standard"
    },
    frontendIPConfigurations: [{
        name: "PublicIPAddress",
        publicIPAddress: {
            id: publicIp.id,
        },
    }],
    //probes: [{
    //    properties: {
    //        protocol: "Tcp",
    //        port: applicationPort,
    //    }
    //}],
});

const inboundNatRule = new azurerm.network.v20200501.InboundNatRule("natrule", {
    loadBalancerName: loadBalancer.name,
    resourceGroupName: resourceGroup.name,
    name: "allow-80",
    frontendPort: applicationPort,
    backendPort: applicationPort,
    frontendIPConfiguration: loadBalancer.properties.frontendIPConfigurations[0]
})

const bpepool = new azurerm.network.v20200501.LoadBalancerBackendAddressPool("bpepool", {
    resourceGroupName: resourceGroup.name,
    name: "bpepool",
    loadBalancerName: loadBalancer.name,
}, { dependsOn: loadBalancer });

const customData = Buffer.from(`#cloud-config
packages:
    - nginx`).toString('base64')

const scaleSet = new azurerm.compute.v20200601.VirtualMachineScaleSet("vmscaleset", {
    resourceGroupName: resourceGroup.name,
    name: "vmscaleset",
    virtualMachineProfile: {
        networkProfile: {
            networkInterfaceConfigurations: [{
                name: "NICConfiguration",
                ipConfigurations: [{
                    name: "VMSetIPConfiguration",
                    loadBalancerBackendAddressPools: [{ id: bpepool.id }],
                    primary: true,
                    subnet: {
                        id: subnet.id,
                    },
                }],
                primary: true,
            }],
        },
        osProfile: {
            adminUsername: adminUser,
            adminPassword: adminPassword,
            computerNamePrefix: "vmlab",
            customData: customData,
            linuxConfiguration: {
                disablePasswordAuthentication: false,
            },
        },
        diagnosticsProfile: {
            bootDiagnostics: {
                enabled: true,
            },
        },
        storageProfile: {
            dataDisks: [{
                caching: "ReadWrite",
                createOption: "Empty",
                diskSizeGB: 10,
                lun: 0,
            }],
            imageReference: {
                offer: "UbuntuServer",
                publisher: "Canonical",
                sku: "16.04-LTS",
                version: "latest",
            },
            osDisk: {
                caching: "ReadWrite",
                createOption: "FromImage",
                managedDisk: { storageAccountType: "Standard_LRS" },
            },
        },
    },
    upgradePolicy: {
        mode: "Manual",
    },
    sku: {
        capacity: 1,
        name: "Standard_DS1_v2",
        tier: "Standard",
    },
    location: location
}, { dependsOn: [bpepool] });

const autoscale = new azurerm.insights.v20150401.AutoscaleSetting("vmss-autoscale", {
    resourceGroupName: resourceGroup.name,
    location: location,
    name: "vmss-autoscale",
    notifications: [{
        operation: "Scale",
        email: {
            customEmails: ["admin@contoso.com"],
            sendToSubscriptionAdministrator: true,
            sendToSubscriptionCoAdministrators: true,
        },
    }],
    profiles: [{
        capacity: {
            default: "1", // strings?
            maximum: "10",
            minimum: "1",
        },
        name: "defaultProfile",
        rules: [
            {
                metricTrigger: {
                    metricName: "Percentage CPU",
                    metricResourceUri: scaleSet.id,
                    operator: "GreaterThan",
                    statistic: "Average",
                    threshold: 75,
                    timeAggregation: "Average",
                    timeGrain: "PT1M",
                    timeWindow: "PT5M",
                },
                scaleAction: {
                    cooldown: "PT1M",
                    direction: "Increase",
                    type: "ChangeCount",
                    value: "1",
                },
            },
            {
                metricTrigger: {
                    metricName: "Percentage CPU",
                    metricResourceUri: scaleSet.id,
                    operator: "LessThan",
                    statistic: "Average",
                    threshold: 25,
                    timeAggregation: "Average",
                    timeGrain: "PT1M",
                    timeWindow: "PT5M",
                },
                scaleAction: {
                    cooldown: "PT1M",
                    direction: "Decrease",
                    type: "ChangeCount",
                    value: "1",
                },
            },
        ],
    }],
    targetResourceUri: scaleSet.id,
});

export const publicAddress = publicIp.properties.dnsSettings.fqdn;
