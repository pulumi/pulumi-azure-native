// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.V20230601.Inputs
{

    /// <summary>
    /// Disk input details.
    /// </summary>
    public sealed class HyperVReplicaAzureDiskInputDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DiskEncryptionSet ARM ID.
        /// </summary>
        [Input("diskEncryptionSetId")]
        public Input<string>? DiskEncryptionSetId { get; set; }

        /// <summary>
        /// The DiskId.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The DiskType.
        /// </summary>
        [Input("diskType")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.V20230601.DiskAccountType>? DiskType { get; set; }

        /// <summary>
        /// The LogStorageAccountId.
        /// </summary>
        [Input("logStorageAccountId")]
        public Input<string>? LogStorageAccountId { get; set; }

        public HyperVReplicaAzureDiskInputDetailsArgs()
        {
        }
        public static new HyperVReplicaAzureDiskInputDetailsArgs Empty => new HyperVReplicaAzureDiskInputDetailsArgs();
    }
}
