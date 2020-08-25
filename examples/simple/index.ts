import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";

const resourceGroup = new azurerm.resources.v20200601.ResourceGroup("rg", {
    resourceGroupName: "azurerm",
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

const staticSite = new azurerm.web.v20190801.StaticSite("staticsite", {
    resourceGroupName: resourceGroup.name,
    location: "westus2",
    repositoryUrl: "",
    branch: "master",
    repositoryToken: "",
    buildProperties: {
        appLocation: "/",
        apiLocation: "api",
        appArtifactLocation: ""
    },
    sku: {
        tier: "Free",
        name: "Free",
    },
    name: "mywebsite24234",
});

const containerinstance = new azurerm.containerinstance.v20191201.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
    // should be autonamed?
    containerGroupName: "abc-1234",
    location: "westus2",
    osType: "Linux",
    containers: [{
        name: "foo",
        image: "nginx",
        resources: {
            requests: {
                memoryInGB: 1,
                cpu: 1,
            },
        },
    }],
});

const vnet = new azurerm.network.v20200501.VirtualNetwork("vnet", {
    resourceGroupName: resourceGroup.name,
    virtualNetworkName: "vnet-1234",
    location: "westus2",
    addressSpace: {
        addressPrefixes: ["10.1.0.0/16"],
    },
    subnets: [{
        name: "default",
        addressPrefix: "10.1.0.0/24",
    }],
});

const subnet = new azurerm.network.v20200501.Subnet("subnet2", {
    resourceGroupName: resourceGroup.name,
    subnetName: "subnet2",
    virtualNetworkName: vnet.name,
    addressPrefix: "10.1.1.0/24",
});

const networkInterface = new azurerm.network.v20200501.NetworkInterface("nic", {
    resourceGroupName: resourceGroup.name,
    networkInterfaceName: "nic-1234",
    location: "westus2",
    ipConfigurations: [{
        name: "ipconfig1",
        subnet: {
            id: pulumi.interpolate`${vnet.id}/subnets/default`,
        },
        privateIPAllocationMethod: "Dynamic",
    }],
});

const virtualmachine  = new azurerm.compute.v20200601.VirtualMachine("vm", {
    resourceGroupName: resourceGroup.name,
    vmName: "abc-1234",
    location: "westus2",
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
});

const appServicePlan  = new azurerm.web.v20190801.AppServicePlan("app-plan", {
    resourceGroupName: resourceGroup.name,
    name: "app-plan",
    location: "westus2",
    kind: "Linux",
    sku: {
        name: "F1",
        capacity: 1
    },
});

const appService = new azurerm.web.v20190801.WebApp("app", {
    resourceGroupName: resourceGroup.name,
    name: "pulumiapp2418a",
    location: "westus2",
    serverFarmId: appServicePlan.id,
});

const storageAccount = new azurerm.storage.v20190601.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    accountName: "pulumi14345sa",
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
    enableHttpsTrafficOnly: true,
});

export const staticWebsiteUrl = pulumi.interpolate`https://${staticSite.defaultHostname}`;

export const existingRg = azurerm.resources.v20200601.getResourceGroup({ resourceGroupName: "Azure-Account-Cleanup" });

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(([resourceGroupName, accountName]) => 
    azurerm.storage.v20190601.listStorageAccountKeys({ resourceGroupName, accountName }));

export const primaryStorageKey = storageAccountKeys.keys[0].value;
