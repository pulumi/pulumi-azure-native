// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get configuration assignment for resource..
 *
 * Uses Azure REST API version 2023-10-01-preview.
 *
 * Other available API versions: 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getConfigurationAssignmentsForResourceGroup(args: GetConfigurationAssignmentsForResourceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationAssignmentsForResourceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:maintenance:getConfigurationAssignmentsForResourceGroup", {
        "configurationAssignmentName": args.configurationAssignmentName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConfigurationAssignmentsForResourceGroupArgs {
    /**
     * Configuration assignment name
     */
    configurationAssignmentName: string;
    /**
     * Resource group name
     */
    resourceGroupName: string;
}

/**
 * Configuration Assignment
 */
export interface GetConfigurationAssignmentsForResourceGroupResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Properties of the configuration assignment
     */
    readonly filter?: outputs.maintenance.ConfigurationAssignmentFilterPropertiesResponse;
    /**
     * Fully qualified identifier of the resource
     */
    readonly id: string;
    /**
     * Location of the resource
     */
    readonly location?: string;
    /**
     * The maintenance configuration Id
     */
    readonly maintenanceConfigurationId?: string;
    /**
     * Name of the resource
     */
    readonly name: string;
    /**
     * The unique resourceId
     */
    readonly resourceId?: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.maintenance.SystemDataResponse;
    /**
     * Type of the resource
     */
    readonly type: string;
}
/**
 * Get configuration assignment for resource..
 *
 * Uses Azure REST API version 2023-10-01-preview.
 *
 * Other available API versions: 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getConfigurationAssignmentsForResourceGroupOutput(args: GetConfigurationAssignmentsForResourceGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConfigurationAssignmentsForResourceGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:maintenance:getConfigurationAssignmentsForResourceGroup", {
        "configurationAssignmentName": args.configurationAssignmentName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConfigurationAssignmentsForResourceGroupOutputArgs {
    /**
     * Configuration assignment name
     */
    configurationAssignmentName: pulumi.Input<string>;
    /**
     * Resource group name
     */
    resourceGroupName: pulumi.Input<string>;
}
