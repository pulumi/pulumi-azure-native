// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * A node pool snapshot resource.
 */
export function getSnapshot(args: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:containerservice/v20241002preview:getSnapshot", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetSnapshotArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the managed cluster resource.
     */
    resourceName: string;
}

/**
 * A node pool snapshot resource.
 */
export interface GetSnapshotResult {
    /**
     * CreationData to be used to specify the source agent pool resource ID to create this snapshot.
     */
    readonly creationData?: outputs.containerservice.v20241002preview.CreationDataResponse;
    /**
     * Whether to use a FIPS-enabled OS.
     */
    readonly enableFIPS: boolean;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The version of Kubernetes.
     */
    readonly kubernetesVersion: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The version of node image.
     */
    readonly nodeImageVersion: string;
    /**
     * Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.
     */
    readonly osSku: string;
    /**
     * The operating system type. The default is Linux.
     */
    readonly osType: string;
    /**
     * The type of a snapshot. The default is NodePool.
     */
    readonly snapshotType?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.containerservice.v20241002preview.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The size of the VM.
     */
    readonly vmSize: string;
}
/**
 * A node pool snapshot resource.
 */
export function getSnapshotOutput(args: GetSnapshotOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSnapshotResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:containerservice/v20241002preview:getSnapshot", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetSnapshotOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the managed cluster resource.
     */
    resourceName: pulumi.Input<string>;
}
