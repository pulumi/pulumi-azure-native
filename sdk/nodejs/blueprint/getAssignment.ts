// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get a blueprint assignment.
 *
 * Uses Azure REST API version 2018-11-01-preview.
 */
export function getAssignment(args: GetAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:blueprint:getAssignment", {
        "assignmentName": args.assignmentName,
        "resourceScope": args.resourceScope,
    }, opts);
}

export interface GetAssignmentArgs {
    /**
     * Name of the blueprint assignment.
     */
    assignmentName: string;
    /**
     * The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
     */
    resourceScope: string;
}

/**
 * Represents a blueprint assignment.
 */
export interface GetAssignmentResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * ID of the published version of a blueprint definition.
     */
    readonly blueprintId?: string;
    /**
     * Multi-line explain this resource.
     */
    readonly description?: string;
    /**
     * One-liner string explain this resource.
     */
    readonly displayName?: string;
    /**
     * String Id used to locate any resource on Azure.
     */
    readonly id: string;
    /**
     * Managed identity for this blueprint assignment.
     */
    readonly identity: outputs.blueprint.ManagedServiceIdentityResponse;
    /**
     * The location of this blueprint assignment.
     */
    readonly location: string;
    /**
     * Defines how resources deployed by a blueprint assignment are locked.
     */
    readonly locks?: outputs.blueprint.AssignmentLockSettingsResponse;
    /**
     * Name of this resource.
     */
    readonly name: string;
    /**
     * Blueprint assignment parameter values.
     */
    readonly parameters: {[key: string]: outputs.blueprint.ParameterValueResponse};
    /**
     * State of the blueprint assignment.
     */
    readonly provisioningState: string;
    /**
     * Names and locations of resource group placeholders.
     */
    readonly resourceGroups: {[key: string]: outputs.blueprint.ResourceGroupValueResponse};
    /**
     * The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
     */
    readonly scope?: string;
    /**
     * Status of blueprint assignment. This field is readonly.
     */
    readonly status: outputs.blueprint.AssignmentStatusResponse;
    /**
     * Type of this resource.
     */
    readonly type: string;
}
/**
 * Get a blueprint assignment.
 *
 * Uses Azure REST API version 2018-11-01-preview.
 */
export function getAssignmentOutput(args: GetAssignmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:blueprint:getAssignment", {
        "assignmentName": args.assignmentName,
        "resourceScope": args.resourceScope,
    }, opts);
}

export interface GetAssignmentOutputArgs {
    /**
     * Name of the blueprint assignment.
     */
    assignmentName: pulumi.Input<string>;
    /**
     * The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
     */
    resourceScope: pulumi.Input<string>;
}
