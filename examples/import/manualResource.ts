// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

import * as msRest from "@azure/ms-rest-js";
import * as msRestNodeAuth from "@azure/ms-rest-nodeauth";
import { ResourceManagementClient } from "@azure/arm-resources";
import { NetworkManagementClient } from "@azure/arm-network";

export interface Inputs {
    resourceGroupName: string;
    virtualNetworkName: string;
    virtualNetworkRange: string;
    location: string;
}

let subscriptionID: string | undefined;

interface Clients {
    resources: ResourceManagementClient;
    network: NetworkManagementClient;
}

async function getManagementClients(): Promise<Clients> {
    let credentials: msRest.ServiceClientCredentials;

    let clientID = process.env["ARM_CLIENT_ID"];
    let clientSecret = process.env["ARM_CLIENT_SECRET"];
    let tenantID = process.env["ARM_TENANT_ID"];
    subscriptionID = process.env["ARM_SUBSCRIPTION_ID"];

    // If they are still empty, try to get the creds from Az CLI.
    if (!clientID || !clientSecret || !tenantID || !subscriptionID) {
        const cliCredentials = await msRestNodeAuth.AzureCliCredentials.create();
        subscriptionID = cliCredentials.subscriptionInfo.id;
        credentials = cliCredentials;
    } else {
        credentials = await msRestNodeAuth.loginWithServicePrincipalSecret(clientID, clientSecret, tenantID);
    }

    return {
        resources: new ResourceManagementClient(credentials, subscriptionID),
        network: new NetworkManagementClient(credentials, subscriptionID),
    };
}

export async function createResources(inputs: Inputs): Promise<string> {
    const {resources, network} = await getManagementClients();

    await resources.resourceGroups.createOrUpdate(inputs.resourceGroupName, {
        location: inputs.location,
    });

    await network.virtualNetworks.createOrUpdate(inputs.resourceGroupName, inputs.virtualNetworkName, {
        location: inputs.location,
        addressSpace: {
            addressPrefixes: [`${inputs.virtualNetworkRange}/16`],
        },
        subnets: [{
            name: "default",
            addressPrefix: `${inputs.virtualNetworkRange}/24`,
        }],
    });

    return subscriptionID!;
}
