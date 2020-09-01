import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "@pulumi/azurerm";
import * as random from "@pulumi/random";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
})

const resourceGroup = new azurerm.resources.latest.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
    tags: {
        Owner: "mikhailshilkov",
        Env: "prod2",
    },
});

const staticSite = new azurerm.web.latest.StaticSite("staticsite", {
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
    name: randomString.result,
});

const containerinstance = new azurerm.containerinstance.latest.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
    // should be autonamed?
    containerGroupName: randomString.result,
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

const vnet = new azurerm.network.latest.VirtualNetwork("vnet", {
    resourceGroupName: resourceGroup.name,
    virtualNetworkName: randomString.result,
    location: "westus2",
    addressSpace: {
        addressPrefixes: ["10.1.0.0/16"],
    },
    subnets: [{
        name: "default",
        addressPrefix: "10.1.0.0/24",
    }],
});

const subnet = new azurerm.network.latest.Subnet("subnet2", {
    resourceGroupName: resourceGroup.name,
    subnetName: randomString.result,
    virtualNetworkName: vnet.name,
    addressPrefix: "10.1.1.0/24",
});

const networkInterface = new azurerm.network.latest.NetworkInterface("nic", {
    resourceGroupName: resourceGroup.name,
    networkInterfaceName: randomString.result,
    location: "westus2",
    ipConfigurations: [{
        name: "ipconfig1",
        subnet: {
            id: pulumi.interpolate`${vnet.id}/subnets/default`,
        },
        privateIPAllocationMethod: "Dynamic",
    }],
});

const virtualmachine  = new azurerm.compute.latest.VirtualMachine("vm", {
    resourceGroupName: resourceGroup.name,
    vmName: randomString.result,
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
        computerName: randomString.result,
        adminUsername: "someusername",
        adminPassword: "someFancyp@wd2!",
    },
});

const appServicePlan  = new azurerm.web.latest.AppServicePlan("app-plan", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    location: "westus2",
    kind: "Linux",
    sku: {
        name: "F1",
        capacity: 1
    },
});

const appService = new azurerm.web.latest.WebApp("app", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    location: "westus2",
    serverFarmId: appServicePlan.id,
});

const storageAccount = new azurerm.storage.latest.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    accountName: randomString.result,
    location: "westus2",
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    kind: "StorageV2",
    enableHttpsTrafficOnly: true,
});

export const staticWebsiteUrl = pulumi.interpolate`https://${staticSite.defaultHostname}`;

export const existingRg = azurerm.resources.latest.getResourceGroup({ resourceGroupName: "Azure-Account-Cleanup" });

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(([resourceGroupName, accountName]) =>
    azurerm.storage.latest.listStorageAccountKeys({ resourceGroupName, accountName }));

export const primaryStorageKey = storageAccountKeys.keys[0].value;
