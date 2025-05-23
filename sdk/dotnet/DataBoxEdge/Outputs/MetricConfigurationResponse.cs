// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.Outputs
{

    /// <summary>
    /// Metric configuration.
    /// </summary>
    [OutputType]
    public sealed class MetricConfigurationResponse
    {
        /// <summary>
        /// Host name for the IoT hub associated to the device.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricCounterSetResponse> CounterSets;
        /// <summary>
        /// The MDM account to which the counters should be pushed.
        /// </summary>
        public readonly string? MdmAccount;
        /// <summary>
        /// The MDM namespace to which the counters should be pushed. This is required if MDMAccount is specified
        /// </summary>
        public readonly string? MetricNameSpace;
        /// <summary>
        /// The Resource ID on which the metrics should be pushed.
        /// </summary>
        public readonly string ResourceId;

        [OutputConstructor]
        private MetricConfigurationResponse(
            ImmutableArray<Outputs.MetricCounterSetResponse> counterSets,

            string? mdmAccount,

            string? metricNameSpace,

            string resourceId)
        {
            CounterSets = counterSets;
            MdmAccount = mdmAccount;
            MetricNameSpace = metricNameSpace;
            ResourceId = resourceId;
        }
    }
}
