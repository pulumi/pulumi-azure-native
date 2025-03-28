// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid
{
    public static class GetPrivateEndpointConnection
    {
        /// <summary>
        /// Get a specific private endpoint connection under a topic, domain, or partner namespace.
        /// 
        /// Uses Azure REST API version 2022-06-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-02-15.
        /// </summary>
        public static Task<GetPrivateEndpointConnectionResult> InvokeAsync(GetPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionResult>("azure-native:eventgrid:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Get a specific private endpoint connection under a topic, domain, or partner namespace.
        /// 
        /// Uses Azure REST API version 2022-06-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-02-15.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:eventgrid:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a specific private endpoint connection under a topic, domain, or partner namespace.
        /// 
        /// Uses Azure REST API version 2022-06-15.
        /// 
        /// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview, 2024-12-15-preview, 2025-02-15.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:eventgrid:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateEndpointConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
        /// </summary>
        [Input("parentName", required: true)]
        public string ParentName { get; set; } = null!;

        /// <summary>
        /// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
        /// </summary>
        [Input("parentType", required: true)]
        public string ParentType { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint connection connection.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public string PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateEndpointConnectionArgs()
        {
        }
        public static new GetPrivateEndpointConnectionArgs Empty => new GetPrivateEndpointConnectionArgs();
    }

    public sealed class GetPrivateEndpointConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
        /// </summary>
        [Input("parentName", required: true)]
        public Input<string> ParentName { get; set; } = null!;

        /// <summary>
        /// The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
        /// </summary>
        [Input("parentType", required: true)]
        public Input<string> ParentType { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint connection connection.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public Input<string> PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPrivateEndpointConnectionInvokeArgs()
        {
        }
        public static new GetPrivateEndpointConnectionInvokeArgs Empty => new GetPrivateEndpointConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateEndpointConnectionResult
    {
        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        public readonly ImmutableArray<string> GroupIds;
        /// <summary>
        /// Fully qualified identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        public readonly Outputs.PrivateEndpointResponse? PrivateEndpoint;
        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        public readonly Outputs.ConnectionStateResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointConnectionResult(
            ImmutableArray<string> groupIds,

            string id,

            string name,

            Outputs.PrivateEndpointResponse? privateEndpoint,

            Outputs.ConnectionStateResponse? privateLinkServiceConnectionState,

            string? provisioningState,

            string type)
        {
            GroupIds = groupIds;
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
