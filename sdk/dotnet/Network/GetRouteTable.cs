// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetRouteTable
    {
        /// <summary>
        /// Gets the specified route table.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2019-06-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("azure-native:network:getRouteTable", args ?? new GetRouteTableArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified route table.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2019-06-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetRouteTableResult> Invoke(GetRouteTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteTableResult>("azure-native:network:getRouteTable", args ?? new GetRouteTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified route table.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2019-06-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetRouteTableResult> Invoke(GetRouteTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteTableResult>("azure-native:network:getRouteTable", args ?? new GetRouteTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteTableArgs : global::Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public string RouteTableName { get; set; } = null!;

        public GetRouteTableArgs()
        {
        }
        public static new GetRouteTableArgs Empty => new GetRouteTableArgs();
    }

    public sealed class GetRouteTableInvokeArgs : global::Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        public GetRouteTableInvokeArgs()
        {
        }
        public static new GetRouteTableInvokeArgs Empty => new GetRouteTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        /// <summary>
        /// Whether to disable the routes learned by BGP on that route table. True means disable.
        /// </summary>
        public readonly bool? DisableBgpRoutePropagation;
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
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the route table resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource GUID property of the route table.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Collection of routes contained within a route table.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteResponse> Routes;
        /// <summary>
        /// A collection of references to subnets.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetResponse> Subnets;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRouteTableResult(
            bool? disableBgpRoutePropagation,

            string etag,

            string? id,

            string? location,

            string name,

            string provisioningState,

            string resourceGuid,

            ImmutableArray<Outputs.RouteResponse> routes,

            ImmutableArray<Outputs.SubnetResponse> subnets,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DisableBgpRoutePropagation = disableBgpRoutePropagation;
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            Routes = routes;
            Subnets = subnets;
            Tags = tags;
            Type = type;
        }
    }
}
