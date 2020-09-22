// Copyright 2020, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

import * as hdinsight from "@pulumi/azure-nextgen/hdinsight/v20180601preview";
import * as resources from "@pulumi/azure-nextgen/resources/latest";
import * as storage from "@pulumi/azure-nextgen/storage/latest";

const config = new pulumi.Config();
const location = config.require("location");
const username = config.require("username");
const password = config.require("password");

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
});

// Create an Azure Resource Group
const resourceGroup = new resources.ResourceGroup("spark-rg", {
    resourceGroupName: randomString.result,
    location: location,
});

// Create a storage account and a container for Spark
const storageAccount = new storage.StorageAccount("sparksa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: "Standard_LRS"
    },
    accountName: randomString.result,
    location: location,
    kind: "StorageV2",
});

const storageContainer = new storage.BlobContainer("spark", {
    containerName: randomString.result,
    accountName: storageAccount.name,
    resourceGroupName: resourceGroup.name,
});

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(
    ([resourceGroupName, accountName]) =>
        storage.listStorageAccountKeys({ resourceGroupName, accountName }));

const primaryStorageKey = storageAccountKeys.keys[0].value;

// Create a Spark cluster in HDInsight
const sparkCluster = new hdinsight.Cluster("myspark", {
    clusterName: randomString.result,
    resourceGroupName: resourceGroup.name,
    location: location,
    properties: {
        clusterVersion: "3.6",
        osType: "Linux",
        clusterDefinition: {
            kind: "spark",
            componentVersion: {
                "spark": "2.3",
            },
            configurations: {
                "gateway": {
                    "restAuthCredential.isEnabled": true,
                    "restAuthCredential.username": username,
                    "restAuthCredential.password": password,
                },
            },
        },
        tier: "Standard",
        storageProfile: {
            storageaccounts: [{
                isDefault: true,
                key: primaryStorageKey,
                name: pulumi.interpolate`${storageAccount.name}.blob.core.windows.net`,
                container: storageContainer.name,
            }]
        },
        computeProfile: {
            roles: [
                {
                    name: "headnode",
                    targetInstanceCount: 1,
                    hardwareProfile: {
                        vmSize: "Standard_D12_v2",
                    },
                    osProfile: {
                        linuxOperatingSystemProfile: {
                            username: username,
                            password: password,
                        },
                    },
                },
                {
                    name: "workernode",
                    targetInstanceCount: 1,
                    hardwareProfile: {
                        vmSize: "Standard_D12_v2",
                    },
                    osProfile: {
                        linuxOperatingSystemProfile: {
                            username: username,
                            password: password,
                        },
                    },
                },
                {
                    name: "zookeepernode",
                    targetInstanceCount: 3,
                    hardwareProfile: {
                        vmSize: "Standard_D12_v2",
                    },
                    osProfile: {
                        linuxOperatingSystemProfile: {
                            username: username,
                            password: password,
                        },
                    },
                },
            ],
        },
    },
});

const httpsEndpoint = sparkCluster.properties.connectivityEndpoints?.apply(
    endpoints => endpoints?.filter(e => e.name?.toLowerCase() == "https").
        map(x => x.location)[0])

// Export the endpoint
export const endpoint = pulumi.interpolate`https://${httpsEndpoint}/`;
