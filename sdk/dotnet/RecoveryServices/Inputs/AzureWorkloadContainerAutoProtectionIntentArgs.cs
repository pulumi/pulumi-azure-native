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
    /// Azure workload specific protection intent item.
    /// </summary>
    public sealed class AzureWorkloadContainerAutoProtectionIntentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of backup management for the backed up item.
        /// </summary>
        [Input("backupManagementType")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.BackupManagementType>? BackupManagementType { get; set; }

        /// <summary>
        /// ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        /// </summary>
        [Input("itemId")]
        public Input<string>? ItemId { get; set; }

        /// <summary>
        /// ID of the backup policy with which this item is backed up.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// backup protectionIntent type.
        /// Expected value is 'AzureWorkloadContainerAutoProtectionIntent'.
        /// </summary>
        [Input("protectionIntentItemType", required: true)]
        public Input<string> ProtectionIntentItemType { get; set; } = null!;

        /// <summary>
        /// Backup state of this backup item.
        /// </summary>
        [Input("protectionState")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.ProtectionStatus>? ProtectionState { get; set; }

        /// <summary>
        /// ARM ID of the resource to be backed up.
        /// </summary>
        [Input("sourceResourceId")]
        public Input<string>? SourceResourceId { get; set; }

        public AzureWorkloadContainerAutoProtectionIntentArgs()
        {
        }
        public static new AzureWorkloadContainerAutoProtectionIntentArgs Empty => new AzureWorkloadContainerAutoProtectionIntentArgs();
    }
}
