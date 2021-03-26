// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as azure_native from "@pulumi/azure-native";

const config = new pulumi.Config();
const managedClustersAzureNativeAksNameParam = config.get("managedClustersAzureNativeAksNameParam") || "azure-native-aks";
const publicIPAddresses87a245eeEdc1470cA08eC456659d9861ExternalidParam = config.get("publicIPAddresses87a245eeEdc1470cA08eC456659d9861ExternalidParam") || "/subscriptions/0282681f-7a9e-424b-80b2-96babd57a8a1/resourceGroups/MC_azure-native-go_azure-native-aks_westus/providers/Microsoft.Network/publicIPAddresses/87a245ee-edc1-470c-a08e-c456659d9861";
const resourceGroupNameParam = config.require("resourceGroupNameParam");
const managedClusterResource = new azure_native.containerservice.v20200701.ManagedCluster("managedClusterResource", {
    addonProfiles: {
        KubeDashboard: {
            enabled: true,
        },
    },
    agentPoolProfiles: [{
        count: 3,
        maxPods: 110,
        mode: "System",
        name: "agentpool",
        nodeLabels: {},
        orchestratorVersion: "1.16.10",
        osDiskSizeGB: 30,
        osType: "Linux",
        type: "VirtualMachineScaleSets",
        vmSize: "Standard_DS2_v2",
    }],
    dnsPrefix: "azurenativeprovider",
    enableRBAC: true,
    kubernetesVersion: "1.16.10",
    linuxProfile: {
        adminUsername: "testuser",
        ssh: {
            publicKeys: [{
                keyData: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC4oZK5s6zWKAG7Pif7qwih7DbheusimL6QYXHcrT2SN8H9ssuVykMuOKoz5iEynq2rbYMoBaxtWrVCpYhFa3n1/yxIxehyBDdiH2SLjFcSYwzL2AiBcGTP8OX/lxiqqOlCXqS1yjpaMK9rW1FBaVWQxiPx+LsRorV074cbJ/ChmEEfAGrez5mc+Ljz2CcPt5s2f1qFQ1Nc4pWRPHqq1UrUU2RT7e70FCo6fekOrLydH24pLwCwcMuf3TsrVmYNt/XglnpSLg9z3sClWjGsZnvT+7is2OqJNvSVxSGHozBGIeevTWPf+UunZllgB7TW6Jjz0cGABN7xqPiGInn33BP9DFP/bEzCP+TrJpt+yqa5wUJyKfJ19ky8zrG7G0QesqKirlEpzLIONGXeM4QQy5xOhcL9FtRL+SMRYWYvbqK8T5UIHsO49C11d5AtBkgLwMuUyG57D4j4l7dU+waU7X9LWx4HQNsNp5lTCJ5Ovh/FmufrgoYX4oKJp9ulOevmBefA58T9I+Hqcb6Li6n3rjCIHleQwf6VihSRL06Hk3YYnPruFCy2oq9mHTPdrkH7mpd++h3BV5rnSjeyaYx30y3+gzomXSeLtl7PeSSREmSX85AzzsDlCWwLqHxinCWsvnP4GIXu9moxZiiLQ3P88hcEV3LqLnCVUeT2fFQsuBnsyw==\n",
            }],
        },
    },
    location: "westus",
    networkProfile: {
        dnsServiceIP: "10.0.0.10",
        dockerBridgeCidr: "172.17.0.1/16",
        loadBalancerProfile: {
            effectiveOutboundIPs: [{
                id: publicIPAddresses87a245eeEdc1470cA08eC456659d9861ExternalidParam,
            }],
            managedOutboundIPs: {
                count: 1,
            },
        },
        loadBalancerSku: "Standard",
        networkPlugin: "kubenet",
        outboundType: "loadBalancer",
        podCidr: "10.244.0.0/16",
        serviceCidr: "10.0.0.0/16",
    },
    nodeResourceGroup: `MC_azure-native-go_${managedClustersAzureNativeAksNameParam}_westus`,
    resourceGroupName: resourceGroupNameParam,
    resourceName: managedClustersAzureNativeAksNameParam,
    servicePrincipalProfile: {
        clientId: "265097a5-0e6d-4315-8c76-998e743f5e2d",
    },
    sku: {
        name: "Basic",
        tier: "Free",
    },
});
const agentPoolResource = new azure_native.containerservice.v20200701.AgentPool("agentPoolResource", {
    agentPoolName: `${managedClustersAzureNativeAksNameParam}/agentpool`,
    count: 3,
    maxPods: 110,
    mode: "System",
    nodeLabels: {},
    orchestratorVersion: "1.16.10",
    osDiskSizeGB: 30,
    osType: "Linux",
    resourceGroupName: resourceGroupNameParam,
    resourceName: managedClusterResource.name,
    type: "VirtualMachineScaleSets",
    vmSize: "Standard_DS2_v2",
});
