// Copyright 2021, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

import * as resources from "@pulumi/azure-native/resources";
import * as extensions from "@pulumi/azure-native/extensions";

const resourceGroup = new resources.ResourceGroup("rg");

const template = {
    "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
    "contentVersion": "1.0.0.0",
    "parameters": {},
    "variables": {},
    "resources": [{
        "apiVersion": "2018-01-01-preview",
        "name": "eaggerer",
        "type": "Microsoft.ServiceBus/namespaces",
        "location": "westus",
        "properties": {},
      }],
    "outputs": {}
};

const arm = new extensions.ArmTemplate("template", {
    resourceGroupName: resourceGroup.name,
    content: JSON.stringify(template),
});

export const result = arm.result;
