// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// Observability profile to enable advanced network metrics and flow logs with historical contexts.
    /// </summary>
    public sealed class AdvancedNetworkingObservabilityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the enablement of Advanced Networking observability functionalities on clusters.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public AdvancedNetworkingObservabilityArgs()
        {
        }
        public static new AdvancedNetworkingObservabilityArgs Empty => new AdvancedNetworkingObservabilityArgs();
    }
}
