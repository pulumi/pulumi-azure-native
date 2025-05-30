// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorSimple.Outputs
{

    /// <summary>
    /// The eligibility result of device, as a failover target device.
    /// </summary>
    [OutputType]
    public sealed class TargetEligibilityResultResponse
    {
        /// <summary>
        /// The eligibility status of device, as a failover target device.
        /// </summary>
        public readonly string? EligibilityStatus;
        /// <summary>
        /// The list of error messages, if a device does not qualify as a failover target device.
        /// </summary>
        public readonly ImmutableArray<Outputs.TargetEligibilityErrorMessageResponse> Messages;

        [OutputConstructor]
        private TargetEligibilityResultResponse(
            string? eligibilityStatus,

            ImmutableArray<Outputs.TargetEligibilityErrorMessageResponse> messages)
        {
            EligibilityStatus = eligibilityStatus;
            Messages = messages;
        }
    }
}
