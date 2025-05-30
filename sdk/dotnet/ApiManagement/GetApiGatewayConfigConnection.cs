// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetApiGatewayConfigConnection
    {
        /// <summary>
        /// Gets an API Management gateway config connection resource description.
        /// 
        /// Uses Azure REST API version 2024-06-01-preview.
        /// 
        /// Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetApiGatewayConfigConnectionResult> InvokeAsync(GetApiGatewayConfigConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiGatewayConfigConnectionResult>("azure-native:apimanagement:getApiGatewayConfigConnection", args ?? new GetApiGatewayConfigConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an API Management gateway config connection resource description.
        /// 
        /// Uses Azure REST API version 2024-06-01-preview.
        /// 
        /// Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiGatewayConfigConnectionResult> Invoke(GetApiGatewayConfigConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiGatewayConfigConnectionResult>("azure-native:apimanagement:getApiGatewayConfigConnection", args ?? new GetApiGatewayConfigConnectionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an API Management gateway config connection resource description.
        /// 
        /// Uses Azure REST API version 2024-06-01-preview.
        /// 
        /// Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiGatewayConfigConnectionResult> Invoke(GetApiGatewayConfigConnectionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiGatewayConfigConnectionResult>("azure-native:apimanagement:getApiGatewayConfigConnection", args ?? new GetApiGatewayConfigConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiGatewayConfigConnectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management gateway config connection.
        /// </summary>
        [Input("configConnectionName", required: true)]
        public string ConfigConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public string GatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApiGatewayConfigConnectionArgs()
        {
        }
        public static new GetApiGatewayConfigConnectionArgs Empty => new GetApiGatewayConfigConnectionArgs();
    }

    public sealed class GetApiGatewayConfigConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management gateway config connection.
        /// </summary>
        [Input("configConnectionName", required: true)]
        public Input<string> ConfigConnectionName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApiGatewayConfigConnectionInvokeArgs()
        {
        }
        public static new GetApiGatewayConfigConnectionInvokeArgs Empty => new GetApiGatewayConfigConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiGatewayConfigConnectionResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The default hostname of the data-plane gateway.
        /// </summary>
        public readonly string DefaultHostname;
        /// <summary>
        /// ETag of the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The hostnames of the data-plane gateway to which requests can be sent.
        /// </summary>
        public readonly ImmutableArray<string> Hostnames;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current provisioning state of the API Management gateway config connection 
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The link to the API Management service workspace.
        /// </summary>
        public readonly string? SourceId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiGatewayConfigConnectionResult(
            string azureApiVersion,

            string defaultHostname,

            string etag,

            ImmutableArray<string> hostnames,

            string id,

            string name,

            string provisioningState,

            string? sourceId,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DefaultHostname = defaultHostname;
            Etag = etag;
            Hostnames = hostnames;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            SourceId = sourceId;
            Type = type;
        }
    }
}
