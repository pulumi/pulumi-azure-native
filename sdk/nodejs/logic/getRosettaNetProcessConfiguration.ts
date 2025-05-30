// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets an integration account RosettaNetProcessConfiguration.
 *
 * Uses Azure REST API version 2016-06-01.
 */
export function getRosettaNetProcessConfiguration(args: GetRosettaNetProcessConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetRosettaNetProcessConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:logic:getRosettaNetProcessConfiguration", {
        "integrationAccountName": args.integrationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "rosettaNetProcessConfigurationName": args.rosettaNetProcessConfigurationName,
    }, opts);
}

export interface GetRosettaNetProcessConfigurationArgs {
    /**
     * The integration account name.
     */
    integrationAccountName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The integration account RosettaNetProcessConfiguration name.
     */
    rosettaNetProcessConfigurationName: string;
}

/**
 * The integration account RosettaNet process configuration.
 */
export interface GetRosettaNetProcessConfigurationResult {
    /**
     * The RosettaNet process configuration activity settings.
     */
    readonly activitySettings: outputs.logic.RosettaNetPipActivitySettingsResponse;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The changed time.
     */
    readonly changedTime: string;
    /**
     * The created time.
     */
    readonly createdTime: string;
    /**
     * The integration account RosettaNet ProcessConfiguration properties.
     */
    readonly description?: string;
    /**
     * The resource id.
     */
    readonly id: string;
    /**
     * The RosettaNet initiator role settings.
     */
    readonly initiatorRoleSettings: outputs.logic.RosettaNetPipRoleSettingsResponse;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The metadata.
     */
    readonly metadata?: {[key: string]: string};
    /**
     * Gets the resource name.
     */
    readonly name: string;
    /**
     * The integration account RosettaNet process code.
     */
    readonly processCode: string;
    /**
     * The integration account RosettaNet process name.
     */
    readonly processName: string;
    /**
     * The integration account RosettaNet process version.
     */
    readonly processVersion: string;
    /**
     * The RosettaNet responder role settings.
     */
    readonly responderRoleSettings: outputs.logic.RosettaNetPipRoleSettingsResponse;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets the resource type.
     */
    readonly type: string;
}
/**
 * Gets an integration account RosettaNetProcessConfiguration.
 *
 * Uses Azure REST API version 2016-06-01.
 */
export function getRosettaNetProcessConfigurationOutput(args: GetRosettaNetProcessConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRosettaNetProcessConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:logic:getRosettaNetProcessConfiguration", {
        "integrationAccountName": args.integrationAccountName,
        "resourceGroupName": args.resourceGroupName,
        "rosettaNetProcessConfigurationName": args.rosettaNetProcessConfigurationName,
    }, opts);
}

export interface GetRosettaNetProcessConfigurationOutputArgs {
    /**
     * The integration account name.
     */
    integrationAccountName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The integration account RosettaNetProcessConfiguration name.
     */
    rosettaNetProcessConfigurationName: pulumi.Input<string>;
}
