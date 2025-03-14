// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a connection monitor by name.
 */
export function getConnectionMonitor(args: GetConnectionMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network/v20240301:getConnectionMonitor", {
        "connectionMonitorName": args.connectionMonitorName,
        "networkWatcherName": args.networkWatcherName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectionMonitorArgs {
    /**
     * The name of the connection monitor.
     */
    connectionMonitorName: string;
    /**
     * The name of the Network Watcher resource.
     */
    networkWatcherName: string;
    /**
     * The name of the resource group containing Network Watcher.
     */
    resourceGroupName: string;
}

/**
 * Information about the connection monitor.
 */
export interface GetConnectionMonitorResult {
    /**
     * Determines if the connection monitor will start automatically once created.
     */
    readonly autoStart?: boolean;
    /**
     * Type of connection monitor.
     */
    readonly connectionMonitorType: string;
    /**
     * Describes the destination of connection monitor.
     */
    readonly destination?: outputs.network.v20240301.ConnectionMonitorDestinationResponse;
    /**
     * List of connection monitor endpoints.
     */
    readonly endpoints?: outputs.network.v20240301.ConnectionMonitorEndpointResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * ID of the connection monitor.
     */
    readonly id: string;
    /**
     * Connection monitor location.
     */
    readonly location?: string;
    /**
     * Monitoring interval in seconds.
     */
    readonly monitoringIntervalInSeconds?: number;
    /**
     * The monitoring status of the connection monitor.
     */
    readonly monitoringStatus: string;
    /**
     * Name of the connection monitor.
     */
    readonly name: string;
    /**
     * Optional notes to be associated with the connection monitor.
     */
    readonly notes?: string;
    /**
     * List of connection monitor outputs.
     */
    readonly outputs?: outputs.network.v20240301.ConnectionMonitorOutputResponse[];
    /**
     * The provisioning state of the connection monitor.
     */
    readonly provisioningState: string;
    /**
     * Describes the source of connection monitor.
     */
    readonly source?: outputs.network.v20240301.ConnectionMonitorSourceResponse;
    /**
     * The date and time when the connection monitor was started.
     */
    readonly startTime: string;
    /**
     * Connection monitor tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * List of connection monitor test configurations.
     */
    readonly testConfigurations?: outputs.network.v20240301.ConnectionMonitorTestConfigurationResponse[];
    /**
     * List of connection monitor test groups.
     */
    readonly testGroups?: outputs.network.v20240301.ConnectionMonitorTestGroupResponse[];
    /**
     * Connection monitor type.
     */
    readonly type: string;
}
/**
 * Gets a connection monitor by name.
 */
export function getConnectionMonitorOutput(args: GetConnectionMonitorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConnectionMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network/v20240301:getConnectionMonitor", {
        "connectionMonitorName": args.connectionMonitorName,
        "networkWatcherName": args.networkWatcherName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConnectionMonitorOutputArgs {
    /**
     * The name of the connection monitor.
     */
    connectionMonitorName: pulumi.Input<string>;
    /**
     * The name of the Network Watcher resource.
     */
    networkWatcherName: pulumi.Input<string>;
    /**
     * The name of the resource group containing Network Watcher.
     */
    resourceGroupName: pulumi.Input<string>;
}
