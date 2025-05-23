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
    /// Server endpoint cloud tiering status object.
    /// </summary>
    [OutputType]
    public sealed class ServerEndpointCloudTieringStatusResponse
    {
        /// <summary>
        /// Information regarding how well the local cache on the server is performing.
        /// </summary>
        public readonly Outputs.CloudTieringCachePerformanceResponse CachePerformance;
        /// <summary>
        /// Status of the date policy
        /// </summary>
        public readonly Outputs.CloudTieringDatePolicyStatusResponse DatePolicyStatus;
        /// <summary>
        /// Information regarding files that failed to be tiered
        /// </summary>
        public readonly Outputs.CloudTieringFilesNotTieringResponse FilesNotTiering;
        /// <summary>
        /// Cloud tiering health state.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// The last updated timestamp of health state
        /// </summary>
        public readonly string HealthLastUpdatedTimestamp;
        /// <summary>
        /// Last cloud tiering result (HResult)
        /// </summary>
        public readonly int LastCloudTieringResult;
        /// <summary>
        /// Last cloud tiering success timestamp
        /// </summary>
        public readonly string LastSuccessTimestamp;
        /// <summary>
        /// Last updated timestamp
        /// </summary>
        public readonly string LastUpdatedTimestamp;
        /// <summary>
        /// Information regarding the low disk mode state
        /// </summary>
        public readonly Outputs.CloudTieringLowDiskModeResponse LowDiskMode;
        /// <summary>
        /// Information regarding how much local space cloud tiering is saving.
        /// </summary>
        public readonly Outputs.CloudTieringSpaceSavingsResponse SpaceSavings;
        /// <summary>
        /// Status of the volume free space policy
        /// </summary>
        public readonly Outputs.CloudTieringVolumeFreeSpacePolicyStatusResponse VolumeFreeSpacePolicyStatus;

        [OutputConstructor]
        private ServerEndpointCloudTieringStatusResponse(
            Outputs.CloudTieringCachePerformanceResponse cachePerformance,

            Outputs.CloudTieringDatePolicyStatusResponse datePolicyStatus,

            Outputs.CloudTieringFilesNotTieringResponse filesNotTiering,

            string health,

            string healthLastUpdatedTimestamp,

            int lastCloudTieringResult,

            string lastSuccessTimestamp,

            string lastUpdatedTimestamp,

            Outputs.CloudTieringLowDiskModeResponse lowDiskMode,

            Outputs.CloudTieringSpaceSavingsResponse spaceSavings,

            Outputs.CloudTieringVolumeFreeSpacePolicyStatusResponse volumeFreeSpacePolicyStatus)
        {
            CachePerformance = cachePerformance;
            DatePolicyStatus = datePolicyStatus;
            FilesNotTiering = filesNotTiering;
            Health = health;
            HealthLastUpdatedTimestamp = healthLastUpdatedTimestamp;
            LastCloudTieringResult = lastCloudTieringResult;
            LastSuccessTimestamp = lastSuccessTimestamp;
            LastUpdatedTimestamp = lastUpdatedTimestamp;
            LowDiskMode = lowDiskMode;
            SpaceSavings = spaceSavings;
            VolumeFreeSpacePolicyStatus = volumeFreeSpacePolicyStatus;
        }
    }
}
