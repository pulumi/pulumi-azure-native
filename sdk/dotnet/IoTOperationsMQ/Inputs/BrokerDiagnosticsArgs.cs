// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Inputs
{

    /// <summary>
    /// Diagnostics setting specific to Broker
    /// </summary>
    public sealed class BrokerDiagnosticsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Diagnostic Service endpoint
        /// </summary>
        [Input("diagnosticServiceEndpoint")]
        public Input<string>? DiagnosticServiceEndpoint { get; set; }

        /// <summary>
        /// Knob to enable/disable metrics. Default = true
        /// </summary>
        [Input("enableMetrics")]
        public Input<bool>? EnableMetrics { get; set; }

        /// <summary>
        /// Enable self check on Broker via Probe.
        /// </summary>
        [Input("enableSelfCheck")]
        public Input<bool>? EnableSelfCheck { get; set; }

        /// <summary>
        /// Enable self tracing on the Broker so that every selfCheckFrequencySeconds a random message is traced even if it didn't have trace context.
        /// </summary>
        [Input("enableSelfTracing")]
        public Input<bool>? EnableSelfTracing { get; set; }

        /// <summary>
        /// Knob to enable/disable entire tracing infrastructure.
        /// </summary>
        [Input("enableTracing")]
        public Input<bool>? EnableTracing { get; set; }

        /// <summary>
        /// Format for the logs generated.
        /// </summary>
        [Input("logFormat")]
        public Input<string>? LogFormat { get; set; }

        /// <summary>
        /// Log level for the Broker.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// Maximum time for the CellMap to live.
        /// </summary>
        [Input("maxCellMapLifetime")]
        public Input<double>? MaxCellMapLifetime { get; set; }

        /// <summary>
        /// Metric update frequency in seconds.
        /// </summary>
        [Input("metricUpdateFrequencySeconds")]
        public Input<double>? MetricUpdateFrequencySeconds { get; set; }

        /// <summary>
        /// Probe Image to run.
        /// </summary>
        [Input("probeImage")]
        public Input<string>? ProbeImage { get; set; }

        /// <summary>
        /// Frequency for the self check to run.
        /// </summary>
        [Input("selfCheckFrequencySeconds")]
        public Input<double>? SelfCheckFrequencySeconds { get; set; }

        /// <summary>
        /// Time out period of the self check.
        /// </summary>
        [Input("selfCheckTimeoutSeconds")]
        public Input<double>? SelfCheckTimeoutSeconds { get; set; }

        /// <summary>
        /// The frequency at which selfTrace should run.
        /// </summary>
        [Input("selfTraceFrequencySeconds")]
        public Input<double>? SelfTraceFrequencySeconds { get; set; }

        /// <summary>
        /// The number of the spans generated by the Tracing.
        /// </summary>
        [Input("spanChannelCapacity")]
        public Input<double>? SpanChannelCapacity { get; set; }

        public BrokerDiagnosticsArgs()
        {
            EnableMetrics = true;
            EnableSelfCheck = true;
            EnableSelfTracing = true;
            EnableTracing = true;
            LogFormat = "text";
            LogLevel = "info,hyper=off,kube_client=off,tower=off,conhash=off,h2=off";
            MaxCellMapLifetime = 60;
            MetricUpdateFrequencySeconds = 30;
            ProbeImage = "sample.azurecr.io/diagnostics-probe:0.5.0";
            SelfCheckFrequencySeconds = 30;
            SelfCheckTimeoutSeconds = 15;
            SelfTraceFrequencySeconds = 30;
            SpanChannelCapacity = 1000;
        }
        public static new BrokerDiagnosticsArgs Empty => new BrokerDiagnosticsArgs();
    }
}
