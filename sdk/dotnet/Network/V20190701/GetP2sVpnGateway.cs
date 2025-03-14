// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20190701
{
    public static class GetP2sVpnGateway
    {
        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// </summary>
        public static Task<GetP2sVpnGatewayResult> InvokeAsync(GetP2sVpnGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetP2sVpnGatewayResult>("azure-native:network/v20190701:getP2sVpnGateway", args ?? new GetP2sVpnGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// </summary>
        public static Output<GetP2sVpnGatewayResult> Invoke(GetP2sVpnGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayResult>("azure-native:network/v20190701:getP2sVpnGateway", args ?? new GetP2sVpnGatewayInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// </summary>
        public static Output<GetP2sVpnGatewayResult> Invoke(GetP2sVpnGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayResult>("azure-native:network/v20190701:getP2sVpnGateway", args ?? new GetP2sVpnGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetP2sVpnGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public string GatewayName { get; set; } = null!;

        /// <summary>
        /// The resource group name of the P2SVpnGateway.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetP2sVpnGatewayArgs()
        {
        }
        public static new GetP2sVpnGatewayArgs Empty => new GetP2sVpnGatewayArgs();
    }

    public sealed class GetP2sVpnGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// The resource group name of the P2SVpnGateway.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetP2sVpnGatewayInvokeArgs()
        {
        }
        public static new GetP2sVpnGatewayInvokeArgs Empty => new GetP2sVpnGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetP2sVpnGatewayResult
    {
        /// <summary>
        /// The reference of the address space resource which represents the custom routes specified by the customer for P2SVpnGateway and P2S VpnClient.
        /// </summary>
        public readonly Outputs.AddressSpaceResponse? CustomRoutes;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The P2SVpnServerConfiguration to which the p2sVpnGateway is attached to.
        /// </summary>
        public readonly Outputs.SubResourceResponse? P2SVpnServerConfiguration;
        /// <summary>
        /// The provisioning state of the P2S VPN gateway resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The VirtualHub to which the gateway belongs.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualHub;
        /// <summary>
        /// The reference of the address space resource which represents Address space for P2S VpnClient.
        /// </summary>
        public readonly Outputs.AddressSpaceResponse? VpnClientAddressPool;
        /// <summary>
        /// All P2S VPN clients' connection health status.
        /// </summary>
        public readonly Outputs.VpnClientConnectionHealthResponse VpnClientConnectionHealth;
        /// <summary>
        /// The scale unit for this p2s vpn gateway.
        /// </summary>
        public readonly int? VpnGatewayScaleUnit;

        [OutputConstructor]
        private GetP2sVpnGatewayResult(
            Outputs.AddressSpaceResponse? customRoutes,

            string etag,

            string? id,

            string location,

            string name,

            Outputs.SubResourceResponse? p2SVpnServerConfiguration,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.SubResourceResponse? virtualHub,

            Outputs.AddressSpaceResponse? vpnClientAddressPool,

            Outputs.VpnClientConnectionHealthResponse vpnClientConnectionHealth,

            int? vpnGatewayScaleUnit)
        {
            CustomRoutes = customRoutes;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            P2SVpnServerConfiguration = p2SVpnServerConfiguration;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            VirtualHub = virtualHub;
            VpnClientAddressPool = vpnClientAddressPool;
            VpnClientConnectionHealth = vpnClientConnectionHealth;
            VpnGatewayScaleUnit = vpnGatewayScaleUnit;
        }
    }
}
