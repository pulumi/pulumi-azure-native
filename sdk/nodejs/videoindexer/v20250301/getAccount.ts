// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the properties of an Azure Video Indexer account.
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:videoindexer/v20250301:getAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccountArgs {
    /**
     * The name of the Azure Video Indexer account.
     */
    accountName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * An Azure Video Indexer account.
 */
export interface GetAccountResult {
    /**
     * The account's data-plane ID. This can be set only when connecting an existing classic account
     */
    readonly accountId?: string;
    /**
     * The account's name
     */
    readonly accountName: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Managed service identity (system assigned and/or user assigned identities)
     */
    readonly identity?: outputs.videoindexer.v20250301.ManagedServiceIdentityResponse;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The openAi services details
     */
    readonly openAiServices?: outputs.videoindexer.v20250301.OpenAiServicesForPutRequestResponse;
    /**
     * Gets the status of the account at the time the operation was called.
     */
    readonly provisioningState: string;
    /**
     * The storage services details
     */
    readonly storageServices?: outputs.videoindexer.v20250301.StorageServicesForPutRequestResponse;
    /**
     * The system meta data relating to this resource.
     */
    readonly systemData: outputs.videoindexer.v20250301.SystemDataResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The account's tenant id
     */
    readonly tenantId: string;
    /**
     * An integer representing the total minutes that have been indexed on the account
     */
    readonly totalMinutesIndexed: number;
    /**
     * An integer representing the total seconds that have been indexed on the account
     */
    readonly totalSecondsIndexed: number;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets the properties of an Azure Video Indexer account.
 */
export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAccountResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:videoindexer/v20250301:getAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAccountOutputArgs {
    /**
     * The name of the Azure Video Indexer account.
     */
    accountName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
