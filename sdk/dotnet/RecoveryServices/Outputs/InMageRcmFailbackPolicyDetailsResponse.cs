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
    /// InMageRcm failback specific policy details.
    /// </summary>
    [OutputType]
    public sealed class InMageRcmFailbackPolicyDetailsResponse
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
        /// Expected value is 'InMageRcmFailback'.
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private InMageRcmFailbackPolicyDetailsResponse(
            int? appConsistentFrequencyInMinutes,

            int? crashConsistentFrequencyInMinutes,

            string instanceType)
        {
            AppConsistentFrequencyInMinutes = appConsistentFrequencyInMinutes;
            CrashConsistentFrequencyInMinutes = crashConsistentFrequencyInMinutes;
            InstanceType = instanceType;
        }
    }
}
