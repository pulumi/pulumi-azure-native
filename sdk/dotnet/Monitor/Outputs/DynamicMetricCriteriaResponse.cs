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
    /// Criterion for dynamic threshold.
    /// </summary>
    [OutputType]
    public sealed class DynamicMetricCriteriaResponse
    {
        /// <summary>
        /// The extent of deviation required to trigger an alert. This will affect how tight the threshold is to the metric series pattern.
        /// </summary>
        public readonly string AlertSensitivity;
        /// <summary>
        /// Specifies the type of threshold criteria
        /// Expected value is 'DynamicThresholdCriterion'.
        /// </summary>
        public readonly string CriterionType;
        /// <summary>
        /// List of dimension conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricDimensionResponse> Dimensions;
        /// <summary>
        /// The minimum number of violations required within the selected lookback time window required to raise an alert.
        /// </summary>
        public readonly Outputs.DynamicThresholdFailingPeriodsResponse FailingPeriods;
        /// <summary>
        /// Use this option to set the date from which to start learning the metric historical data and calculate the dynamic thresholds (in ISO8601 format)
        /// </summary>
        public readonly string? IgnoreDataBefore;
        /// <summary>
        /// Name of the metric.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// Namespace of the metric.
        /// </summary>
        public readonly string? MetricNamespace;
        /// <summary>
        /// Name of the criteria.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The operator used to compare the metric value against the threshold.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric validation to be skipped.
        /// </summary>
        public readonly bool? SkipMetricValidation;
        /// <summary>
        /// the criteria time aggregation types.
        /// </summary>
        public readonly string TimeAggregation;

        [OutputConstructor]
        private DynamicMetricCriteriaResponse(
            string alertSensitivity,

            string criterionType,

            ImmutableArray<Outputs.MetricDimensionResponse> dimensions,

            Outputs.DynamicThresholdFailingPeriodsResponse failingPeriods,

            string? ignoreDataBefore,

            string metricName,

            string? metricNamespace,

            string name,

            string @operator,

            bool? skipMetricValidation,

            string timeAggregation)
        {
            AlertSensitivity = alertSensitivity;
            CriterionType = criterionType;
            Dimensions = dimensions;
            FailingPeriods = failingPeriods;
            IgnoreDataBefore = ignoreDataBefore;
            MetricName = metricName;
            MetricNamespace = metricNamespace;
            Name = name;
            Operator = @operator;
            SkipMetricValidation = skipMetricValidation;
            TimeAggregation = timeAggregation;
        }
    }
}
