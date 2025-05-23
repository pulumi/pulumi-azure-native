// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// Azure Monitor Workspace Logs specific configurations.
    /// </summary>
    public sealed class AzureMonitorWorkspaceLogsExporterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API configurations for Azure Monitor workspace exporter.
        /// </summary>
        [Input("api", required: true)]
        public Input<Inputs.AzureMonitorWorkspaceLogsApiConfigArgs> Api { get; set; } = null!;

        /// <summary>
        /// Cache configurations.
        /// </summary>
        [Input("cache")]
        public Input<Inputs.CacheConfigurationArgs>? Cache { get; set; }

        /// <summary>
        /// Concurrency configuration for the exporter.
        /// </summary>
        [Input("concurrency")]
        public Input<Inputs.ConcurrencyConfigurationArgs>? Concurrency { get; set; }

        public AzureMonitorWorkspaceLogsExporterArgs()
        {
        }
        public static new AzureMonitorWorkspaceLogsExporterArgs Empty => new AzureMonitorWorkspaceLogsExporterArgs();
    }
}
