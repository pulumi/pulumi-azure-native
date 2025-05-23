// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageSync.Outputs
{

    /// <summary>
    /// Status of the volume free space policy
    /// </summary>
    [OutputType]
    public sealed class CloudTieringVolumeFreeSpacePolicyStatusResponse
    {
        /// <summary>
        /// Current volume free space percentage.
        /// </summary>
        public readonly int CurrentVolumeFreeSpacePercent;
        /// <summary>
        /// In the case where multiple server endpoints are present in a volume, an effective free space policy is applied.
        /// </summary>
        public readonly int EffectiveVolumeFreeSpacePolicy;
        /// <summary>
        /// Last updated timestamp
        /// </summary>
        public readonly string LastUpdatedTimestamp;

        [OutputConstructor]
        private CloudTieringVolumeFreeSpacePolicyStatusResponse(
            int currentVolumeFreeSpacePercent,

            int effectiveVolumeFreeSpacePolicy,

            string lastUpdatedTimestamp)
        {
            CurrentVolumeFreeSpacePercent = currentVolumeFreeSpacePercent;
            EffectiveVolumeFreeSpacePolicy = effectiveVolumeFreeSpacePolicy;
            LastUpdatedTimestamp = lastUpdatedTimestamp;
        }
    }
}
