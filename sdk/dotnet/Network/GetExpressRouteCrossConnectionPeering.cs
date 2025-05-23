// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetExpressRouteCrossConnectionPeering
    {
        /// <summary>
        /// Gets the specified peering for the ExpressRouteCrossConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetExpressRouteCrossConnectionPeeringResult> InvokeAsync(GetExpressRouteCrossConnectionPeeringArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExpressRouteCrossConnectionPeeringResult>("azure-native:network:getExpressRouteCrossConnectionPeering", args ?? new GetExpressRouteCrossConnectionPeeringArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified peering for the ExpressRouteCrossConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExpressRouteCrossConnectionPeeringResult> Invoke(GetExpressRouteCrossConnectionPeeringInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExpressRouteCrossConnectionPeeringResult>("azure-native:network:getExpressRouteCrossConnectionPeering", args ?? new GetExpressRouteCrossConnectionPeeringInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified peering for the ExpressRouteCrossConnection.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetExpressRouteCrossConnectionPeeringResult> Invoke(GetExpressRouteCrossConnectionPeeringInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExpressRouteCrossConnectionPeeringResult>("azure-native:network:getExpressRouteCrossConnectionPeering", args ?? new GetExpressRouteCrossConnectionPeeringInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExpressRouteCrossConnectionPeeringArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ExpressRouteCrossConnection.
        /// </summary>
        [Input("crossConnectionName", required: true)]
        public string CrossConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public string PeeringName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetExpressRouteCrossConnectionPeeringArgs()
        {
        }
        public static new GetExpressRouteCrossConnectionPeeringArgs Empty => new GetExpressRouteCrossConnectionPeeringArgs();
    }

    public sealed class GetExpressRouteCrossConnectionPeeringInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ExpressRouteCrossConnection.
        /// </summary>
        [Input("crossConnectionName", required: true)]
        public Input<string> CrossConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetExpressRouteCrossConnectionPeeringInvokeArgs()
        {
        }
        public static new GetExpressRouteCrossConnectionPeeringInvokeArgs Empty => new GetExpressRouteCrossConnectionPeeringInvokeArgs();
    }


    [OutputType]
    public sealed class GetExpressRouteCrossConnectionPeeringResult
    {
        /// <summary>
        /// The Azure ASN.
        /// </summary>
        public readonly int AzureASN;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        public readonly string? GatewayManagerEtag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The IPv6 peering configuration.
        /// </summary>
        public readonly Outputs.Ipv6ExpressRouteCircuitPeeringConfigResponse? Ipv6PeeringConfig;
        /// <summary>
        /// Who was the last to modify the peering.
        /// </summary>
        public readonly string LastModifiedBy;
        /// <summary>
        /// The Microsoft peering configuration.
        /// </summary>
        public readonly Outputs.ExpressRouteCircuitPeeringConfigResponse? MicrosoftPeeringConfig;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The peer ASN.
        /// </summary>
        public readonly double? PeerASN;
        /// <summary>
        /// The peering type.
        /// </summary>
        public readonly string? PeeringType;
        /// <summary>
        /// The primary port.
        /// </summary>
        public readonly string PrimaryAzurePort;
        /// <summary>
        /// The primary address prefix.
        /// </summary>
        public readonly string? PrimaryPeerAddressPrefix;
        /// <summary>
        /// The provisioning state of the express route cross connection peering resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The secondary port.
        /// </summary>
        public readonly string SecondaryAzurePort;
        /// <summary>
        /// The secondary address prefix.
        /// </summary>
        public readonly string? SecondaryPeerAddressPrefix;
        /// <summary>
        /// The shared key.
        /// </summary>
        public readonly string? SharedKey;
        /// <summary>
        /// The peering state.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The VLAN ID.
        /// </summary>
        public readonly int? VlanId;

        [OutputConstructor]
        private GetExpressRouteCrossConnectionPeeringResult(
            int azureASN,

            string azureApiVersion,

            string etag,

            string? gatewayManagerEtag,

            string? id,

            Outputs.Ipv6ExpressRouteCircuitPeeringConfigResponse? ipv6PeeringConfig,

            string lastModifiedBy,

            Outputs.ExpressRouteCircuitPeeringConfigResponse? microsoftPeeringConfig,

            string? name,

            double? peerASN,

            string? peeringType,

            string primaryAzurePort,

            string? primaryPeerAddressPrefix,

            string provisioningState,

            string secondaryAzurePort,

            string? secondaryPeerAddressPrefix,

            string? sharedKey,

            string? state,

            int? vlanId)
        {
            AzureASN = azureASN;
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            GatewayManagerEtag = gatewayManagerEtag;
            Id = id;
            Ipv6PeeringConfig = ipv6PeeringConfig;
            LastModifiedBy = lastModifiedBy;
            MicrosoftPeeringConfig = microsoftPeeringConfig;
            Name = name;
            PeerASN = peerASN;
            PeeringType = peeringType;
            PrimaryAzurePort = primaryAzurePort;
            PrimaryPeerAddressPrefix = primaryPeerAddressPrefix;
            ProvisioningState = provisioningState;
            SecondaryAzurePort = secondaryAzurePort;
            SecondaryPeerAddressPrefix = secondaryPeerAddressPrefix;
            SharedKey = sharedKey;
            State = state;
            VlanId = vlanId;
        }
    }
}
