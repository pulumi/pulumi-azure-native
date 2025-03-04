// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.V20210301.Outputs
{

    /// <summary>
    /// An indication of Cache health. Gives more information about health than just that related to provisioning.
    /// </summary>
    [OutputType]
    public sealed class CacheHealthResponse
    {
        /// <summary>
        /// Outstanding conditions that need to be investigated and resolved.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConditionResponse> Conditions;
        /// <summary>
        /// List of Cache health states.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Describes explanation of state.
        /// </summary>
        public readonly string? StatusDescription;

        [OutputConstructor]
        private CacheHealthResponse(
            ImmutableArray<Outputs.ConditionResponse> conditions,

            string? state,

            string? statusDescription)
        {
            Conditions = conditions;
            State = state;
            StatusDescription = statusDescription;
        }
    }
}
