// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Impact.Inputs
{

    /// <summary>
    /// Details about impacted performance metrics. Applicable for performance related impact
    /// </summary>
    public sealed class PerformanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Observed value for the metric
        /// </summary>
        [Input("actual")]
        public Input<double>? Actual { get; set; }

        /// <summary>
        /// Threshold value for the metric
        /// </summary>
        [Input("expected")]
        public Input<double>? Expected { get; set; }

        /// <summary>
        /// Max and Min Threshold values for the metric
        /// </summary>
        [Input("expectedValueRange")]
        public Input<Inputs.ExpectedValueRangeArgs>? ExpectedValueRange { get; set; }

        /// <summary>
        /// Name of the Metric examples:  Disk, IOPs, CPU, GPU, Memory, details can be found from /impactCategories API
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// Unit of the metric ex: Bytes, Percentage, Count, Seconds, Milliseconds, Bytes/Second, Count/Second, etc.., Other
        /// </summary>
        [Input("unit")]
        public InputUnion<string, Pulumi.AzureNative.Impact.MetricUnit>? Unit { get; set; }

        public PerformanceArgs()
        {
        }
        public static new PerformanceArgs Empty => new PerformanceArgs();
    }
}
