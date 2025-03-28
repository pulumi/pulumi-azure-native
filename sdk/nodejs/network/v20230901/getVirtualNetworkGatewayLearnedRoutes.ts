// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * This operation retrieves a list of routes the virtual network gateway has learned, including routes learned from BGP peers.
 */
export function getVirtualNetworkGatewayLearnedRoutes(args: GetVirtualNetworkGatewayLearnedRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkGatewayLearnedRoutesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20230901:getVirtualNetworkGatewayLearnedRoutes", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayLearnedRoutesArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
    /**
     * The name of the virtual network gateway.
     */
    virtualNetworkGatewayName: string;
}

/**
 * List of virtual network gateway routes.
 */
export interface GetVirtualNetworkGatewayLearnedRoutesResult {
    /**
     * List of gateway routes.
     */
    readonly value?: outputs.network.v20230901.GatewayRouteResponse[];
}
/**
 * This operation retrieves a list of routes the virtual network gateway has learned, including routes learned from BGP peers.
 */
export function getVirtualNetworkGatewayLearnedRoutesOutput(args: GetVirtualNetworkGatewayLearnedRoutesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualNetworkGatewayLearnedRoutesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20230901:getVirtualNetworkGatewayLearnedRoutes", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayLearnedRoutesOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the virtual network gateway.
     */
    virtualNetworkGatewayName: pulumi.Input<string>;
}
