// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the specified load balancer inbound NAT rule.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getInboundNatRule(args: GetInboundNatRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetInboundNatRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network:getInboundNatRule", {
        "expand": args.expand,
        "inboundNatRuleName": args.inboundNatRuleName,
        "loadBalancerName": args.loadBalancerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetInboundNatRuleArgs {
    /**
     * Expands referenced resources.
     */
    expand?: string;
    /**
     * The name of the inbound NAT rule.
     */
    inboundNatRuleName: string;
    /**
     * The name of the load balancer.
     */
    loadBalancerName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Inbound NAT rule of the load balancer.
 */
export interface GetInboundNatRuleResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * A reference to backendAddressPool resource.
     */
    readonly backendAddressPool?: outputs.network.SubResourceResponse;
    /**
     * A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
     */
    readonly backendIPConfiguration: outputs.network.NetworkInterfaceIPConfigurationResponse;
    /**
     * The port used for the internal endpoint. Acceptable values range from 1 to 65535.
     */
    readonly backendPort?: number;
    /**
     * Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
     */
    readonly enableFloatingIP?: boolean;
    /**
     * Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
     */
    readonly enableTcpReset?: boolean;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * A reference to frontend IP addresses.
     */
    readonly frontendIPConfiguration?: outputs.network.SubResourceResponse;
    /**
     * The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
     */
    readonly frontendPort?: number;
    /**
     * The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
     */
    readonly frontendPortRangeEnd?: number;
    /**
     * The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
     */
    readonly frontendPortRangeStart?: number;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
     */
    readonly idleTimeoutInMinutes?: number;
    /**
     * The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * The reference to the transport protocol used by the load balancing rule.
     */
    readonly protocol?: string;
    /**
     * The provisioning state of the inbound NAT rule resource.
     */
    readonly provisioningState: string;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
/**
 * Gets the specified load balancer inbound NAT rule.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getInboundNatRuleOutput(args: GetInboundNatRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInboundNatRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network:getInboundNatRule", {
        "expand": args.expand,
        "inboundNatRuleName": args.inboundNatRuleName,
        "loadBalancerName": args.loadBalancerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetInboundNatRuleOutputArgs {
    /**
     * Expands referenced resources.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the inbound NAT rule.
     */
    inboundNatRuleName: pulumi.Input<string>;
    /**
     * The name of the load balancer.
     */
    loadBalancerName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
