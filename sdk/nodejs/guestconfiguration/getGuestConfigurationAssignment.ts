// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get information about a guest configuration assignment
 *
 * Uses Azure REST API version 2024-04-05.
 *
 * Other available API versions: 2022-01-25. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native guestconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGuestConfigurationAssignment(args: GetGuestConfigurationAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetGuestConfigurationAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:guestconfiguration:getGuestConfigurationAssignment", {
        "guestConfigurationAssignmentName": args.guestConfigurationAssignmentName,
        "resourceGroupName": args.resourceGroupName,
        "vmName": args.vmName,
    }, opts);
}

export interface GetGuestConfigurationAssignmentArgs {
    /**
     * The guest configuration assignment name.
     */
    guestConfigurationAssignmentName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the virtual machine.
     */
    vmName: string;
}

/**
 * Guest configuration assignment is an association between a machine and guest configuration.
 */
export interface GetGuestConfigurationAssignmentResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * ARM resource id of the guest configuration assignment.
     */
    readonly id: string;
    /**
     * Region where the VM is located.
     */
    readonly location?: string;
    /**
     * The guest configuration assignment name.
     */
    readonly name: string;
    /**
     * Properties of the Guest configuration assignment.
     */
    readonly properties: outputs.guestconfiguration.GuestConfigurationAssignmentPropertiesResponse;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.guestconfiguration.SystemDataResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
/**
 * Get information about a guest configuration assignment
 *
 * Uses Azure REST API version 2024-04-05.
 *
 * Other available API versions: 2022-01-25. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native guestconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getGuestConfigurationAssignmentOutput(args: GetGuestConfigurationAssignmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGuestConfigurationAssignmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:guestconfiguration:getGuestConfigurationAssignment", {
        "guestConfigurationAssignmentName": args.guestConfigurationAssignmentName,
        "resourceGroupName": args.resourceGroupName,
        "vmName": args.vmName,
    }, opts);
}

export interface GetGuestConfigurationAssignmentOutputArgs {
    /**
     * The guest configuration assignment name.
     */
    guestConfigurationAssignmentName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the virtual machine.
     */
    vmName: pulumi.Input<string>;
}
