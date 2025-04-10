// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * List all the subscriptions in the private store context
 *
 * Uses Azure REST API version 2023-01-01.
 */
export function listPrivateStoreSubscriptionsContext(args: ListPrivateStoreSubscriptionsContextArgs, opts?: pulumi.InvokeOptions): Promise<ListPrivateStoreSubscriptionsContextResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:marketplace:listPrivateStoreSubscriptionsContext", {
        "privateStoreId": args.privateStoreId,
    }, opts);
}

export interface ListPrivateStoreSubscriptionsContextArgs {
    /**
     * The store ID - must use the tenant ID
     */
    privateStoreId: string;
}

/**
 * List of subscription Ids in the private store
 */
export interface ListPrivateStoreSubscriptionsContextResult {
    readonly subscriptionsIds?: string[];
}
/**
 * List all the subscriptions in the private store context
 *
 * Uses Azure REST API version 2023-01-01.
 */
export function listPrivateStoreSubscriptionsContextOutput(args: ListPrivateStoreSubscriptionsContextOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListPrivateStoreSubscriptionsContextResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:marketplace:listPrivateStoreSubscriptionsContext", {
        "privateStoreId": args.privateStoreId,
    }, opts);
}

export interface ListPrivateStoreSubscriptionsContextOutputArgs {
    /**
     * The store ID - must use the tenant ID
     */
    privateStoreId: pulumi.Input<string>;
}
