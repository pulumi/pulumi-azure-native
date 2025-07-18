// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp
{
    /// <summary>
    /// Backup under a Backup Vault
    /// 
    /// Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2022-11-01.
    /// 
    /// Other available API versions: 2022-11-01-preview, 2023-05-01-preview, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01-preview, 2025-01-01, 2025-01-01-preview, 2025-03-01, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:netapp:Backup")]
    public partial class Backup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// UUID v4 used to identify the Backup
        /// </summary>
        [Output("backupId")]
        public Output<string> BackupId { get; private set; } = null!;

        /// <summary>
        /// ResourceId used to identify the backup policy
        /// </summary>
        [Output("backupPolicyResourceId")]
        public Output<string> BackupPolicyResourceId { get; private set; } = null!;

        /// <summary>
        /// Type of backup Manual or Scheduled
        /// </summary>
        [Output("backupType")]
        public Output<string> BackupType { get; private set; } = null!;

        /// <summary>
        /// The creation date of the backup
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// Failure reason
        /// </summary>
        [Output("failureReason")]
        public Output<string> FailureReason { get; private set; } = null!;

        /// <summary>
        /// Label for backup
        /// </summary>
        [Output("label")]
        public Output<string?> Label { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Size of backup in bytes
        /// </summary>
        [Output("size")]
        public Output<double> Size { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot
        /// </summary>
        [Output("snapshotName")]
        public Output<string?> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
        /// </summary>
        [Output("useExistingSnapshot")]
        public Output<bool?> UseExistingSnapshot { get; private set; } = null!;

        /// <summary>
        /// ResourceId used to identify the Volume
        /// </summary>
        [Output("volumeResourceId")]
        public Output<string> VolumeResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("azure-native:netapp:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:netapp:Backup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20221101:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20221101preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20230501preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20230701preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20231101:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20231101preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240101:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240301:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240301preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240501:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240501preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240701:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240701preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240901:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20240901preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20250101:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20250101preview:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20250301:Backup" },
                    new global::Pulumi.Alias { Type = "azure-native:netapp/v20250301preview:Backup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, options);
        }
    }

    public sealed class BackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the backup
        /// </summary>
        [Input("backupName")]
        public Input<string>? BackupName { get; set; }

        /// <summary>
        /// The name of the Backup Vault
        /// </summary>
        [Input("backupVaultName", required: true)]
        public Input<string> BackupVaultName { get; set; } = null!;

        /// <summary>
        /// Label for backup
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the snapshot
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
        /// </summary>
        [Input("useExistingSnapshot")]
        public Input<bool>? UseExistingSnapshot { get; set; }

        /// <summary>
        /// ResourceId used to identify the Volume
        /// </summary>
        [Input("volumeResourceId", required: true)]
        public Input<string> VolumeResourceId { get; set; } = null!;

        public BackupArgs()
        {
            UseExistingSnapshot = false;
        }
        public static new BackupArgs Empty => new BackupArgs();
    }
}
