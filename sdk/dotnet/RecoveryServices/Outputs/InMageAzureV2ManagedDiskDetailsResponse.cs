// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// InMageAzureV2 Managed disk details.
    /// </summary>
    [OutputType]
    public sealed class InMageAzureV2ManagedDiskDetailsResponse
    {
        /// <summary>
        /// The DiskEncryptionSet ARM ID.
        /// </summary>
        public readonly string? DiskEncryptionSetId;
        /// <summary>
        /// The disk id.
        /// </summary>
        public readonly string? DiskId;
        /// <summary>
        /// The replica disk type.
        /// </summary>
        public readonly string? ReplicaDiskType;
        /// <summary>
        /// Seed managed disk Id.
        /// </summary>
        public readonly string? SeedManagedDiskId;
        /// <summary>
        /// The target disk name.
        /// </summary>
        public readonly string? TargetDiskName;

        [OutputConstructor]
        private InMageAzureV2ManagedDiskDetailsResponse(
            string? diskEncryptionSetId,

            string? diskId,

            string? replicaDiskType,

            string? seedManagedDiskId,

            string? targetDiskName)
        {
            DiskEncryptionSetId = diskEncryptionSetId;
            DiskId = diskId;
            ReplicaDiskType = replicaDiskType;
            SeedManagedDiskId = seedManagedDiskId;
            TargetDiskName = targetDiskName;
        }
    }
}
