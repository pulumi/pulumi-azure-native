// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230601Preview.Outputs
{

    [OutputType]
    public sealed class FeatureAttributionDriftMonitoringSignalResponse
    {
        /// <summary>
        /// [Required] A list of metrics to calculate and their associated thresholds.
        /// </summary>
        public readonly Outputs.FeatureAttributionMetricThresholdResponse MetricThreshold;
        /// <summary>
        /// The current notification mode for this signal.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// [Required] The data which drift will be calculated for.
        /// </summary>
        public readonly ImmutableArray<object> ProductionData;
        /// <summary>
        /// Property dictionary. Properties can be added, but not removed or altered.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Properties;
        /// <summary>
        /// [Required] The data to calculate drift against.
        /// </summary>
        public readonly object ReferenceData;
        /// <summary>
        /// 
        /// Expected value is 'FeatureAttributionDrift'.
        /// </summary>
        public readonly string SignalType;

        [OutputConstructor]
        private FeatureAttributionDriftMonitoringSignalResponse(
            Outputs.FeatureAttributionMetricThresholdResponse metricThreshold,

            string? mode,

            ImmutableArray<object> productionData,

            ImmutableDictionary<string, string>? properties,

            object referenceData,

            string signalType)
        {
            MetricThreshold = metricThreshold;
            Mode = mode;
            ProductionData = productionData;
            Properties = properties;
            ReferenceData = referenceData;
            SignalType = signalType;
        }
    }
}
