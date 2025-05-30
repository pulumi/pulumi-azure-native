// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a WorkloadNetworkPortMirroring
 *
 * Uses Azure REST API version 2023-09-01.
 *
 * Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWorkloadNetworkPortMirroring(args: GetWorkloadNetworkPortMirroringArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkloadNetworkPortMirroringResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:avs:getWorkloadNetworkPortMirroring", {
        "portMirroringId": args.portMirroringId,
        "privateCloudName": args.privateCloudName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWorkloadNetworkPortMirroringArgs {
    /**
     * ID of the NSX port mirroring profile.
     */
    portMirroringId: string;
    /**
     * Name of the private cloud
     */
    privateCloudName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * NSX Port Mirroring
 */
export interface GetWorkloadNetworkPortMirroringResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Destination VM Group.
     */
    readonly destination?: string;
    /**
     * Direction of port mirroring profile.
     */
    readonly direction?: string;
    /**
     * Display name of the port mirroring profile.
     */
    readonly displayName?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state
     */
    readonly provisioningState: string;
    /**
     * NSX revision number.
     */
    readonly revision?: number;
    /**
     * Source VM Group.
     */
    readonly source?: string;
    /**
     * Port Mirroring Status.
     */
    readonly status: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.avs.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a WorkloadNetworkPortMirroring
 *
 * Uses Azure REST API version 2023-09-01.
 *
 * Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getWorkloadNetworkPortMirroringOutput(args: GetWorkloadNetworkPortMirroringOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkloadNetworkPortMirroringResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:avs:getWorkloadNetworkPortMirroring", {
        "portMirroringId": args.portMirroringId,
        "privateCloudName": args.privateCloudName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetWorkloadNetworkPortMirroringOutputArgs {
    /**
     * ID of the NSX port mirroring profile.
     */
    portMirroringId: pulumi.Input<string>;
    /**
     * Name of the private cloud
     */
    privateCloudName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
