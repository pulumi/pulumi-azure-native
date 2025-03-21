// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20230401Preview.Outputs
{

    [OutputType]
    public sealed class ModelPerformanceSignalBaseResponse
    {
        /// <summary>
        /// [Required] The data to calculate drift against.
        /// </summary>
        public readonly Outputs.MonitoringInputDataResponse BaselineData;
        /// <summary>
        /// The data segment.
        /// </summary>
        public readonly Outputs.MonitoringDataSegmentResponse? DataSegment;
        /// <summary>
        /// The amount of time a single monitor should look back over the target data on a given run.
        /// </summary>
        public readonly string? LookbackPeriod;
        /// <summary>
        /// [Required] A list of metrics to calculate and their associated thresholds.
        /// </summary>
        public readonly Union<Outputs.ClassificationModelPerformanceMetricThresholdResponse, Outputs.RegressionModelPerformanceMetricThresholdResponse> MetricThreshold;
        /// <summary>
        /// The current notification mode for this signal.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// 
        /// Expected value is 'ModelPerformance'.
        /// </summary>
        public readonly string SignalType;
        /// <summary>
        /// [Required] The data produced by the production service which drift will be calculated for.
        /// </summary>
        public readonly Outputs.MonitoringInputDataResponse TargetData;

        [OutputConstructor]
        private ModelPerformanceSignalBaseResponse(
            Outputs.MonitoringInputDataResponse baselineData,

            Outputs.MonitoringDataSegmentResponse? dataSegment,

            string? lookbackPeriod,

            Union<Outputs.ClassificationModelPerformanceMetricThresholdResponse, Outputs.RegressionModelPerformanceMetricThresholdResponse> metricThreshold,

            string? mode,

            string signalType,

            Outputs.MonitoringInputDataResponse targetData)
        {
            BaselineData = baselineData;
            DataSegment = dataSegment;
            LookbackPeriod = lookbackPeriod;
            MetricThreshold = metricThreshold;
            Mode = mode;
            SignalType = signalType;
            TargetData = targetData;
        }
    }
}
