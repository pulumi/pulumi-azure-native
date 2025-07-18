// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get virtual network.
 *
 * Uses Azure REST API version 2018-09-15.
 */
export function getVirtualNetwork(args: GetVirtualNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devtestlab:getVirtualNetwork", {
        "expand": args.expand,
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualNetworkArgs {
    /**
     * Specify the $expand query. Example: 'properties($expand=externalSubnets)'
     */
    expand?: string;
    /**
     * The name of the lab.
     */
    labName: string;
    /**
     * The name of the VirtualNetwork
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A virtual network.
 */
export interface GetVirtualNetworkResult {
    /**
     * The allowed subnets of the virtual network.
     */
    readonly allowedSubnets?: outputs.devtestlab.SubnetResponse[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The creation date of the virtual network.
     */
    readonly createdDate: string;
    /**
     * The description of the virtual network.
     */
    readonly description?: string;
    /**
     * The Microsoft.Network resource identifier of the virtual network.
     */
    readonly externalProviderResourceId?: string;
    /**
     * The external subnet properties.
     */
    readonly externalSubnets: outputs.devtestlab.ExternalSubnetResponse[];
    /**
     * The identifier of the resource.
     */
    readonly id: string;
    /**
     * The location of the resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState: string;
    /**
     * The subnet overrides of the virtual network.
     */
    readonly subnetOverrides?: outputs.devtestlab.SubnetOverrideResponse[];
    /**
     * The tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier: string;
}
/**
 * Get virtual network.
 *
 * Uses Azure REST API version 2018-09-15.
 */
export function getVirtualNetworkOutput(args: GetVirtualNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devtestlab:getVirtualNetwork", {
        "expand": args.expand,
        "labName": args.labName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetVirtualNetworkOutputArgs {
    /**
     * Specify the $expand query. Example: 'properties($expand=externalSubnets)'
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    labName: pulumi.Input<string>;
    /**
     * The name of the VirtualNetwork
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
