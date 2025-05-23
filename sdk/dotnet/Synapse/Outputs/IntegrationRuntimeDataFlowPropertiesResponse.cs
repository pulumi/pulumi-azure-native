// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Outputs
{

    /// <summary>
    /// Data flow properties for managed integration runtime.
    /// </summary>
    [OutputType]
    public sealed class IntegrationRuntimeDataFlowPropertiesResponse
    {
        /// <summary>
        /// Compute type of the cluster which will execute data flow job.
        /// </summary>
        public readonly string? ComputeType;
        /// <summary>
        /// Core count of the cluster which will execute data flow job. Supported values are: 8, 16, 32, 48, 80, 144 and 272.
        /// </summary>
        public readonly int? CoreCount;
        /// <summary>
        /// Time to live (in minutes) setting of the cluster which will execute data flow job.
        /// </summary>
        public readonly int? TimeToLive;

        [OutputConstructor]
        private IntegrationRuntimeDataFlowPropertiesResponse(
            string? computeType,

            int? coreCount,

            int? timeToLive)
        {
            ComputeType = computeType;
            CoreCount = coreCount;
            TimeToLive = timeToLive;
        }
    }
}
