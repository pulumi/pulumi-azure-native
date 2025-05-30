// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of awsLogsMetricFilter
    /// </summary>
    [OutputType]
    public sealed class AwsLogsMetricFilterPropertiesResponse
    {
        /// <summary>
        /// The name of the metric filter.
        /// </summary>
        public readonly string? FilterName;
        /// <summary>
        /// A filter pattern for extracting metric data out of ingested log events. For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
        /// </summary>
        public readonly string? FilterPattern;
        /// <summary>
        /// The name of an existing log group that you want to associate with this metric filter.
        /// </summary>
        public readonly string? LogGroupName;
        /// <summary>
        /// The metric transformations.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricTransformationResponse> MetricTransformations;

        [OutputConstructor]
        private AwsLogsMetricFilterPropertiesResponse(
            string? filterName,

            string? filterPattern,

            string? logGroupName,

            ImmutableArray<Outputs.MetricTransformationResponse> metricTransformations)
        {
            FilterName = filterName;
            FilterPattern = filterPattern;
            LogGroupName = logGroupName;
            MetricTransformations = metricTransformations;
        }
    }
}
