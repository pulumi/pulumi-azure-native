// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetP2sVpnGateway
    {
        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetP2sVpnGatewayResult> InvokeAsync(GetP2sVpnGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetP2sVpnGatewayResult>("azure-native:network:getP2sVpnGateway", args ?? new GetP2sVpnGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetP2sVpnGatewayResult> Invoke(GetP2sVpnGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayResult>("azure-native:network:getP2sVpnGateway", args ?? new GetP2sVpnGatewayInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the details of a virtual wan p2s vpn gateway.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetP2sVpnGatewayResult> Invoke(GetP2sVpnGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayResult>("azure-native:network:getP2sVpnGateway", args ?? new GetP2sVpnGatewayInvokeArgs(), options.WithDefaults());
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
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// List of all customer specified DNS servers IP addresses.
        /// </summary>
        public readonly ImmutableArray<string> CustomDnsServers;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
        /// </summary>
        public readonly bool? IsRoutingPreferenceInternet;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of all p2s connection configurations of the gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.P2SConnectionConfigurationResponse> P2SConnectionConfigurations;
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
        /// All P2S VPN clients' connection health status.
        /// </summary>
        public readonly Outputs.VpnClientConnectionHealthResponse VpnClientConnectionHealth;
        /// <summary>
        /// The scale unit for this p2s vpn gateway.
        /// </summary>
        public readonly int? VpnGatewayScaleUnit;
        /// <summary>
        /// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VpnServerConfiguration;

        [OutputConstructor]
        private GetP2sVpnGatewayResult(
            string azureApiVersion,

            ImmutableArray<string> customDnsServers,

            string etag,

            string? id,

            bool? isRoutingPreferenceInternet,

            string location,

            string name,

            ImmutableArray<Outputs.P2SConnectionConfigurationResponse> p2SConnectionConfigurations,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.SubResourceResponse? virtualHub,

            Outputs.VpnClientConnectionHealthResponse vpnClientConnectionHealth,

            int? vpnGatewayScaleUnit,

            Outputs.SubResourceResponse? vpnServerConfiguration)
        {
            AzureApiVersion = azureApiVersion;
            CustomDnsServers = customDnsServers;
            Etag = etag;
            Id = id;
            IsRoutingPreferenceInternet = isRoutingPreferenceInternet;
            Location = location;
            Name = name;
            P2SConnectionConfigurations = p2SConnectionConfigurations;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            VirtualHub = virtualHub;
            VpnClientConnectionHealth = vpnClientConnectionHealth;
            VpnGatewayScaleUnit = vpnGatewayScaleUnit;
            VpnServerConfiguration = vpnServerConfiguration;
        }
    }
}
