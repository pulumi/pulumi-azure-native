// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieve the Dsc node configurations by node configuration.
 */
export function getDscNodeConfiguration(args: GetDscNodeConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetDscNodeConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:automation/v20220808:getDscNodeConfiguration", {
        "automationAccountName": args.automationAccountName,
        "nodeConfigurationName": args.nodeConfigurationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDscNodeConfigurationArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: string;
    /**
     * The Dsc node configuration name.
     */
    nodeConfigurationName: string;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: string;
}

/**
 * Definition of the dsc node configuration.
 */
export interface GetDscNodeConfigurationResult {
    /**
     * Gets or sets the configuration of the node.
     */
    readonly configuration?: outputs.automation.v20220808.DscConfigurationAssociationPropertyResponse;
    /**
     * Gets or sets creation time.
     */
    readonly creationTime?: string;
    /**
     * Fully qualified resource Id for the resource
     */
    readonly id: string;
    /**
     * If a new build version of NodeConfiguration is required.
     */
    readonly incrementNodeConfigurationBuild?: boolean;
    /**
     * Gets or sets the last modified time.
     */
    readonly lastModifiedTime?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Number of nodes with this node configuration assigned
     */
    readonly nodeCount?: number;
    /**
     * Source of node configuration.
     */
    readonly source?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
/**
 * Retrieve the Dsc node configurations by node configuration.
 */
export function getDscNodeConfigurationOutput(args: GetDscNodeConfigurationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDscNodeConfigurationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:automation/v20220808:getDscNodeConfiguration", {
        "automationAccountName": args.automationAccountName,
        "nodeConfigurationName": args.nodeConfigurationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDscNodeConfigurationOutputArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * The Dsc node configuration name.
     */
    nodeConfigurationName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
