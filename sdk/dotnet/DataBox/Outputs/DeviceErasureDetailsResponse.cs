// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Outputs
{

    /// <summary>
    /// Device erasure details with erasure completion status and erasureordestructionlog sas key
    /// </summary>
    [OutputType]
    public sealed class DeviceErasureDetailsResponse
    {
        /// <summary>
        /// Holds the device erasure completion status
        /// </summary>
        public readonly string DeviceErasureStatus;
        /// <summary>
        /// Shared access key to download cleanup or destruction certificate for device
        /// </summary>
        public readonly string ErasureOrDestructionCertificateSasKey;

        [OutputConstructor]
        private DeviceErasureDetailsResponse(
            string deviceErasureStatus,

            string erasureOrDestructionCertificateSasKey)
        {
            DeviceErasureStatus = deviceErasureStatus;
            ErasureOrDestructionCertificateSasKey = erasureOrDestructionCertificateSasKey;
        }
    }
}
