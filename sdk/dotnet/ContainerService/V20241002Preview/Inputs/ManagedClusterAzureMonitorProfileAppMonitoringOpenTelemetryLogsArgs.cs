// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20241002Preview.Inputs
{

    /// <summary>
    /// Application Monitoring Open Telemetry Metrics Profile for Kubernetes Application Container Logs and Traces. Collects OpenTelemetry logs and traces of the application using Azure Monitor OpenTelemetry based SDKs. See aka.ms/AzureMonitorApplicationMonitoring for an overview.
    /// </summary>
    public sealed class ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if Application Monitoring Open Telemetry Logs and traces is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Open Telemetry host port for Open Telemetry logs and traces. If not specified, the default port is 28331.
        /// </summary>
        [Input("port")]
        public Input<double>? Port { get; set; }

        public ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogsArgs()
        {
        }
        public static new ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogsArgs Empty => new ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryLogsArgs();
    }
}
