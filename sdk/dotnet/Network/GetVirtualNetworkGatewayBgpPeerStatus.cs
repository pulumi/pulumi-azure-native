// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetVirtualNetworkGatewayBgpPeerStatus
    {
        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVirtualNetworkGatewayBgpPeerStatusResult> InvokeAsync(GetVirtualNetworkGatewayBgpPeerStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusArgs(), options.WithDefaults());

        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVirtualNetworkGatewayBgpPeerStatusResult> Invoke(GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The GetBgpPeerStatus operation retrieves the status of all BGP peers.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVirtualNetworkGatewayBgpPeerStatusResult> Invoke(GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayBgpPeerStatusResult>("azure-native:network:getVirtualNetworkGatewayBgpPeerStatus", args ?? new GetVirtualNetworkGatewayBgpPeerStatusInvokeArgs(), options.WithDefaults());
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
