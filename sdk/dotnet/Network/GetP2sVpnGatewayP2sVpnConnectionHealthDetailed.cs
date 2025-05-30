// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetP2sVpnGatewayP2sVpnConnectionHealthDetailed
    {
        /// <summary>
        /// Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult> InvokeAsync(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult>("azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed", args ?? new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult> Invoke(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult>("azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed", args ?? new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult> Invoke(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult>("azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed", args ?? new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the P2SVpnGateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public string GatewayName { get; set; } = null!;

        /// <summary>
        /// The sas-url to download the P2S Vpn connection health detail.
        /// </summary>
        [Input("outputBlobSasUrl")]
        public string? OutputBlobSasUrl { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("vpnUserNamesFilter")]
        private List<string>? _vpnUserNamesFilter;

        /// <summary>
        /// The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
        /// </summary>
        public List<string> VpnUserNamesFilter
        {
            get => _vpnUserNamesFilter ?? (_vpnUserNamesFilter = new List<string>());
            set => _vpnUserNamesFilter = value;
        }

        public GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs()
        {
        }
        public static new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs Empty => new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs();
    }

    public sealed class GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the P2SVpnGateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// The sas-url to download the P2S Vpn connection health detail.
        /// </summary>
        [Input("outputBlobSasUrl")]
        public Input<string>? OutputBlobSasUrl { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("vpnUserNamesFilter")]
        private InputList<string>? _vpnUserNamesFilter;

        /// <summary>
        /// The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
        /// </summary>
        public InputList<string> VpnUserNamesFilter
        {
            get => _vpnUserNamesFilter ?? (_vpnUserNamesFilter = new InputList<string>());
            set => _vpnUserNamesFilter = value;
        }

        public GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs()
        {
        }
        public static new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs Empty => new GetP2sVpnGatewayP2sVpnConnectionHealthDetailedInvokeArgs();
    }


    [OutputType]
    public sealed class GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult
    {
        /// <summary>
        /// Returned sas url of the blob to which the p2s vpn connection detailed health will be written.
        /// </summary>
        public readonly string? SasUrl;

        [OutputConstructor]
        private GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult(string? sasUrl)
        {
            SasUrl = sasUrl;
        }
    }
}
