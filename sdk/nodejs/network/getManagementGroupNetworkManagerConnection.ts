// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a specified connection created by this management group.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagementGroupNetworkManagerConnection(args: GetManagementGroupNetworkManagerConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementGroupNetworkManagerConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:network:getManagementGroupNetworkManagerConnection", {
        "managementGroupId": args.managementGroupId,
        "networkManagerConnectionName": args.networkManagerConnectionName,
    }, opts);
}

export interface GetManagementGroupNetworkManagerConnectionArgs {
    /**
     * The management group Id which uniquely identify the Microsoft Azure management group.
     */
    managementGroupId: string;
    /**
     * Name for the network manager connection.
     */
    networkManagerConnectionName: string;
}

/**
 * The Network Manager Connection resource
 */
export interface GetManagementGroupNetworkManagerConnectionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * A description of the network manager connection.
     */
    readonly description?: string;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource ID.
     */
    readonly id: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Network Manager Id.
     */
    readonly networkManagerId?: string;
    /**
     * The system metadata related to this resource.
     */
    readonly systemData: outputs.network.SystemDataResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
/**
 * Get a specified connection created by this management group.
 *
 * Uses Azure REST API version 2024-05-01.
 *
 * Other available API versions: 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagementGroupNetworkManagerConnectionOutput(args: GetManagementGroupNetworkManagerConnectionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetManagementGroupNetworkManagerConnectionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:network:getManagementGroupNetworkManagerConnection", {
        "managementGroupId": args.managementGroupId,
        "networkManagerConnectionName": args.networkManagerConnectionName,
    }, opts);
}

export interface GetManagementGroupNetworkManagerConnectionOutputArgs {
    /**
     * The management group Id which uniquely identify the Microsoft Azure management group.
     */
    managementGroupId: pulumi.Input<string>;
    /**
     * Name for the network manager connection.
     */
    networkManagerConnectionName: pulumi.Input<string>;
}
