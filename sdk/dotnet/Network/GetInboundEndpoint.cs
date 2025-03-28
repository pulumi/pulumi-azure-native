// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetInboundEndpoint
    {
        /// <summary>
        /// Gets properties of an inbound endpoint for a DNS resolver.
        /// 
        /// Uses Azure REST API version 2022-07-01.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2023-07-01-preview.
        /// </summary>
        public static Task<GetInboundEndpointResult> InvokeAsync(GetInboundEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInboundEndpointResult>("azure-native:network:getInboundEndpoint", args ?? new GetInboundEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets properties of an inbound endpoint for a DNS resolver.
        /// 
        /// Uses Azure REST API version 2022-07-01.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2023-07-01-preview.
        /// </summary>
        public static Output<GetInboundEndpointResult> Invoke(GetInboundEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInboundEndpointResult>("azure-native:network:getInboundEndpoint", args ?? new GetInboundEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets properties of an inbound endpoint for a DNS resolver.
        /// 
        /// Uses Azure REST API version 2022-07-01.
        /// 
        /// Other available API versions: 2020-04-01-preview, 2023-07-01-preview.
        /// </summary>
        public static Output<GetInboundEndpointResult> Invoke(GetInboundEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInboundEndpointResult>("azure-native:network:getInboundEndpoint", args ?? new GetInboundEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInboundEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNS resolver.
        /// </summary>
        [Input("dnsResolverName", required: true)]
        public string DnsResolverName { get; set; } = null!;

        /// <summary>
        /// The name of the inbound endpoint for the DNS resolver.
        /// </summary>
        [Input("inboundEndpointName", required: true)]
        public string InboundEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetInboundEndpointArgs()
        {
        }
        public static new GetInboundEndpointArgs Empty => new GetInboundEndpointArgs();
    }

    public sealed class GetInboundEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNS resolver.
        /// </summary>
        [Input("dnsResolverName", required: true)]
        public Input<string> DnsResolverName { get; set; } = null!;

        /// <summary>
        /// The name of the inbound endpoint for the DNS resolver.
        /// </summary>
        [Input("inboundEndpointName", required: true)]
        public Input<string> InboundEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetInboundEndpointInvokeArgs()
        {
        }
        public static new GetInboundEndpointInvokeArgs Empty => new GetInboundEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetInboundEndpointResult
    {
        /// <summary>
        /// ETag of the inbound endpoint.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP configurations for the inbound endpoint.
        /// </summary>
        public readonly ImmutableArray<Outputs.InboundEndpointIPConfigurationResponse> IpConfigurations;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resourceGuid property of the inbound endpoint resource.
        /// </summary>
        public readonly string ResourceGuid;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInboundEndpointResult(
            string etag,

            string id,

            ImmutableArray<Outputs.InboundEndpointIPConfigurationResponse> ipConfigurations,

            string location,

            string name,

            string provisioningState,

            string resourceGuid,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Id = id;
            IpConfigurations = ipConfigurations;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
