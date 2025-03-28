// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves the details of a nat rule.
 */
export function getVirtualNetworkGatewayNatRule(args: GetVirtualNetworkGatewayNatRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNetworkGatewayNatRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20240301:getVirtualNetworkGatewayNatRule", {
        "natRuleName": args.natRuleName,
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayNatRuleArgs {
    /**
     * The name of the nat rule.
     */
    natRuleName: string;
    /**
     * The resource group name of the Virtual Network Gateway.
     */
    resourceGroupName: string;
    /**
     * The name of the gateway.
     */
    virtualNetworkGatewayName: string;
}

/**
 * VirtualNetworkGatewayNatRule Resource.
 */
export interface GetVirtualNetworkGatewayNatRuleResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The private IP address external mapping for NAT.
     */
    readonly externalMappings?: outputs.network.v20240301.VpnNatRuleMappingResponse[];
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * The private IP address internal mapping for NAT.
     */
    readonly internalMappings?: outputs.network.v20240301.VpnNatRuleMappingResponse[];
    /**
     * The IP Configuration ID this NAT rule applies to.
     */
    readonly ipConfigurationId?: string;
    /**
     * The Source NAT direction of a VPN NAT.
     */
    readonly mode?: string;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The provisioning state of the NAT Rule resource.
     */
    readonly provisioningState: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Retrieves the details of a nat rule.
 */
export function getVirtualNetworkGatewayNatRuleOutput(args: GetVirtualNetworkGatewayNatRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualNetworkGatewayNatRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20240301:getVirtualNetworkGatewayNatRule", {
        "natRuleName": args.natRuleName,
        "resourceGroupName": args.resourceGroupName,
        "virtualNetworkGatewayName": args.virtualNetworkGatewayName,
    }, opts);
}

export interface GetVirtualNetworkGatewayNatRuleOutputArgs {
    /**
     * The name of the nat rule.
     */
    natRuleName: pulumi.Input<string>;
    /**
     * The resource group name of the Virtual Network Gateway.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the gateway.
     */
    virtualNetworkGatewayName: pulumi.Input<string>;
}
