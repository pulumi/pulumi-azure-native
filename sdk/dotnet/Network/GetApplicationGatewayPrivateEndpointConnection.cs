// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetApplicationGatewayPrivateEndpointConnection
    {
        /// <summary>
        /// Gets the specified private endpoint connection on application gateway.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Task<GetApplicationGatewayPrivateEndpointConnectionResult> InvokeAsync(GetApplicationGatewayPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationGatewayPrivateEndpointConnectionResult>("azure-native:network:getApplicationGatewayPrivateEndpointConnection", args ?? new GetApplicationGatewayPrivateEndpointConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified private endpoint connection on application gateway.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetApplicationGatewayPrivateEndpointConnectionResult> Invoke(GetApplicationGatewayPrivateEndpointConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationGatewayPrivateEndpointConnectionResult>("azure-native:network:getApplicationGatewayPrivateEndpointConnection", args ?? new GetApplicationGatewayPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified private endpoint connection on application gateway.
        /// 
        /// Uses Azure REST API version 2023-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
        /// </summary>
        public static Output<GetApplicationGatewayPrivateEndpointConnectionResult> Invoke(GetApplicationGatewayPrivateEndpointConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationGatewayPrivateEndpointConnectionResult>("azure-native:network:getApplicationGatewayPrivateEndpointConnection", args ?? new GetApplicationGatewayPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationGatewayPrivateEndpointConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application gateway.
        /// </summary>
        [Input("applicationGatewayName", required: true)]
        public string ApplicationGatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the application gateway private endpoint connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationGatewayPrivateEndpointConnectionArgs()
        {
        }
        public static new GetApplicationGatewayPrivateEndpointConnectionArgs Empty => new GetApplicationGatewayPrivateEndpointConnectionArgs();
    }

    public sealed class GetApplicationGatewayPrivateEndpointConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application gateway.
        /// </summary>
        [Input("applicationGatewayName", required: true)]
        public Input<string> ApplicationGatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the application gateway private endpoint connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApplicationGatewayPrivateEndpointConnectionInvokeArgs()
        {
        }
        public static new GetApplicationGatewayPrivateEndpointConnectionInvokeArgs Empty => new GetApplicationGatewayPrivateEndpointConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationGatewayPrivateEndpointConnectionResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The consumer link id.
        /// </summary>
        public readonly string LinkIdentifier;
        /// <summary>
        /// Name of the private endpoint connection on an application gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource of private end point.
        /// </summary>
        public readonly Outputs.PrivateEndpointResponse PrivateEndpoint;
        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// The provisioning state of the application gateway private endpoint connection resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationGatewayPrivateEndpointConnectionResult(
            string etag,

            string? id,

            string linkIdentifier,

            string? name,

            Outputs.PrivateEndpointResponse privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponse? privateLinkServiceConnectionState,

            string provisioningState,

            string type)
        {
            Etag = etag;
            Id = id;
            LinkIdentifier = linkIdentifier;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
