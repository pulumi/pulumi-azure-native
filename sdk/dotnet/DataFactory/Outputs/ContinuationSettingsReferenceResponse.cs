// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Continuation settings for execute data flow activity.
    /// </summary>
    [OutputType]
    public sealed class ContinuationSettingsReferenceResponse
    {
        /// <summary>
        /// Continuation TTL in minutes.
        /// </summary>
        public readonly object? ContinuationTtlInMinutes;
        /// <summary>
        /// Customized checkpoint key.
        /// </summary>
        public readonly object? CustomizedCheckpointKey;
        /// <summary>
        /// Idle condition.
        /// </summary>
        public readonly object? IdleCondition;

        [OutputConstructor]
        private ContinuationSettingsReferenceResponse(
            object? continuationTtlInMinutes,

            object? customizedCheckpointKey,

            object? idleCondition)
        {
            ContinuationTtlInMinutes = continuationTtlInMinutes;
            CustomizedCheckpointKey = customizedCheckpointKey;
            IdleCondition = idleCondition;
        }
    }
}
