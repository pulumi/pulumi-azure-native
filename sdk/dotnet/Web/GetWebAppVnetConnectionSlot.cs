// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class GetWebAppVnetConnectionSlot
    {
        /// <summary>
        /// Description for Gets a virtual network the app (or deployment slot) is connected to by name.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebAppVnetConnectionSlotResult> InvokeAsync(GetWebAppVnetConnectionSlotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppVnetConnectionSlotResult>("azure-native:web:getWebAppVnetConnectionSlot", args ?? new GetWebAppVnetConnectionSlotArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets a virtual network the app (or deployment slot) is connected to by name.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppVnetConnectionSlotResult> Invoke(GetWebAppVnetConnectionSlotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppVnetConnectionSlotResult>("azure-native:web:getWebAppVnetConnectionSlot", args ?? new GetWebAppVnetConnectionSlotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Gets a virtual network the app (or deployment slot) is connected to by name.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppVnetConnectionSlotResult> Invoke(GetWebAppVnetConnectionSlotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppVnetConnectionSlotResult>("azure-native:web:getWebAppVnetConnectionSlot", args ?? new GetWebAppVnetConnectionSlotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppVnetConnectionSlotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will get the named virtual network for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        /// <summary>
        /// Name of the virtual network.
        /// </summary>
        [Input("vnetName", required: true)]
        public string VnetName { get; set; } = null!;

        public GetWebAppVnetConnectionSlotArgs()
        {
        }
        public static new GetWebAppVnetConnectionSlotArgs Empty => new GetWebAppVnetConnectionSlotArgs();
    }

    public sealed class GetWebAppVnetConnectionSlotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will get the named virtual network for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        /// <summary>
        /// Name of the virtual network.
        /// </summary>
        [Input("vnetName", required: true)]
        public Input<string> VnetName { get; set; } = null!;

        public GetWebAppVnetConnectionSlotInvokeArgs()
        {
        }
        public static new GetWebAppVnetConnectionSlotInvokeArgs Empty => new GetWebAppVnetConnectionSlotInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppVnetConnectionSlotResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
        /// Point-To-Site VPN connection.
        /// </summary>
        public readonly string? CertBlob;
        /// <summary>
        /// The client certificate thumbprint.
        /// </summary>
        public readonly string CertThumbprint;
        /// <summary>
        /// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
        /// </summary>
        public readonly string? DnsServers;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Flag that is used to denote if this is VNET injection
        /// </summary>
        public readonly bool? IsSwift;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if a resync is required; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool ResyncRequired;
        /// <summary>
        /// The routes that this Virtual Network connection uses.
        /// </summary>
        public readonly ImmutableArray<Outputs.VnetRouteResponse> Routes;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The Virtual Network's resource ID.
        /// </summary>
        public readonly string? VnetResourceId;

        [OutputConstructor]
        private GetWebAppVnetConnectionSlotResult(
            string azureApiVersion,

            string? certBlob,

            string certThumbprint,

            string? dnsServers,

            string id,

            bool? isSwift,

            string? kind,

            string name,

            bool resyncRequired,

            ImmutableArray<Outputs.VnetRouteResponse> routes,

            string type,

            string? vnetResourceId)
        {
            AzureApiVersion = azureApiVersion;
            CertBlob = certBlob;
            CertThumbprint = certThumbprint;
            DnsServers = dnsServers;
            Id = id;
            IsSwift = isSwift;
            Kind = kind;
            Name = name;
            ResyncRequired = resyncRequired;
            Routes = routes;
            Type = type;
            VnetResourceId = vnetResourceId;
        }
    }
}
