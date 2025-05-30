// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Outputs
{

    /// <summary>
    /// This is the Autoscale profile for the cluster. This will allow customer to create cluster enabled with Autoscale.
    /// </summary>
    [OutputType]
    public sealed class AutoscaleProfileResponse
    {
        /// <summary>
        /// User to specify which type of Autoscale to be implemented - Scheduled Based or Load Based.
        /// </summary>
        public readonly string? AutoscaleType;
        /// <summary>
        /// This indicates whether auto scale is enabled on HDInsight on AKS cluster.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// This property is for graceful decommission timeout; It has a default setting of 3600 seconds before forced shutdown takes place. This is the maximal time to wait for running containers and applications to complete before transition a DECOMMISSIONING node into DECOMMISSIONED. The default value is 3600 seconds. Negative value (like -1) is handled as infinite timeout.
        /// </summary>
        public readonly int? GracefulDecommissionTimeout;
        /// <summary>
        /// Profiles of load based Autoscale.
        /// </summary>
        public readonly Outputs.LoadBasedConfigResponse? LoadBasedConfig;
        /// <summary>
        /// Profiles of schedule based Autoscale.
        /// </summary>
        public readonly Outputs.ScheduleBasedConfigResponse? ScheduleBasedConfig;

        [OutputConstructor]
        private AutoscaleProfileResponse(
            string? autoscaleType,

            bool enabled,

            int? gracefulDecommissionTimeout,

            Outputs.LoadBasedConfigResponse? loadBasedConfig,

            Outputs.ScheduleBasedConfigResponse? scheduleBasedConfig)
        {
            AutoscaleType = autoscaleType;
            Enabled = enabled;
            GracefulDecommissionTimeout = gracefulDecommissionTimeout;
            LoadBasedConfig = loadBasedConfig;
            ScheduleBasedConfig = scheduleBasedConfig;
        }
    }
}
