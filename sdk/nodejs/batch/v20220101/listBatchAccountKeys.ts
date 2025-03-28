// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * This operation applies only to Batch accounts with allowedAuthenticationModes containing 'SharedKey'. If the Batch account doesn't contain 'SharedKey' in its allowedAuthenticationMode, clients cannot use shared keys to authenticate, and must use another allowedAuthenticationModes instead. In this case, getting the keys will fail.
 */
export function listBatchAccountKeys(args: ListBatchAccountKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListBatchAccountKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:batch/v20220101:listBatchAccountKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListBatchAccountKeysArgs {
    /**
     * The name of the Batch account.
     */
    accountName: string;
    /**
     * The name of the resource group that contains the Batch account.
     */
    resourceGroupName: string;
}

/**
 * A set of Azure Batch account keys.
 */
export interface ListBatchAccountKeysResult {
    /**
     * The Batch account name.
     */
    readonly accountName: string;
    /**
     * The primary key associated with the account.
     */
    readonly primary: string;
    /**
     * The secondary key associated with the account.
     */
    readonly secondary: string;
}
/**
 * This operation applies only to Batch accounts with allowedAuthenticationModes containing 'SharedKey'. If the Batch account doesn't contain 'SharedKey' in its allowedAuthenticationMode, clients cannot use shared keys to authenticate, and must use another allowedAuthenticationModes instead. In this case, getting the keys will fail.
 */
export function listBatchAccountKeysOutput(args: ListBatchAccountKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<ListBatchAccountKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:batch/v20220101:listBatchAccountKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListBatchAccountKeysOutputArgs {
    /**
     * The name of the Batch account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account.
     */
    resourceGroupName: pulumi.Input<string>;
}
