// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridConnectivity
{
    public static class GetEndpoint
    {
        /// <summary>
        /// Gets the endpoint to the resource.
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// 
        /// Other available API versions: 2023-03-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridconnectivity [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("azure-native:hybridconnectivity:getEndpoint", args ?? new GetEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the endpoint to the resource.
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// 
        /// Other available API versions: 2023-03-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridconnectivity [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("azure-native:hybridconnectivity:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the endpoint to the resource.
        /// 
        /// Uses Azure REST API version 2024-12-01.
        /// 
        /// Other available API versions: 2023-03-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridconnectivity [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("azure-native:hybridconnectivity:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetEndpointArgs()
        {
        }
        public static new GetEndpointArgs Empty => new GetEndpointArgs();
    }

    public sealed class GetEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The endpoint name.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The fully qualified Azure Resource manager identifier of the resource.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetEndpointInvokeArgs()
        {
        }
        public static new GetEndpointInvokeArgs Empty => new GetEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource Id of the connectivity endpoint (optional).
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetEndpointResult(
            string azureApiVersion,

            string id,

            string name,

            string provisioningState,

            string? resourceId,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceId = resourceId;
            SystemData = systemData;
            Type = type;
        }
    }
}
