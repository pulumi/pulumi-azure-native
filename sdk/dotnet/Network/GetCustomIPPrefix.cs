// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetCustomIPPrefix
    {
        /// <summary>
        /// Gets the specified custom IP prefix in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetCustomIPPrefixResult> InvokeAsync(GetCustomIPPrefixArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomIPPrefixResult>("azure-native:network:getCustomIPPrefix", args ?? new GetCustomIPPrefixArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified custom IP prefix in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCustomIPPrefixResult> Invoke(GetCustomIPPrefixInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomIPPrefixResult>("azure-native:network:getCustomIPPrefix", args ?? new GetCustomIPPrefixInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified custom IP prefix in a specified resource group.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCustomIPPrefixResult> Invoke(GetCustomIPPrefixInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomIPPrefixResult>("azure-native:network:getCustomIPPrefix", args ?? new GetCustomIPPrefixInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomIPPrefixArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom IP prefix.
        /// </summary>
        [Input("customIpPrefixName", required: true)]
        public string CustomIpPrefixName { get; set; } = null!;

        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCustomIPPrefixArgs()
        {
        }
        public static new GetCustomIPPrefixArgs Empty => new GetCustomIPPrefixArgs();
    }

    public sealed class GetCustomIPPrefixInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom IP prefix.
        /// </summary>
        [Input("customIpPrefixName", required: true)]
        public Input<string> CustomIpPrefixName { get; set; } = null!;

        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetCustomIPPrefixInvokeArgs()
        {
        }
        public static new GetCustomIPPrefixInvokeArgs Empty => new GetCustomIPPrefixInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomIPPrefixResult
    {
        /// <summary>
        /// The ASN for CIDR advertising. Should be an integer as string.
        /// </summary>
        public readonly string? Asn;
        /// <summary>
        /// Authorization message for WAN validation.
        /// </summary>
        public readonly string? AuthorizationMessage;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The list of all Children for IPv6 /48 CustomIpPrefix.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> ChildCustomIpPrefixes;
        /// <summary>
        /// The prefix range in CIDR notation. Should include the start address and the prefix length.
        /// </summary>
        public readonly string? Cidr;
        /// <summary>
        /// The commissioned state of the Custom IP Prefix.
        /// </summary>
        public readonly string? CommissionedState;
        /// <summary>
        /// The Parent CustomIpPrefix for IPv6 /64 CustomIpPrefix.
        /// </summary>
        public readonly Outputs.SubResourceResponse? CustomIpPrefixParent;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Whether to do express route advertise.
        /// </summary>
        public readonly bool? ExpressRouteAdvertise;
        /// <summary>
        /// The extended location of the custom IP prefix.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// The reason why resource is in failed state.
        /// </summary>
        public readonly string FailedReason;
        /// <summary>
        /// The Geo for CIDR advertising. Should be an Geo code.
        /// </summary>
        public readonly string? Geo;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether to Advertise the range to Internet.
        /// </summary>
        public readonly bool? NoInternetAdvertise;
        /// <summary>
        /// Type of custom IP prefix. Should be Singular, Parent, or Child.
        /// </summary>
        public readonly string? PrefixType;
        /// <summary>
        /// The provisioning state of the custom IP prefix resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The list of all referenced PublicIpPrefixes.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> PublicIpPrefixes;
        /// <summary>
        /// The resource GUID property of the custom IP prefix resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Signed message for WAN validation.
        /// </summary>
        public readonly string? SignedMessage;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A list of availability zones denoting the IP allocated for the resource needs to come from.
        /// </summary>
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private GetCustomIPPrefixResult(
            string? asn,

            string? authorizationMessage,

            string azureApiVersion,

            ImmutableArray<Outputs.SubResourceResponse> childCustomIpPrefixes,

            string? cidr,

            string? commissionedState,

            Outputs.SubResourceResponse? customIpPrefixParent,

            string etag,

            bool? expressRouteAdvertise,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string failedReason,

            string? geo,

            string? id,

            string? location,

            string name,

            bool? noInternetAdvertise,

            string? prefixType,

            string provisioningState,

            ImmutableArray<Outputs.SubResourceResponse> publicIpPrefixes,

            string resourceGuid,

            string? signedMessage,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<string> zones)
        {
            Asn = asn;
            AuthorizationMessage = authorizationMessage;
            AzureApiVersion = azureApiVersion;
            ChildCustomIpPrefixes = childCustomIpPrefixes;
            Cidr = cidr;
            CommissionedState = commissionedState;
            CustomIpPrefixParent = customIpPrefixParent;
            Etag = etag;
            ExpressRouteAdvertise = expressRouteAdvertise;
            ExtendedLocation = extendedLocation;
            FailedReason = failedReason;
            Geo = geo;
            Id = id;
            Location = location;
            Name = name;
            NoInternetAdvertise = noInternetAdvertise;
            PrefixType = prefixType;
            ProvisioningState = provisioningState;
            PublicIpPrefixes = publicIpPrefixes;
            ResourceGuid = resourceGuid;
            SignedMessage = signedMessage;
            Tags = tags;
            Type = type;
            Zones = zones;
        }
    }
}
