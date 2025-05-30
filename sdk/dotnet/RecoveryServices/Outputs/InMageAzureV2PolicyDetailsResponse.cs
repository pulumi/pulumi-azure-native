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
    /// InMage Azure v2 specific protection profile details.
    /// </summary>
    [OutputType]
    public sealed class InMageAzureV2PolicyDetailsResponse
    {
        /// <summary>
        /// The app consistent snapshot frequency in minutes.
        /// </summary>
        public readonly int? AppConsistentFrequencyInMinutes;
        /// <summary>
        /// The crash consistent snapshot frequency in minutes.
        /// </summary>
        public readonly int? CrashConsistentFrequencyInMinutes;
        /// <summary>
        /// Gets the class type. Overridden in derived classes.
        /// Expected value is 'InMageAzureV2'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// A value indicating whether multi-VM sync has to be enabled.
        /// </summary>
        public readonly string? MultiVmSyncStatus;
        /// <summary>
        /// The duration in minutes until which the recovery points need to be stored.
        /// </summary>
        public readonly int? RecoveryPointHistory;
        /// <summary>
        /// The recovery point threshold in minutes.
        /// </summary>
        public readonly int? RecoveryPointThresholdInMinutes;

        [OutputConstructor]
        private InMageAzureV2PolicyDetailsResponse(
            int? appConsistentFrequencyInMinutes,

            int? crashConsistentFrequencyInMinutes,

            string instanceType,

            string? multiVmSyncStatus,

            int? recoveryPointHistory,

            int? recoveryPointThresholdInMinutes)
        {
            AppConsistentFrequencyInMinutes = appConsistentFrequencyInMinutes;
            CrashConsistentFrequencyInMinutes = crashConsistentFrequencyInMinutes;
            InstanceType = instanceType;
            MultiVmSyncStatus = multiVmSyncStatus;
            RecoveryPointHistory = recoveryPointHistory;
            RecoveryPointThresholdInMinutes = recoveryPointThresholdInMinutes;
        }
    }
}
