// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.V20210601
{
    public static class GetPrivateLinkHub
    {
        /// <summary>
        /// Gets a privateLinkHub
        /// </summary>
        public static Task<GetPrivateLinkHubResult> InvokeAsync(GetPrivateLinkHubArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateLinkHubResult>("azure-native:synapse/v20210601:getPrivateLinkHub", args ?? new GetPrivateLinkHubArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a privateLinkHub
        /// </summary>
        public static Output<GetPrivateLinkHubResult> Invoke(GetPrivateLinkHubInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateLinkHubResult>("azure-native:synapse/v20210601:getPrivateLinkHub", args ?? new GetPrivateLinkHubInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a privateLinkHub
        /// </summary>
        public static Output<GetPrivateLinkHubResult> Invoke(GetPrivateLinkHubInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateLinkHubResult>("azure-native:synapse/v20210601:getPrivateLinkHub", args ?? new GetPrivateLinkHubInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateLinkHubArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the privateLinkHub
        /// </summary>
        [Input("privateLinkHubName", required: true)]
        public string PrivateLinkHubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateLinkHubArgs()
        {
        }
        public static new GetPrivateLinkHubArgs Empty => new GetPrivateLinkHubArgs();
    }

    public sealed class GetPrivateLinkHubInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the privateLinkHub
        /// </summary>
        [Input("privateLinkHubName", required: true)]
        public Input<string> PrivateLinkHubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPrivateLinkHubInvokeArgs()
        {
        }
        public static new GetPrivateLinkHubInvokeArgs Empty => new GetPrivateLinkHubInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateLinkHubResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of private endpoint connections
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionForPrivateLinkHubBasicResponse> PrivateEndpointConnections;
        /// <summary>
        /// PrivateLinkHub provisioning state
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateLinkHubResult(
            string id,

            string location,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionForPrivateLinkHubBasicResponse> privateEndpointConnections,

            string? provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
