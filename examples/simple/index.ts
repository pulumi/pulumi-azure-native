import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";

const resourceGroup = new azurerm.core.ResourceGroup("rg", {
    name: "azurerm",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

const staticSite = new azurerm.web.StaticSite("staticsite", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    properties: {
        repositoryUrl: "",
        branch: "master",
        repositoryToken: "",
        buildProperties: {
            appLocation: "/",
            apiLocation: "api",
            appArtifactLocation: ""
        }
    },
    sku: {
        tier: "Free",
        name: "Free",
    },
    name: "mywebsite24234",
});

export const staticWebsiteUrl = pulumi.interpolate`https://${staticSite.properties.defaultHostname}`;

const containerinstance = new azurerm.containerinstance.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
    // should be autonamed?
    name: "abc-1234",
    location: "westus2",
    // should be inlined via 'x-ms-client-flatten'
    properties: {
        osType: "Linux",
        containers: [{
            name: "foo",
            // should be inlined via 'x-ms-client-flatten'
            properties: {
                image: "nginx",
                resources: {
                    requests: {
                        memoryInGB: 1,
                        cpu: 1,
                    },
                },
            },
        }],
    },
});

const vnet = new azurerm.network.VirtualNetwork("vnet", {
    resourceGroupName: resourceGroup.name,
    name: "vnet-1234",
    location: "westus2",
    properties: {
        addressSpace: {
            addressPrefixes: ["10.1.0.0/16"],
        },
        subnets: [{
            name: "default",
            properties: {
                addressPrefix: "10.1.0.0/24",
            },
        }],
    },
});

const subnet = new azurerm.network.Subnet("subnet2", {
    resourceGroupName: resourceGroup.name,
    name: "subnet2",
    virtualNetworkName: vnet.name,
    properties: {
        addressPrefix: "10.1.1.0/24"
    },
});

const networkInterface = new azurerm.network.NetworkInterface("nic", {
    resourceGroupName: resourceGroup.name,
    name: "nic-1234",
    location: "westus2",
    properties: {
        ipConfigurations: [{
            name: "ipconfig1",
            properties: {
                subnet: {
                    id: pulumi.interpolate`${vnet.id}/subnets/default`,
                },
                privateIPAllocationMethod: "Dynamic",
            },
        }],
    },
});

const virtualmachine  = new azurerm.compute.VirtualMachine("vm", {
    resourceGroupName: resourceGroup.name,
    name: "abc-1234",
    location: "westus2",
    properties: {
        hardwareProfile: {
            vmSize: "Standard_A0",
        },
        networkProfile: {
            networkInterfaces: [{
                id: networkInterface.id,
            }],
        },
        storageProfile: {
            imageReference: {
                publisher: "Canonical",
                offer: "UbuntuServer",
                sku: "18.04-LTS",
                version: "latest"
            }
        },
        osProfile: {
            computerName: "foo",
            adminUsername: "someusername",
            adminPassword: "someFancyp@wd2!",
        },
    },
});

const appServicePlan  = new azurerm.web.AppServicePlan("app-plan", {
    resourceGroupName: resourceGroup.name,
    name: "app-plan",
    location: "westus2",
    kind: "Linux",
    sku: {
        name: "F1",
        capacity: 1
    },
});

const appService = new azurerm.web.WebApp("app", {
    resourceGroupName: resourceGroup.name,
    name: "pulumiapp2418a",
    location: "westus2",
    properties: {
        serverFarmId: appServicePlan.id,
    },
});

const storageAccount = new azurerm.storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    name: "pulumi14345sa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
    properties: {
        supportsHttpsTrafficOnly: true,
    }
});

export const existingRg = azurerm.core.getResourceGroup({ name: "Azure-Account-Cleanup" });

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(([resourceGroupName, accountName]) => 
    azurerm.storage.listStorageAccountKeys({ resourceGroupName, accountName }));

export const primaryStorageKey = storageAccountKeys.keys[0].value;
