// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class FeatureAttributionDriftMonitoringSignalArgs : global::Pulumi.ResourceArgs
    {
        [Input("featureDataTypeOverride")]
        private InputMap<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringFeatureDataType>>? _featureDataTypeOverride;

        /// <summary>
        /// A dictionary that maps feature names to their respective data types.
        /// </summary>
        public InputMap<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringFeatureDataType>> FeatureDataTypeOverride
        {
            get => _featureDataTypeOverride ?? (_featureDataTypeOverride = new InputMap<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringFeatureDataType>>());
            set => _featureDataTypeOverride = value;
        }

        /// <summary>
        /// [Required] The settings for computing feature importance.
        /// </summary>
        [Input("featureImportanceSettings", required: true)]
        public Input<Inputs.FeatureImportanceSettingsArgs> FeatureImportanceSettings { get; set; } = null!;

        /// <summary>
        /// [Required] A list of metrics to calculate and their associated thresholds.
        /// </summary>
        [Input("metricThreshold", required: true)]
        public Input<Inputs.FeatureAttributionMetricThresholdArgs> MetricThreshold { get; set; } = null!;

        [Input("notificationTypes")]
        private InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringNotificationType>>? _notificationTypes;

        /// <summary>
        /// The current notification mode for this signal.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringNotificationType>> NotificationTypes
        {
            get => _notificationTypes ?? (_notificationTypes = new InputList<Union<string, Pulumi.AzureNative.MachineLearningServices.MonitoringNotificationType>>());
            set => _notificationTypes = value;
        }

        [Input("productionData", required: true)]
        private InputList<object>? _productionData;

        /// <summary>
        /// [Required] The data which drift will be calculated for.
        /// </summary>
        public InputList<object> ProductionData
        {
            get => _productionData ?? (_productionData = new InputList<object>());
            set => _productionData = value;
        }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Property dictionary. Properties can be added, but not removed or altered.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// [Required] The data to calculate drift against.
        /// </summary>
        [Input("referenceData", required: true)]
        public object ReferenceData { get; set; } = null!;

        /// <summary>
        /// 
        /// Expected value is 'FeatureAttributionDrift'.
        /// </summary>
        [Input("signalType", required: true)]
        public Input<string> SignalType { get; set; } = null!;

        public FeatureAttributionDriftMonitoringSignalArgs()
        {
        }
        public static new FeatureAttributionDriftMonitoringSignalArgs Empty => new FeatureAttributionDriftMonitoringSignalArgs();
    }
}
