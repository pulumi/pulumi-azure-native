// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Dashboard
{
    public static class GetPrivateEndpointConnection
    {
        /// <summary>
        /// The Private Endpoint Connection resource.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-10-01-preview, 2023-09-01, 2023-10-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetPrivateEndpointConnectionResult> InvokeAsync(GetPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionResult>("azure-native:dashboard:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// The Private Endpoint Connection resource.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-10-01-preview, 2023-09-01, 2023-10-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:dashboard:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The Private Endpoint Connection resource.
        /// 
        /// Uses Azure REST API version 2024-10-01.
        /// 
        /// Other available API versions: 2022-08-01, 2022-10-01-preview, 2023-09-01, 2023-10-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:dashboard:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateEndpointConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private endpoint connection name of Azure Managed Grafana.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public string PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The workspace name of Azure Managed Grafana.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetPrivateEndpointConnectionArgs()
        {
        }
        public static new GetPrivateEndpointConnectionArgs Empty => new GetPrivateEndpointConnectionArgs();
    }

    public sealed class GetPrivateEndpointConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private endpoint connection name of Azure Managed Grafana.
        /// </summary>
        [Input("privateEndpointConnectionName", required: true)]
        public Input<string> PrivateEndpointConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The workspace name of Azure Managed Grafana.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetPrivateEndpointConnectionInvokeArgs()
        {
        }
        public static new GetPrivateEndpointConnectionInvokeArgs Empty => new GetPrivateEndpointConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateEndpointConnectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The private endpoint connection group ids.
        /// </summary>
        public readonly ImmutableArray<string> GroupIds;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource of private end point.
        /// </summary>
        public readonly Outputs.PrivateEndpointResponse? PrivateEndpoint;
        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStateResponse PrivateLinkServiceConnectionState;
        /// <summary>
        /// The provisioning state of the private endpoint connection resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointConnectionResult(
            string azureApiVersion,

            ImmutableArray<string> groupIds,

            string id,

            string name,

            Outputs.PrivateEndpointResponse? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponse privateLinkServiceConnectionState,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            GroupIds = groupIds;
            Id = id;
            Name = name;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Type = type;
        }
    }
}
