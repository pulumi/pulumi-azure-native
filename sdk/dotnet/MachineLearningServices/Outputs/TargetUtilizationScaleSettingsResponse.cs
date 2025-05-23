// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class TargetUtilizationScaleSettingsResponse
    {
        /// <summary>
        /// The maximum number of instances that the deployment can scale to. The quota will be reserved for max_instances.
        /// </summary>
        public readonly int? MaxInstances;
        /// <summary>
        /// The minimum number of instances to always be present.
        /// </summary>
        public readonly int? MinInstances;
        /// <summary>
        /// The polling interval in ISO 8691 format. Only supports duration with precision as low as Seconds.
        /// </summary>
        public readonly string? PollingInterval;
        /// <summary>
        /// 
        /// Expected value is 'TargetUtilization'.
        /// </summary>
        public readonly string ScaleType;
        /// <summary>
        /// Target CPU usage for the autoscaler.
        /// </summary>
        public readonly int? TargetUtilizationPercentage;

        [OutputConstructor]
        private TargetUtilizationScaleSettingsResponse(
            int? maxInstances,

            int? minInstances,

            string? pollingInterval,

            string scaleType,

            int? targetUtilizationPercentage)
        {
            MaxInstances = maxInstances;
            MinInstances = minInstances;
            PollingInterval = pollingInterval;
            ScaleType = scaleType;
            TargetUtilizationPercentage = targetUtilizationPercentage;
        }
    }
}
