// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20230702Preview.Inputs
{

    /// <summary>
    /// This addon can be used to configure network monitoring and generate network monitoring data in Prometheus format
    /// </summary>
    public sealed class NetworkMonitoringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or disable the network monitoring plugin on the cluster
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public NetworkMonitoringArgs()
        {
        }
        public static new NetworkMonitoringArgs Empty => new NetworkMonitoringArgs();
    }
}
