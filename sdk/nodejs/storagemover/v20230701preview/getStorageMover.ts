// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a Storage Mover resource.
 */
export function getStorageMover(args: GetStorageMoverArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageMoverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:storagemover/v20230701preview:getStorageMover", {
        "resourceGroupName": args.resourceGroupName,
        "storageMoverName": args.storageMoverName,
    }, opts);
}

export interface GetStorageMoverArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the Storage Mover resource.
     */
    storageMoverName: string;
}

/**
 * The Storage Mover resource, which is a container for a group of Agents, Projects, and Endpoints.
 */
export interface GetStorageMoverResult {
    /**
     * A description for the Storage Mover.
     */
    readonly description?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The provisioning state of this resource.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.storagemover.v20230701preview.SystemDataResponse;
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
 * Gets a Storage Mover resource.
 */
export function getStorageMoverOutput(args: GetStorageMoverOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetStorageMoverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:storagemover/v20230701preview:getStorageMover", {
        "resourceGroupName": args.resourceGroupName,
        "storageMoverName": args.storageMoverName,
    }, opts);
}

export interface GetStorageMoverOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Storage Mover resource.
     */
    storageMoverName: pulumi.Input<string>;
}
