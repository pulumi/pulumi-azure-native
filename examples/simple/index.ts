import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

import * as compute from "@pulumi/azure-nextgen/compute/latest";
// TODO change to latest when https://github.com/Azure/azure-rest-api-specs/issues/11634 is fixed
import * as containerinstance from "@pulumi/azure-nextgen/containerinstance/v20191201";
import * as eventgrid from "@pulumi/azure-nextgen/eventgrid/latest";
import * as managedidentity from "@pulumi/azure-nextgen/managedidentity/latest";
import * as network from "@pulumi/azure-nextgen/network/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as storage from "@pulumi/azure-nextgen/storage/latest";
// TODO change to latest when https://github.com/Azure/azure-rest-api-specs/issues/12005 is fixed
import * as web from "@pulumi/azure-nextgen/web/v20200601";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString.result,
    location: "westus2",
});

const staticSite = new web.StaticSite("staticsite", {
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

const userIdentity = new managedidentity.UserAssignedIdentity("uai", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    resourceName: randomString.result,
});

const container = new containerinstance.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
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
    identity: {
        type: "UserAssigned",
        userAssignedIdentities: userIdentity.id.apply(id => {
            const dict: { [key: string] : object } = {};
            dict[id] = {};
            return dict;
        }),
    },
});

const vnet = new network.VirtualNetwork("vnet", {
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

const subnet = new network.Subnet("subnet2", {
    resourceGroupName: resourceGroup.name,
    subnetName: randomString.result,
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

const publicIP = new network.PublicIPAddress("pip", {
    resourceGroupName: resourceGroup.name,
    publicIpAddressName: randomString.result,
    location: "westus2",
    sku: {
        name: "Basic",  
    },
    publicIPAddressVersion: "IPv4",
    publicIPAllocationMethod: "Dynamic",
});

const virtualmachine  = new compute.VirtualMachine("vm", {
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

const appServicePlan  = new web.AppServicePlan("app-plan", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    location: "westus2",
    sku: {
        name: "S1",
        capacity: 1
    },
});

const appService = new web.WebApp("app", {
    resourceGroupName: resourceGroup.name,
    name: randomString.result,
    location: "westus2",
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
    accountName: randomString.result,
    location: "westus2",
    sku: {
        name: "Standard_LRS",
    },
    kind: "StorageV2",
    enableHttpsTrafficOnly: true,
});

var queue = new storage.Queue("queue", {
    resourceGroupName: resourceGroup.name,
    accountName: storageAccount.name,
    queueName: "event-grid-dest",
});

const eventGridSub = new eventgrid.EventSubscription("egsub", {
    eventSubscriptionName: randomString.result,
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
