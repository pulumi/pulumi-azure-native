// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Get synchronization settings set on a share
 *
 * Uses Azure REST API version 2021-08-01.
 */
export function listShareSubscriptionSourceShareSynchronizationSettings(args: ListShareSubscriptionSourceShareSynchronizationSettingsArgs, opts?: pulumi.InvokeOptions): Promise<ListShareSubscriptionSourceShareSynchronizationSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:datashare:listShareSubscriptionSourceShareSynchronizationSettings", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "shareSubscriptionName": args.shareSubscriptionName,
        "skipToken": args.skipToken,
    }, opts);
}

export interface ListShareSubscriptionSourceShareSynchronizationSettingsArgs {
    /**
     * The name of the share account.
     */
    accountName: string;
    /**
     * The resource group name.
     */
    resourceGroupName: string;
    /**
     * The name of the shareSubscription.
     */
    shareSubscriptionName: string;
    /**
     * Continuation token
     */
    skipToken?: string;
}

/**
 * List response for get source share Synchronization settings
 */
export interface ListShareSubscriptionSourceShareSynchronizationSettingsResult {
    /**
     * The Url of next result page.
     */
    readonly nextLink?: string;
    /**
     * Collection of items of type DataTransferObjects.
     */
    readonly value: outputs.datashare.ScheduledSourceSynchronizationSettingResponse[];
}
/**
 * Get synchronization settings set on a share
 *
 * Uses Azure REST API version 2021-08-01.
 */
export function listShareSubscriptionSourceShareSynchronizationSettingsOutput(args: ListShareSubscriptionSourceShareSynchronizationSettingsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListShareSubscriptionSourceShareSynchronizationSettingsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:datashare:listShareSubscriptionSourceShareSynchronizationSettings", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "shareSubscriptionName": args.shareSubscriptionName,
        "skipToken": args.skipToken,
    }, opts);
}

export interface ListShareSubscriptionSourceShareSynchronizationSettingsOutputArgs {
    /**
     * The name of the share account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the shareSubscription.
     */
    shareSubscriptionName: pulumi.Input<string>;
    /**
     * Continuation token
     */
    skipToken?: pulumi.Input<string>;
}
