// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get bare metal machine key set of the provided cluster.
 *
 * Uses Azure REST API version 2025-02-01.
 *
 * Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBareMetalMachineKeySet(args: GetBareMetalMachineKeySetArgs, opts?: pulumi.InvokeOptions): Promise<GetBareMetalMachineKeySetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:networkcloud:getBareMetalMachineKeySet", {
        "bareMetalMachineKeySetName": args.bareMetalMachineKeySetName,
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBareMetalMachineKeySetArgs {
    /**
     * The name of the bare metal machine key set.
     */
    bareMetalMachineKeySetName: string;
    /**
     * The name of the cluster.
     */
    clusterName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

export interface GetBareMetalMachineKeySetResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
     */
    readonly azureGroupId: string;
    /**
     * The more detailed status of the key set.
     */
    readonly detailedStatus: string;
    /**
     * The descriptive message about the current detailed status.
     */
    readonly detailedStatusMessage: string;
    /**
     * Resource ETag.
     */
    readonly etag: string;
    /**
     * The date and time after which the users in this key set will be removed from the bare metal machines.
     */
    readonly expiration: string;
    /**
     * The extended location of the cluster associated with the resource.
     */
    readonly extendedLocation: outputs.networkcloud.ExtendedLocationResponse;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The list of IP addresses of jump hosts with management network access from which a login will be allowed for the users.
     */
    readonly jumpHostsAllowed: string[];
    /**
     * The last time this key set was validated.
     */
    readonly lastValidation: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The name of the group that users will be assigned to on the operating system of the machines.
     */
    readonly osGroupName?: string;
    /**
     * The access level allowed for the users in this key set.
     */
    readonly privilegeLevel: string;
    /**
     * The provisioning state of the bare metal machine key set.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.networkcloud.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The unique list of permitted users.
     */
    readonly userList: outputs.networkcloud.KeySetUserResponse[];
    /**
     * The status evaluation of each user.
     */
    readonly userListStatus: outputs.networkcloud.KeySetUserStatusResponse[];
}
/**
 * Get bare metal machine key set of the provided cluster.
 *
 * Uses Azure REST API version 2025-02-01.
 *
 * Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBareMetalMachineKeySetOutput(args: GetBareMetalMachineKeySetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBareMetalMachineKeySetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:networkcloud:getBareMetalMachineKeySet", {
        "bareMetalMachineKeySetName": args.bareMetalMachineKeySetName,
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetBareMetalMachineKeySetOutputArgs {
    /**
     * The name of the bare metal machine key set.
     */
    bareMetalMachineKeySetName: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
