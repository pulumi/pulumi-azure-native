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
    /// InMageRcm un-protected disk details.
    /// </summary>
    [OutputType]
    public sealed class InMageRcmUnProtectedDiskDetailsResponse
    {
        /// <summary>
        /// The disk capacity in bytes.
        /// </summary>
        public readonly double CapacityInBytes;
        /// <summary>
        /// The disk Id.
        /// </summary>
        public readonly string DiskId;
        /// <summary>
        /// The disk name.
        /// </summary>
        public readonly string DiskName;

        [OutputConstructor]
        private InMageRcmUnProtectedDiskDetailsResponse(
            double capacityInBytes,

            string diskId,

            string diskName)
        {
            CapacityInBytes = capacityInBytes;
            DiskId = diskId;
            DiskName = diskName;
        }
    }
}
