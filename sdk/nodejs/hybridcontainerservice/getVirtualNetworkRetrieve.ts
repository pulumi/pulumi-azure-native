// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the Hybrid AKS virtual network
 *
 * Uses Azure REST API version 2022-09-01-preview.
 */
export function getVirtualNetworkRetrieve(args: GetVirtualNetworkRetrieveArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkRetrieveResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:hybridcontainerservice:getVirtualNetworkRetrieve", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworksName": args.virtualNetworksName,
    }, opts);
}

export interface GetVirtualNetworkRetrieveArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * Parameter for the name of the virtual network
     */
    virtualNetworksName: string;
}

/**
 * The virtualNetworks resource definition.
 */
export interface GetVirtualNetworkRetrieveResult {
    readonly extendedLocation?: outputs.hybridcontainerservice.VirtualNetworksResponseExtendedLocation;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
     */
    readonly properties: outputs.hybridcontainerservice.VirtualNetworksPropertiesResponse;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    readonly systemData: outputs.hybridcontainerservice.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets the Hybrid AKS virtual network
 *
 * Uses Azure REST API version 2022-09-01-preview.
 */
export function getVirtualNetworkRetrieveOutput(args: GetVirtualNetworkRetrieveOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualNetworkRetrieveResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:hybridcontainerservice:getVirtualNetworkRetrieve", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworksName": args.virtualNetworksName,
    }, opts);
}

export interface GetVirtualNetworkRetrieveOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Parameter for the name of the virtual network
     */
    virtualNetworksName: pulumi.Input<string>;
}
