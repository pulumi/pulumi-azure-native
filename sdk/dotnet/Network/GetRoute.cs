// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetRoute
    {
        /// <summary>
        /// Gets the specified route from a route table.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("azure-native:network:getRoute", args ?? new GetRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified route from a route table.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteResult>("azure-native:network:getRoute", args ?? new GetRouteInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified route from a route table.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteResult>("azure-native:network:getRoute", args ?? new GetRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route.
        /// </summary>
        [Input("routeName", required: true)]
        public string RouteName { get; set; } = null!;

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public string RouteTableName { get; set; } = null!;

        public GetRouteArgs()
        {
        }
        public static new GetRouteArgs Empty => new GetRouteArgs();
    }

    public sealed class GetRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route.
        /// </summary>
        [Input("routeName", required: true)]
        public Input<string> RouteName { get; set; } = null!;

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        public GetRouteInvokeArgs()
        {
        }
        public static new GetRouteInvokeArgs Empty => new GetRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteResult
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        public readonly string? AddressPrefix;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
        /// </summary>
        public readonly bool HasBgpOverride;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        public readonly string? NextHopIpAddress;
        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        public readonly string NextHopType;
        /// <summary>
        /// The provisioning state of the route resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetRouteResult(
            string? addressPrefix,

            string azureApiVersion,

            string etag,

            bool hasBgpOverride,

            string? id,

            string? name,

            string? nextHopIpAddress,

            string nextHopType,

            string provisioningState,

            string? type)
        {
            AddressPrefix = addressPrefix;
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            HasBgpOverride = hasBgpOverride;
            Id = id;
            Name = name;
            NextHopIpAddress = nextHopIpAddress;
            NextHopType = nextHopType;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
