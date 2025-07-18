// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The request to update subscriptions needed to be monitored by the Elastic monitor resource.
 *
 * Uses Azure REST API version 2025-01-15-preview.
 *
 * Other available API versions: 2024-05-01-preview, 2024-06-15-preview, 2024-10-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elastic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getMonitoredSubscription(args: GetMonitoredSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoredSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:elastic:getMonitoredSubscription", {
        "configurationName": args.configurationName,
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMonitoredSubscriptionArgs {
    /**
     * The configuration name. Only 'default' value is supported.
     */
    configurationName: string;
    /**
     * Monitor resource name
     */
    monitorName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * The request to update subscriptions needed to be monitored by the Elastic monitor resource.
 */
export interface GetMonitoredSubscriptionResult {
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The id of the monitored subscription resource.
     */
    readonly id: string;
    /**
     * Name of the monitored subscription resource.
     */
    readonly name: string;
    /**
     * The request to update subscriptions needed to be monitored by the Elastic monitor resource.
     */
    readonly properties: outputs.elastic.SubscriptionListResponse;
    /**
     * The type of the monitored subscription resource.
     */
    readonly type: string;
}
/**
 * The request to update subscriptions needed to be monitored by the Elastic monitor resource.
 *
 * Uses Azure REST API version 2025-01-15-preview.
 *
 * Other available API versions: 2024-05-01-preview, 2024-06-15-preview, 2024-10-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elastic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getMonitoredSubscriptionOutput(args: GetMonitoredSubscriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMonitoredSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:elastic:getMonitoredSubscription", {
        "configurationName": args.configurationName,
        "monitorName": args.monitorName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetMonitoredSubscriptionOutputArgs {
    /**
     * The configuration name. Only 'default' value is supported.
     */
    configurationName: pulumi.Input<string>;
    /**
     * Monitor resource name
     */
    monitorName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
