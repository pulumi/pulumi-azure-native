import * as pulumi from "@pulumi/pulumi";

import * as compute from "@pulumi/azure-native/compute";
import * as eventgrid from "@pulumi/azure-native/eventgrid";
import * as network from "@pulumi/azure-native/network";
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pulumi/azure-native/storage";
import * as web from "@pulumi/azure-native/web";

const resourceGroup = new resources.ResourceGroup("rg", {
    location: "eastus2", // explicit location is here to test the location propagation from resource group to other resources
});

const staticSite = new web.StaticSite("staticsite", {
    resourceGroupName: resourceGroup.name,
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
    location: "westeurope", // it's still possible to set an explict location for a resource
});

const vnet = new network.VirtualNetwork("vnet", {
    resourceGroupName: resourceGroup.name,
    addressSpace: {
        addressPrefixes: ["10.1.0.0/16"],
    },
    subnets: [{
        name: "default",
        addressPrefix: "10.1.0.0/24",
    }],
});

const subnet = new network.Subnet("subnet2", {
    resourceGroupName: resourceGroup.name,
    virtualNetworkName: vnet.name,
    addressPrefix: "10.1.1.0/24",
    serviceEndpoints: [{ service: "Microsoft.Storage" }],
    delegations: [{
        name: "webapp",
        serviceName: "Microsoft.Web/serverFarms",
    }],
});

const networkInterface = new network.NetworkInterface("nic", {
    resourceGroupName: resourceGroup.name,
    ipConfigurations: [{
        name: "ipconfig1",
        subnet: {
            id: pulumi.interpolate`${vnet.id}/subnets/default`,
        },
        privateIPAllocationMethod: network.IPAllocationMethod.Dynamic,
    }],
});

const publicIP = new network.PublicIPAddress("pip", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: network.PublicIPAddressSkuName.Basic,
    },
    publicIPAddressVersion: network.IPVersion.IPv4,
    publicIPAllocationMethod: network.IPAllocationMethod.Dynamic,
});

const virtualmachine  = new compute.VirtualMachine("vm", {
    resourceGroupName: resourceGroup.name,
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
        computerName: "mycomputer",
        adminUsername: "someusername",
        adminPassword: "someFancyp@wd2!",
    },
});

const appServicePlan  = new web.AppServicePlan("app-plan", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: "S1",
        capacity: 1
    },
});

const appService = new web.WebApp("app", {
    resourceGroupName: resourceGroup.name,
    serverFarmId: appServicePlan.id,
    kind: "app",
    siteConfig: {
        appSettings: [{
            name: "test",
            value: "this is a slot setting",
        }],
    },
});

new web.WebAppSlotConfigurationNames("names", {
    resourceGroupName: resourceGroup.name,
    name: appService.name,
    appSettingNames: ["test"],
});

new web.WebAppSwiftVirtualNetworkConnection("swiftconn", {
    subnetResourceId: subnet.id,
    name: appService.name,
    resourceGroupName: resourceGroup.name,
    swiftSupported: true,
});

const storageAccount = new storage.StorageAccount("sa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: storage.SkuName.Standard_LRS,
    },
    kind: storage.Kind.StorageV2,
    enableHttpsTrafficOnly: true,
});

var queue = new storage.Queue("queue", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
});

const eventGridSub = new eventgrid.EventSubscription("egsub", {
    scope: resourceGroup.id,
    destination: {        
        endpointType: "StorageQueue",
        queueName: queue.name,
        resourceId: storageAccount.id,
    },
    filter: {
        isSubjectCaseSensitive: true,
    }
});

new storage.BlobServiceProperties("blobprops", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    blobServicesName: "default",
    isVersioningEnabled: true,
    cors: {
        corsRules: [{
            allowedHeaders: ["*"],
            allowedMethods: ["GET", "PUT"],
            allowedOrigins: ["*"],
            exposedHeaders: ["*"],
            maxAgeInSeconds: 1000,
        }],
    },
});

export const staticWebsiteUrl = pulumi.interpolate`https://${staticSite.defaultHostname}`;

export const existingRg = resources.getResourceGroup({ resourceGroupName: "Azure-Account-Cleanup" });

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(([resourceGroupName, accountName]) =>
    storage.listStorageAccountKeys({ resourceGroupName, accountName }));

export const primaryStorageKey = storageAccountKeys.keys[0].value;
