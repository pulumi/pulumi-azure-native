// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a machine pool
 *
 * Uses Azure REST API version 2024-02-01.
 *
 * Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPool(args: GetPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:devcenter:getPool", {
        "poolName": args.poolName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPoolArgs {
    /**
     * Name of the pool.
     */
    poolName: string;
    /**
     * The name of the project.
     */
    projectName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * A pool of Virtual Machines.
 */
export interface GetPoolResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Indicates the number of provisioned Dev Boxes in this pool.
     */
    readonly devBoxCount: number;
    /**
     * Name of a Dev Box definition in parent Project of this Pool
     */
    readonly devBoxDefinitionName: string;
    /**
     * The display name of the pool.
     */
    readonly displayName?: string;
    /**
     * Overall health status of the Pool. Indicates whether or not the Pool is available to create Dev Boxes.
     */
    readonly healthStatus: string;
    /**
     * Details on the Pool health status to help diagnose issues. This is only populated when the pool status indicates the pool is in a non-healthy state
     */
    readonly healthStatusDetails: outputs.devcenter.HealthStatusDetailResponse[];
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
     */
    readonly licenseType: string;
    /**
     * Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
     */
    readonly localAdministrator: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The regions of the managed virtual network (required when managedNetworkType is Managed).
     */
    readonly managedVirtualNetworkRegions?: string[];
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Name of a Network Connection in parent Project of this Pool
     */
    readonly networkConnectionName: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
     */
    readonly singleSignOnStatus?: string;
    /**
     * Stop on disconnect configuration settings for Dev Boxes created in this pool.
     */
    readonly stopOnDisconnect?: outputs.devcenter.StopOnDisconnectConfigurationResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.devcenter.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
     */
    readonly virtualNetworkType?: string;
}
/**
 * Gets a machine pool
 *
 * Uses Azure REST API version 2024-02-01.
 *
 * Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPoolOutput(args: GetPoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:devcenter:getPool", {
        "poolName": args.poolName,
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPoolOutputArgs {
    /**
     * Name of the pool.
     */
    poolName: pulumi.Input<string>;
    /**
     * The name of the project.
     */
    projectName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
