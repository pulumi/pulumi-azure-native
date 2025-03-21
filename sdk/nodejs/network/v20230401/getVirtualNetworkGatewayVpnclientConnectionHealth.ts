// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get VPN client connection health detail per P2S client connection of the virtual network gateway in the specified resource group.
 */
export function getVirtualNetworkGatewayVpnclientConnectionHealth(args: GetVirtualNetworkGatewayVpnclientConnectionHealthArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkGatewayVpnclientConnectionHealthResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20230401:getVirtualNetworkGatewayVpnclientConnectionHealth", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayVpnclientConnectionHealthArgs {
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
 * List of virtual network gateway vpn client connection health.
 */
export interface GetVirtualNetworkGatewayVpnclientConnectionHealthResult {
    /**
     * List of vpn client connection health.
     */
    readonly value?: outputs.network.v20230401.VpnClientConnectionHealthDetailResponse[];
}
/**
 * Get VPN client connection health detail per P2S client connection of the virtual network gateway in the specified resource group.
 */
export function getVirtualNetworkGatewayVpnclientConnectionHealthOutput(args: GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualNetworkGatewayVpnclientConnectionHealthResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20230401:getVirtualNetworkGatewayVpnclientConnectionHealth", {
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs {
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the virtual network gateway.
     */
    virtualNetworkGatewayName: pulumi.Input<string>;
}
