// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// Definition of Prometheus metrics forwarding configuration.
    /// </summary>
    [OutputType]
    public sealed class PrometheusForwarderDataSourceResponse
    {
        /// <summary>
        /// The list of label inclusion filters in the form of label "name-value" pairs.
        /// Currently only one label is supported: 'microsoft_metrics_include_label'.
        /// Label values are matched case-insensitively.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? LabelIncludeFilter;
        /// <summary>
        /// A friendly name for the data source. 
        /// This name should be unique across all data sources (regardless of type) within the data collection rule.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// List of streams that this data source will be sent to.
        /// </summary>
        public readonly ImmutableArray<string> Streams;

        [OutputConstructor]
        private PrometheusForwarderDataSourceResponse(
            ImmutableDictionary<string, string>? labelIncludeFilter,

            string? name,

            ImmutableArray<string> streams)
        {
            LabelIncludeFilter = labelIncludeFilter;
            Name = name;
            Streams = streams;
        }
    }
}
