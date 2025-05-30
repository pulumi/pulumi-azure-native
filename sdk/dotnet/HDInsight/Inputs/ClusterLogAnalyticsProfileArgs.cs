// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// Cluster log analytics profile to enable or disable OMS agent for cluster.
    /// </summary>
    public sealed class ClusterLogAnalyticsProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Collection of logs to be enabled or disabled for log analytics.
        /// </summary>
        [Input("applicationLogs")]
        public Input<Inputs.ClusterLogAnalyticsApplicationLogsArgs>? ApplicationLogs { get; set; }

        /// <summary>
        /// True if log analytics is enabled for the cluster, otherwise false.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// True if metrics are enabled, otherwise false.
        /// </summary>
        [Input("metricsEnabled")]
        public Input<bool>? MetricsEnabled { get; set; }

        public ClusterLogAnalyticsProfileArgs()
        {
        }
        public static new ClusterLogAnalyticsProfileArgs Empty => new ClusterLogAnalyticsProfileArgs();
    }
}
