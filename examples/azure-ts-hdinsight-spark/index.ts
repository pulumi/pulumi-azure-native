// Copyright 2020, Pulumi Corporation.  All rights reserved.

import * as azurerm from "../../sdk/nodejs";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const location = config.require("location");
const username = config.require("username");
const password = config.require("password");

// Create an Azure Resource Group
const resourceGroup = new azurerm.resources.v20200601.ResourceGroup("spark-rg", {
    name: "azurerm", // ERR Required but not detected at compile time
    location: location,
    tags: {
        Owner: "azurerm-test",
    },
});

// Create a storage account and a container for Spark
const storageAccount = new azurerm.storage.v20190601.StorageAccount("sparksa", {
    resourceGroupName: resourceGroup.name,
    sku: {
        name: "Standard_LRS",
        tier: "Standard",
    },
    name: "sparksa12345", // ERR format validation
    location: location,
    kind: "StorageV2",
});

const storageContainer = new azurerm.storage.v20190601.BlobContainer("spark", {
    name: "spark-sc12345",
    accountName: storageAccount.name,
    resourceGroupName: resourceGroup.name,
});

const storageAccountKeys = pulumi.all([resourceGroup.name, storageAccount.name, storageAccount.id]).apply(
    ([resourceGroupName, accountName]) =>
        azurerm.storage.v20190601.listStorageAccountKeys({ resourceGroupName, accountName }));

const primaryStorageKey = storageAccountKeys.keys[0].value;

// Create a Spark cluster in HDInsight
const sparkCluster = new azurerm.hdinsight.v20180601beta1.Cluster("myspark", {
    name: "spark-cluster12345",
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
