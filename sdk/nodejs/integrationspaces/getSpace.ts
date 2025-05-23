// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a Space
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getSpace(args: GetSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:integrationspaces:getSpace", {
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetSpaceArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the space
     */
    spaceName: string;
}

/**
 * An integration space.
 */
export interface GetSpaceResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The description of the resource.
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
     * The status of the last operation.
     */
    readonly provisioningState: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.integrationspaces.SystemDataResponse;
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
 * Get a Space
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getSpaceOutput(args: GetSpaceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSpaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:integrationspaces:getSpace", {
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetSpaceOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the space
     */
    spaceName: pulumi.Input<string>;
}
