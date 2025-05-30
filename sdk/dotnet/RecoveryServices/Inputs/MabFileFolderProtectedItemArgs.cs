// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// MAB workload-specific backup item.
    /// </summary>
    public sealed class MabFileFolderProtectedItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the backup set the backup item belongs to
        /// </summary>
        [Input("backupSetName")]
        public Input<string>? BackupSetName { get; set; }

        /// <summary>
        /// Name of the computer associated with this backup item.
        /// </summary>
        [Input("computerName")]
        public Input<string>? ComputerName { get; set; }

        /// <summary>
        /// Unique name of container
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
        /// </summary>
        [Input("createMode")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.CreateMode>? CreateMode { get; set; }

        /// <summary>
        /// Sync time for deferred deletion in UTC
        /// </summary>
        [Input("deferredDeleteSyncTimeInUTC")]
        public Input<double>? DeferredDeleteSyncTimeInUTC { get; set; }

        /// <summary>
        /// Time for deferred deletion in UTC
        /// </summary>
        [Input("deferredDeleteTimeInUTC")]
        public Input<string>? DeferredDeleteTimeInUTC { get; set; }

        /// <summary>
        /// Time remaining before the DS marked for deferred delete is permanently deleted
        /// </summary>
        [Input("deferredDeleteTimeRemaining")]
        public Input<string>? DeferredDeleteTimeRemaining { get; set; }

        /// <summary>
        /// Additional information with this backup item.
        /// </summary>
        [Input("extendedInfo")]
        public Input<Inputs.MabFileFolderProtectedItemExtendedInfoArgs>? ExtendedInfo { get; set; }

        /// <summary>
        /// Friendly name of this backup item.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Flag to identify whether datasource is protected in archive
        /// </summary>
        [Input("isArchiveEnabled")]
        public Input<bool>? IsArchiveEnabled { get; set; }

        /// <summary>
        /// Flag to identify whether the deferred deleted DS is to be purged soon
        /// </summary>
        [Input("isDeferredDeleteScheduleUpcoming")]
        public Input<bool>? IsDeferredDeleteScheduleUpcoming { get; set; }

        /// <summary>
        /// Flag to identify that deferred deleted DS is to be moved into Pause state
        /// </summary>
        [Input("isRehydrate")]
        public Input<bool>? IsRehydrate { get; set; }

        /// <summary>
        /// Flag to identify whether the DS is scheduled for deferred delete
        /// </summary>
        [Input("isScheduledForDeferredDelete")]
        public Input<bool>? IsScheduledForDeferredDelete { get; set; }

        /// <summary>
        /// Status of last backup operation.
        /// </summary>
        [Input("lastBackupStatus")]
        public Input<string>? LastBackupStatus { get; set; }

        /// <summary>
        /// Timestamp of the last backup operation on this backup item.
        /// </summary>
        [Input("lastBackupTime")]
        public Input<string>? LastBackupTime { get; set; }

        /// <summary>
        /// Timestamp when the last (latest) backup copy was created for this backup item.
        /// </summary>
        [Input("lastRecoveryPoint")]
        public Input<string>? LastRecoveryPoint { get; set; }

        /// <summary>
        /// ID of the backup policy with which this item is backed up.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Name of the policy used for protection
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// backup item type.
        /// Expected value is 'MabFileFolderProtectedItem'.
        /// </summary>
        [Input("protectedItemType", required: true)]
        public Input<string> ProtectedItemType { get; set; } = null!;

        /// <summary>
        /// Protected, ProtectionStopped, IRPending or ProtectionError
        /// </summary>
        [Input("protectionState")]
        public Input<string>? ProtectionState { get; set; }

        [Input("resourceGuardOperationRequests")]
        private InputList<string>? _resourceGuardOperationRequests;

        /// <summary>
        /// ResourceGuardOperationRequests on which LAC check will be performed
        /// </summary>
        public InputList<string> ResourceGuardOperationRequests
        {
            get => _resourceGuardOperationRequests ?? (_resourceGuardOperationRequests = new InputList<string>());
            set => _resourceGuardOperationRequests = value;
        }

        /// <summary>
        /// Soft delete retention period in days
        /// </summary>
        [Input("softDeleteRetentionPeriodInDays")]
        public Input<int>? SoftDeleteRetentionPeriodInDays { get; set; }

        /// <summary>
        /// ARM ID of the resource to be backed up.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        public MabFileFolderProtectedItemArgs()
        {
        }
        public static new MabFileFolderProtectedItemArgs Empty => new MabFileFolderProtectedItemArgs();
    }
}
