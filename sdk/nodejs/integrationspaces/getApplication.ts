// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a Application
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getApplication(args: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:integrationspaces:getApplication", {
        "applicationName": args.applicationName,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetApplicationArgs {
    /**
     * The name of the Application
     */
    applicationName: string;
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
 * An integration application under space.
 */
export interface GetApplicationResult {
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
     * The tracking data stores.
     */
    readonly trackingDataStores?: {[key: string]: outputs.integrationspaces.TrackingDataStoreResponse};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get a Application
 *
 * Uses Azure REST API version 2023-11-14-preview.
 */
export function getApplicationOutput(args: GetApplicationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApplicationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:integrationspaces:getApplication", {
        "applicationName": args.applicationName,
        "resourceGroupName": args.resourceGroupName,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetApplicationOutputArgs {
    /**
     * The name of the Application
     */
    applicationName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the space
     */
    spaceName: pulumi.Input<string>;
}
