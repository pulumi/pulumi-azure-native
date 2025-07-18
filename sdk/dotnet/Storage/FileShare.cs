// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage
{
    /// <summary>
    /// Properties of the file share, including Id, resource name, resource type, Etag.
    /// 
    /// Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:storage:FileShare")]
    public partial class FileShare : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
        /// </summary>
        [Output("accessTier")]
        public Output<string?> AccessTier { get; private set; } = null!;

        /// <summary>
        /// Indicates the last modification time for share access tier.
        /// </summary>
        [Output("accessTierChangeTime")]
        public Output<string> AccessTierChangeTime { get; private set; } = null!;

        /// <summary>
        /// Indicates if there is a pending transition for access tier.
        /// </summary>
        [Output("accessTierStatus")]
        public Output<string> AccessTierStatus { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the share was deleted.
        /// </summary>
        [Output("deleted")]
        public Output<bool> Deleted { get; private set; } = null!;

        /// <summary>
        /// The deleted time if the share was deleted.
        /// </summary>
        [Output("deletedTime")]
        public Output<string> DeletedTime { get; private set; } = null!;

        /// <summary>
        /// The authentication protocol that is used for the file share. Can only be specified when creating a share.
        /// </summary>
        [Output("enabledProtocols")]
        public Output<string?> EnabledProtocols { get; private set; } = null!;

        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// File Share Paid Bursting properties.
        /// </summary>
        [Output("fileSharePaidBursting")]
        public Output<Outputs.FileSharePropertiesResponseFileSharePaidBursting?> FileSharePaidBursting { get; private set; } = null!;

        /// <summary>
        /// The calculated burst IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type.
        /// </summary>
        [Output("includedBurstIops")]
        public Output<int> IncludedBurstIops { get; private set; } = null!;

        /// <summary>
        /// Returns the date and time the share was last modified.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
        /// </summary>
        [Output("leaseDuration")]
        public Output<string> LeaseDuration { get; private set; } = null!;

        /// <summary>
        /// Lease state of the share.
        /// </summary>
        [Output("leaseState")]
        public Output<string> LeaseState { get; private set; } = null!;

        /// <summary>
        /// The lease status of the share.
        /// </summary>
        [Output("leaseStatus")]
        public Output<string> LeaseStatus { get; private set; } = null!;

        /// <summary>
        /// The calculated maximum burst credits for the share. This property is only for file shares created under Files Provisioned v2 account type.
        /// </summary>
        [Output("maxBurstCreditsForIops")]
        public Output<double> MaxBurstCreditsForIops { get; private set; } = null!;

        /// <summary>
        /// A name-value pair to associate with the share as metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Returns the next allowed provisioned bandwidth downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
        /// </summary>
        [Output("nextAllowedProvisionedBandwidthDowngradeTime")]
        public Output<string> NextAllowedProvisionedBandwidthDowngradeTime { get; private set; } = null!;

        /// <summary>
        /// Returns the next allowed provisioned IOPS downgrade time for the share. This property is only for file shares created under Files Provisioned v2 account type.
        /// </summary>
        [Output("nextAllowedProvisionedIopsDowngradeTime")]
        public Output<string> NextAllowedProvisionedIopsDowngradeTime { get; private set; } = null!;

        /// <summary>
        /// Returns the next allowed provisioned storage size downgrade time for the share. This property is only for file shares created under Files Provisioned v1 SSD and Files Provisioned v2 account type
        /// </summary>
        [Output("nextAllowedQuotaDowngradeTime")]
        public Output<string> NextAllowedQuotaDowngradeTime { get; private set; } = null!;

        /// <summary>
        /// The provisioned bandwidth of the share, in mebibytes per second. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned bandwidth.
        /// </summary>
        [Output("provisionedBandwidthMibps")]
        public Output<int?> ProvisionedBandwidthMibps { get; private set; } = null!;

        /// <summary>
        /// The provisioned IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned IOPS.
        /// </summary>
        [Output("provisionedIops")]
        public Output<int?> ProvisionedIops { get; private set; } = null!;

        /// <summary>
        /// Remaining retention days for share that was soft deleted.
        /// </summary>
        [Output("remainingRetentionDays")]
        public Output<int> RemainingRetentionDays { get; private set; } = null!;

        /// <summary>
        /// The property is for NFS share only. The default is NoRootSquash.
        /// </summary>
        [Output("rootSquash")]
        public Output<string?> RootSquash { get; private set; } = null!;

        /// <summary>
        /// The provisioned size of the share, in gibibytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400. For file shares created under Files Provisioned v2 account type, please refer to the GetFileServiceUsage API response for the minimum and maximum allowed provisioned storage size.
        /// </summary>
        [Output("shareQuota")]
        public Output<int?> ShareQuota { get; private set; } = null!;

        /// <summary>
        /// The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
        /// </summary>
        [Output("shareUsageBytes")]
        public Output<double> ShareUsageBytes { get; private set; } = null!;

        /// <summary>
        /// List of stored access policies specified on the share.
        /// </summary>
        [Output("signedIdentifiers")]
        public Output<ImmutableArray<Outputs.SignedIdentifierResponse>> SignedIdentifiers { get; private set; } = null!;

        /// <summary>
        /// Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
        /// </summary>
        [Output("snapshotTime")]
        public Output<string> SnapshotTime { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The version of the share.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a FileShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileShare(string name, FileShareArgs args, CustomResourceOptions? options = null)
            : base("azure-native:storage:FileShare", name, args ?? new FileShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileShare(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:storage:FileShare", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20190401:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20190601:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20200801preview:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210101:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210201:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210401:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210601:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210801:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20210901:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220501:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20220901:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230101:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230401:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20230501:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20240101:FileShare" },
                    new global::Pulumi.Alias { Type = "azure-native:storage/v20250101:FileShare" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FileShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileShare Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FileShare(name, id, options);
        }
    }

    public sealed class FileShareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
        /// </summary>
        [Input("accessTier")]
        public InputUnion<string, Pulumi.AzureNative.Storage.ShareAccessTier>? AccessTier { get; set; }

        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The authentication protocol that is used for the file share. Can only be specified when creating a share.
        /// </summary>
        [Input("enabledProtocols")]
        public InputUnion<string, Pulumi.AzureNative.Storage.EnabledProtocols>? EnabledProtocols { get; set; }

        /// <summary>
        /// Optional, used to expand the properties within share's properties. Valid values are: snapshots. Should be passed as a string with delimiter ','
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// File Share Paid Bursting properties.
        /// </summary>
        [Input("fileSharePaidBursting")]
        public Input<Inputs.FileSharePropertiesFileSharePaidBurstingArgs>? FileSharePaidBursting { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A name-value pair to associate with the share as metadata.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The provisioned bandwidth of the share, in mebibytes per second. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned bandwidth.
        /// </summary>
        [Input("provisionedBandwidthMibps")]
        public Input<int>? ProvisionedBandwidthMibps { get; set; }

        /// <summary>
        /// The provisioned IOPS of the share. This property is only for file shares created under Files Provisioned v2 account type. Please refer to the GetFileServiceUsage API response for the minimum and maximum allowed value for provisioned IOPS.
        /// </summary>
        [Input("provisionedIops")]
        public Input<int>? ProvisionedIops { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The property is for NFS share only. The default is NoRootSquash.
        /// </summary>
        [Input("rootSquash")]
        public InputUnion<string, Pulumi.AzureNative.Storage.RootSquashType>? RootSquash { get; set; }

        /// <summary>
        /// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
        /// </summary>
        [Input("shareName")]
        public Input<string>? ShareName { get; set; }

        /// <summary>
        /// The provisioned size of the share, in gibibytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400. For file shares created under Files Provisioned v2 account type, please refer to the GetFileServiceUsage API response for the minimum and maximum allowed provisioned storage size.
        /// </summary>
        [Input("shareQuota")]
        public Input<int>? ShareQuota { get; set; }

        [Input("signedIdentifiers")]
        private InputList<Inputs.SignedIdentifierArgs>? _signedIdentifiers;

        /// <summary>
        /// List of stored access policies specified on the share.
        /// </summary>
        public InputList<Inputs.SignedIdentifierArgs> SignedIdentifiers
        {
            get => _signedIdentifiers ?? (_signedIdentifiers = new InputList<Inputs.SignedIdentifierArgs>());
            set => _signedIdentifiers = value;
        }

        public FileShareArgs()
        {
        }
        public static new FileShareArgs Empty => new FileShareArgs();
    }
}
