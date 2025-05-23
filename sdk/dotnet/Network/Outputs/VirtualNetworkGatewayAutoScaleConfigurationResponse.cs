// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Virtual Network Gateway Autoscale Configuration details
    /// </summary>
    [OutputType]
    public sealed class VirtualNetworkGatewayAutoScaleConfigurationResponse
    {
        /// <summary>
        /// The bounds of the autoscale configuration
        /// </summary>
        public readonly Outputs.VirtualNetworkGatewayAutoScaleBoundsResponse? Bounds;

        [OutputConstructor]
        private VirtualNetworkGatewayAutoScaleConfigurationResponse(Outputs.VirtualNetworkGatewayAutoScaleBoundsResponse? bounds)
        {
            Bounds = bounds;
        }
    }
}
