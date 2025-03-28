// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Dapr PubSub Event Subscription.
 */
export function getDaprSubscription(args: GetDaprSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetDaprSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:app/v20230801preview:getDaprSubscription", {
        "environmentName": args.environmentName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDaprSubscriptionArgs {
    /**
     * Name of the Managed Environment.
     */
    environmentName: string;
    /**
     * Name of the Dapr subscription.
     */
    name: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Dapr PubSub Event Subscription.
 */
export interface GetDaprSubscriptionResult {
    /**
     * Bulk subscription options
     */
    readonly bulkSubscribe?: outputs.app.v20230801preview.DaprSubscriptionBulkSubscribeOptionsResponse;
    /**
     * Deadletter topic name
     */
    readonly deadLetterTopic?: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Subscription metadata
     */
    readonly metadata?: {[key: string]: string};
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Dapr PubSub component name
     */
    readonly pubsubName?: string;
    /**
     * Subscription routes
     */
    readonly routes?: outputs.app.v20230801preview.DaprSubscriptionRoutesResponse;
    /**
     * Application scopes to restrict the subscription to specific apps.
     */
    readonly scopes?: string[];
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.app.v20230801preview.SystemDataResponse;
    /**
     * Topic name
     */
    readonly topic?: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Dapr PubSub Event Subscription.
 */
export function getDaprSubscriptionOutput(args: GetDaprSubscriptionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDaprSubscriptionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:app/v20230801preview:getDaprSubscription", {
        "environmentName": args.environmentName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDaprSubscriptionOutputArgs {
    /**
     * Name of the Managed Environment.
     */
    environmentName: pulumi.Input<string>;
    /**
     * Name of the Dapr subscription.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
