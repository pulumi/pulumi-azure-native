// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets properties of a specified share.
 *
 * Uses Azure REST API version 2024-01-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getFileShare(args: GetFileShareArgs, opts?: pulumi.InvokeOptions): Promise<GetFileShareResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:storage:getFileShare", {
        "accountName": args.accountName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "shareName": args.shareName,
    }, opts);
}

export interface GetFileShareArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: string;
    /**
     * Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
     */
    expand?: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: string;
    /**
     * The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
     */
    shareName: string;
}

/**
 * Properties of the file share, including Id, resource name, resource type, Etag.
 */
export interface GetFileShareResult {
    /**
     * Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
     */
    readonly accessTier?: string;
    /**
     * Indicates the last modification time for share access tier.
     */
    readonly accessTierChangeTime: string;
    /**
     * Indicates if there is a pending transition for access tier.
     */
    readonly accessTierStatus: string;
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * Indicates whether the share was deleted.
     */
    readonly deleted: boolean;
    /**
     * The deleted time if the share was deleted.
     */
    readonly deletedTime: string;
    /**
     * The authentication protocol that is used for the file share. Can only be specified when creating a share.
     */
    readonly enabledProtocols?: string;
    /**
     * Resource Etag.
     */
    readonly etag: string;
    /**
     * File Share Paid Bursting properties.
     */
    readonly fileSharePaidBursting?: outputs.storage.FileSharePropertiesResponseFileSharePaidBursting;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The calculated burst IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type.
     */
    readonly includedBurstIops: number;
    /**
     * Returns the date and time the share was last modified.
     */
    readonly lastModifiedTime: string;
    /**
     * Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
     */
    readonly leaseDuration: string;
    /**
     * Lease state of the share.
     */
    readonly leaseState: string;
    /**
     * The lease status of the share.
     */
    readonly leaseStatus: string;
    /**
     * The calculated maximum burst credits for the share. This property is only for file shares created under Files Provisioned v2 account type.
     */
    readonly maxBurstCreditsForIops: number;
    /**
     * A name-value pair to associate with the share as metadata.
     */
    readonly metadata?: {[key: string]: string};
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Returns the next allowed provisioned bandwidth downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
     */
    readonly nextAllowedProvisionedBandwidthDowngradeTime: string;
    /**
     * Returns the next allowed provisioned IOPS downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
     */
    readonly nextAllowedProvisionedIopsDowngradeTime: string;
    /**
     * Returns the next allowed provisioned storage size downgrade time for the share. This property is only for file shares created under Files Provisioned v1 SSD and Files Provisioned v2 account type
     */
    readonly nextAllowedQuotaDowngradeTime: string;
    /**
     * The provisioned bandwidth of the share, in mebibytes per second. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned bandwidth.
     */
    readonly provisionedBandwidthMibps?: number;
    /**
     * The provisioned IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned IOPS.
     */
    readonly provisionedIops?: number;
    /**
     * Remaining retention days for share that was soft deleted.
     */
    readonly remainingRetentionDays: number;
    /**
     * The property is for NFS share only. The default is NoRootSquash.
     */
    readonly rootSquash?: string;
    /**
     * The provisioned size of the share, in gibibytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400. For file shares created under Files Provisioned v2 account type, please refer to the GetFileServiceUsage API response for the minimum and maximum allowed provisioned storage size.
     */
    readonly shareQuota?: number;
    /**
     * The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
     */
    readonly shareUsageBytes: number;
    /**
     * List of stored access policies specified on the share.
     */
    readonly signedIdentifiers?: outputs.storage.SignedIdentifierResponse[];
    /**
     * Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
     */
    readonly snapshotTime: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The version of the share.
     */
    readonly version: string;
}
/**
 * Gets properties of a specified share.
 *
 * Uses Azure REST API version 2024-01-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getFileShareOutput(args: GetFileShareOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFileShareResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:storage:getFileShare", {
        "accountName": args.accountName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "shareName": args.shareName,
    }, opts);
}

export interface GetFileShareOutputArgs {
    /**
     * The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
     */
    accountName: pulumi.Input<string>;
    /**
     * Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
     */
    shareName: pulumi.Input<string>;
}
