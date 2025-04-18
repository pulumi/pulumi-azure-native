// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
 *
 * Uses Azure REST API version 2024-04-01.
 *
 * Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagedAzResiliencyStatus(args: GetManagedAzResiliencyStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedAzResiliencyStatusResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:servicefabric:getManagedAzResiliencyStatus", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedAzResiliencyStatusArgs {
    /**
     * The name of the cluster resource.
     */
    clusterName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * Describes the result of the request to list Managed VM Sizes for Service Fabric Managed Clusters.
 */
export interface GetManagedAzResiliencyStatusResult {
    /**
     * List of Managed VM Sizes for Service Fabric Managed Clusters.
     */
    readonly baseResourceStatus?: outputs.servicefabric.ResourceAzStatusResponse[];
    /**
     * URL to get the next set of Managed VM Sizes if there are any.
     */
    readonly isClusterZoneResilient: boolean;
}
/**
 * Action to get Az Resiliency Status of all the Base resources constituting Service Fabric Managed Clusters.
 *
 * Uses Azure REST API version 2024-04-01.
 *
 * Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getManagedAzResiliencyStatusOutput(args: GetManagedAzResiliencyStatusOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetManagedAzResiliencyStatusResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:servicefabric:getManagedAzResiliencyStatus", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedAzResiliencyStatusOutputArgs {
    /**
     * The name of the cluster resource.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
