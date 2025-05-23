// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a Big Data pool.
 *
 * Uses Azure REST API version 2021-06-01.
 *
 * Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBigDataPool(args: GetBigDataPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetBigDataPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:synapse:getBigDataPool", {
        "bigDataPoolName": args.bigDataPoolName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetBigDataPoolArgs {
    /**
     * Big Data pool name
     */
    bigDataPoolName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * A Big Data pool
 */
export interface GetBigDataPoolResult {
    /**
     * Auto-pausing properties
     */
    readonly autoPause?: outputs.synapse.AutoPausePropertiesResponse;
    /**
     * Auto-scaling properties
     */
    readonly autoScale?: outputs.synapse.AutoScalePropertiesResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The cache size
     */
    readonly cacheSize?: number;
    /**
     * The time when the Big Data pool was created.
     */
    readonly creationDate: string;
    /**
     * List of custom libraries/packages associated with the spark pool.
     */
    readonly customLibraries?: outputs.synapse.LibraryInfoResponse[];
    /**
     * The default folder where Spark logs will be written.
     */
    readonly defaultSparkLogFolder?: string;
    /**
     * Dynamic Executor Allocation
     */
    readonly dynamicExecutorAllocation?: outputs.synapse.DynamicExecutorAllocationResponse;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Whether autotune is required or not.
     */
    readonly isAutotuneEnabled?: boolean;
    /**
     * Whether compute isolation is required or not.
     */
    readonly isComputeIsolationEnabled?: boolean;
    /**
     * The time when the Big Data pool was updated successfully.
     */
    readonly lastSucceededTimestamp: string;
    /**
     * Library version requirements
     */
    readonly libraryRequirements?: outputs.synapse.LibraryRequirementsResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The number of nodes in the Big Data pool.
     */
    readonly nodeCount?: number;
    /**
     * The level of compute power that each node in the Big Data pool has.
     */
    readonly nodeSize?: string;
    /**
     * The kind of nodes that the Big Data pool provides.
     */
    readonly nodeSizeFamily?: string;
    /**
     * The state of the Big Data pool.
     */
    readonly provisioningState?: string;
    /**
     * Whether session level packages enabled.
     */
    readonly sessionLevelPackagesEnabled?: boolean;
    /**
     * Spark configuration file to specify additional properties
     */
    readonly sparkConfigProperties?: outputs.synapse.SparkConfigPropertiesResponse;
    /**
     * The Spark events folder
     */
    readonly sparkEventsFolder?: string;
    /**
     * The Apache Spark version.
     */
    readonly sparkVersion?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a Big Data pool.
 *
 * Uses Azure REST API version 2021-06-01.
 *
 * Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getBigDataPoolOutput(args: GetBigDataPoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBigDataPoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:synapse:getBigDataPool", {
        "bigDataPoolName": args.bigDataPoolName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetBigDataPoolOutputArgs {
    /**
     * Big Data pool name
     */
    bigDataPoolName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
