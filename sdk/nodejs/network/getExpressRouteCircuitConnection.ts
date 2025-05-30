// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets the specified Express Route Circuit Connection from the specified express route circuit.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getExpressRouteCircuitConnection(args: GetExpressRouteCircuitConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetExpressRouteCircuitConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network:getExpressRouteCircuitConnection", {
        "circuitName": args.circuitName,
        "connectionName": args.connectionName,
        "peeringName": args.peeringName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetExpressRouteCircuitConnectionArgs {
    /**
     * The name of the express route circuit.
     */
    circuitName: string;
    /**
     * The name of the express route circuit connection.
     */
    connectionName: string;
    /**
     * The name of the peering.
     */
    peeringName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
 */
export interface GetExpressRouteCircuitConnectionResult {
    /**
     * /29 IP address space to carve out Customer addresses for tunnels.
     */
    readonly addressPrefix?: string;
    /**
     * The authorization key.
     */
    readonly authorizationKey?: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Express Route Circuit connection state.
     */
    readonly circuitConnectionStatus: string;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
     */
    readonly expressRouteCircuitPeering?: outputs.network.SubResourceResponse;
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * IPv6 Address PrefixProperties of the express route circuit connection.
     */
    readonly ipv6CircuitConnectionConfig?: outputs.network.Ipv6CircuitConnectionConfigResponse;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * Reference to Express Route Circuit Private Peering Resource of the peered circuit.
     */
    readonly peerExpressRouteCircuitPeering?: outputs.network.SubResourceResponse;
    /**
     * The provisioning state of the express route circuit connection resource.
     */
    readonly provisioningState: string;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
/**
 * Gets the specified Express Route Circuit Connection from the specified express route circuit.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getExpressRouteCircuitConnectionOutput(args: GetExpressRouteCircuitConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetExpressRouteCircuitConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network:getExpressRouteCircuitConnection", {
        "circuitName": args.circuitName,
        "connectionName": args.connectionName,
        "peeringName": args.peeringName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetExpressRouteCircuitConnectionOutputArgs {
    /**
     * The name of the express route circuit.
     */
    circuitName: pulumi.Input<string>;
    /**
     * The name of the express route circuit connection.
     */
    connectionName: pulumi.Input<string>;
    /**
     * The name of the peering.
     */
    peeringName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
