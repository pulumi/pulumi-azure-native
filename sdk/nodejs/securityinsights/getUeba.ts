// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a setting.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 */
export function getUeba(args: GetUebaArgs, opts?: pulumi.InvokeOptions): Promise<GetUebaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:securityinsights:getUeba", {
        "resourceGroupName": args.resourceGroupName,
        "settingsName": args.settingsName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetUebaArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
     */
    settingsName: string;
    /**
     * The name of the workspace.
     */
    workspaceName: string;
}

/**
 * Settings with single toggle.
 */
export interface GetUebaResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The relevant data sources that enriched by ueba
     */
    readonly dataSources?: string[];
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
     */
    readonly id: string;
    /**
     * The kind of the setting
     * Expected value is 'Ueba'.
     */
    readonly kind: "Ueba";
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.securityinsights.SystemDataResponse;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets a setting.
 *
 * Uses Azure REST API version 2025-01-01-preview.
 */
export function getUebaOutput(args: GetUebaOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUebaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:securityinsights:getUeba", {
        "resourceGroupName": args.resourceGroupName,
        "settingsName": args.settingsName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetUebaOutputArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The setting name. Supports - Anomalies, EyesOn, EntityAnalytics, Ueba
     */
    settingsName: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
