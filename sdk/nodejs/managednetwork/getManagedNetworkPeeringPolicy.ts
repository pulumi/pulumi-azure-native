// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The Get ManagedNetworkPeeringPolicies operation gets a Managed Network Peering Policy resource, specified by the  resource group, Managed Network name, and peering policy name
 *
 * Uses Azure REST API version 2019-06-01-preview.
 */
export function getManagedNetworkPeeringPolicy(args: GetManagedNetworkPeeringPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedNetworkPeeringPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:managednetwork:getManagedNetworkPeeringPolicy", {
        "managedNetworkName": args.managedNetworkName,
        "managedNetworkPeeringPolicyName": args.managedNetworkPeeringPolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedNetworkPeeringPolicyArgs {
    /**
     * The name of the Managed Network.
     */
    managedNetworkName: string;
    /**
     * The name of the Managed Network Peering Policy.
     */
    managedNetworkPeeringPolicyName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * The Managed Network Peering Policy resource
 */
export interface GetManagedNetworkPeeringPolicyResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Gets or sets the properties of a Managed Network Policy
     */
    readonly properties: outputs.managednetwork.ManagedNetworkPeeringPolicyPropertiesResponse;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
/**
 * The Get ManagedNetworkPeeringPolicies operation gets a Managed Network Peering Policy resource, specified by the  resource group, Managed Network name, and peering policy name
 *
 * Uses Azure REST API version 2019-06-01-preview.
 */
export function getManagedNetworkPeeringPolicyOutput(args: GetManagedNetworkPeeringPolicyOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetManagedNetworkPeeringPolicyResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:managednetwork:getManagedNetworkPeeringPolicy", {
        "managedNetworkName": args.managedNetworkName,
        "managedNetworkPeeringPolicyName": args.managedNetworkPeeringPolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedNetworkPeeringPolicyOutputArgs {
    /**
     * The name of the Managed Network.
     */
    managedNetworkName: pulumi.Input<string>;
    /**
     * The name of the Managed Network Peering Policy.
     */
    managedNetworkPeeringPolicyName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
