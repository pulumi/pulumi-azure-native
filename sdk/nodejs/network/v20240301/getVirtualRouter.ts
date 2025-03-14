// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the specified Virtual Router.
 */
export function getVirtualRouter(args: GetVirtualRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualRouterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20240301:getVirtualRouter", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "virtualRouterName": args.virtualRouterName,
    }, opts);
}

export interface GetVirtualRouterArgs {
    /**
     * Expands referenced resources.
     */
    expand?: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The name of the Virtual Router.
     */
    virtualRouterName: string;
}

/**
 * VirtualRouter Resource.
 */
export interface GetVirtualRouterResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The Gateway on which VirtualRouter is hosted.
     */
    readonly hostedGateway?: outputs.network.v20240301.SubResourceResponse;
    /**
     * The Subnet on which VirtualRouter is hosted.
     */
    readonly hostedSubnet?: outputs.network.v20240301.SubResourceResponse;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * List of references to VirtualRouterPeerings.
     */
    readonly peerings: outputs.network.v20240301.SubResourceResponse[];
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * VirtualRouter ASN.
     */
    readonly virtualRouterAsn?: number;
    /**
     * VirtualRouter IPs.
     */
    readonly virtualRouterIps?: string[];
}
/**
 * Gets the specified Virtual Router.
 */
export function getVirtualRouterOutput(args: GetVirtualRouterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualRouterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20240301:getVirtualRouter", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "virtualRouterName": args.virtualRouterName,
    }, opts);
}

export interface GetVirtualRouterOutputArgs {
    /**
     * Expands referenced resources.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Virtual Router.
     */
    virtualRouterName: pulumi.Input<string>;
}
