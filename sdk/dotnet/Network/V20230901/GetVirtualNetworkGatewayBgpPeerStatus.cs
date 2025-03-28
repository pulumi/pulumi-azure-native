// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20230901
{
    public static class GetVirtualNetworkGatewayBgpPeerStatus
    {
        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// </summary>
        public static Task<GetVirtualNetworkGatewayBgpPeerStatusResult> InvokeAsync(GetVirtualNetworkGatewayBgpPeerStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network/v20230901:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusArgs(), options.WithDefaults());

        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// </summary>
        public static Output<GetVirtualNetworkGatewayBgpPeerStatusResult> Invoke(GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network/v20230901:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// </summary>
        public static Output<GetVirtualNetworkGatewayBgpPeerStatusResult> Invoke(GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network/v20230901:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkGatewayBgpPeerStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address of the peer to retrieve the status of.
        /// </summary>
        [Input("peer")]
        public string? Peer { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network gateway.
        /// </summary>
        [Input("virtualNetworkGatewayName", required: true)]
        public string VirtualNetworkGatewayName { get; set; } = null!;

        public GetVirtualNetworkGatewayBgpPeerStatusArgs()
        {
        }
        public static new GetVirtualNetworkGatewayBgpPeerStatusArgs Empty => new GetVirtualNetworkGatewayBgpPeerStatusArgs();
    }

    public sealed class GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address of the peer to retrieve the status of.
        /// </summary>
        [Input("peer")]
        public Input<string>? Peer { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network gateway.
        /// </summary>
        [Input("virtualNetworkGatewayName", required: true)]
        public Input<string> VirtualNetworkGatewayName { get; set; } = null!;

        public GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs()
        {
        }
        public static new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs Empty => new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualNetworkGatewayBgpPeerStatusResult
    {
        /// <summary>
        /// List of BGP peers.
        /// </summary>
        public readonly ImmutableArray<Outputs.BgpPeerStatusResponse> Value;

        [OutputConstructor]
        private GetVirtualNetworkGatewayBgpPeerStatusResult(ImmutableArray<Outputs.BgpPeerStatusResponse> value)
        {
            Value = value;
        }
    }
}
