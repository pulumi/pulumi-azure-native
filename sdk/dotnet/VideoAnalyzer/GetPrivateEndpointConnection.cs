// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer
{
    public static class GetPrivateEndpointConnection
    {
        /// <summary>
        /// Get private endpoint connection under video analyzer account.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Task<GetPrivateEndpointConnectionResult> InvokeAsync(GetPrivateEndpointConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointConnectionResult>("azure-native:videoanalyzer:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Get private endpoint connection under video analyzer account.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:videoanalyzer:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get private endpoint connection under video analyzer account.
        /// 
        /// Uses Azure REST API version 2021-11-01-preview.
        /// </summary>
        public static Output<GetPrivateEndpointConnectionResult> Invoke(GetPrivateEndpointConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateEndpointConnectionResult>("azure-native:videoanalyzer:getPrivateEndpointConnection", args ?? new GetPrivateEndpointConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateEndpointConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Video Analyzer account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Private endpoint connection name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// The Video Analyzer account name.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Private endpoint connection name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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

            string id,

            string name,

            Outputs.PrivateEndpointResponse? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStateResponse privateLinkServiceConnectionState,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
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
