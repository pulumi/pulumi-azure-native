// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get the details of the serviceGroup
 *
 * Uses Azure REST API version 2024-02-01-preview.
 */
export function getServiceGroup(args: GetServiceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:management:getServiceGroup", {
        "serviceGroupName": args.serviceGroupName,
    }, opts);
}

export interface GetServiceGroupArgs {
    /**
     * ServiceGroup Name.
     */
    serviceGroupName: string;
}

/**
 * The serviceGroup details.
 */
export interface GetServiceGroupResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The kind of the serviceGroup.
     */
    readonly kind?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * ServiceGroup creation request body parameters.
     */
    readonly properties: outputs.management.ServiceGroupPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.management.SystemDataResponse;
    /**
     * The serviceGroup tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Get the details of the serviceGroup
 *
 * Uses Azure REST API version 2024-02-01-preview.
 */
export function getServiceGroupOutput(args: GetServiceGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetServiceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:management:getServiceGroup", {
        "serviceGroupName": args.serviceGroupName,
    }, opts);
}

export interface GetServiceGroupOutputArgs {
    /**
     * ServiceGroup Name.
     */
    serviceGroupName: pulumi.Input<string>;
}
