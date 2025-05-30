// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetApiWiki
    {
        /// <summary>
        /// Gets the details of the Wiki for an API specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetApiWikiResult> InvokeAsync(GetApiWikiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiWikiResult>("azure-native:apimanagement:getApiWiki", args ?? new GetApiWikiArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the Wiki for an API specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiWikiResult> Invoke(GetApiWikiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiWikiResult>("azure-native:apimanagement:getApiWiki", args ?? new GetApiWikiInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the Wiki for an API specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetApiWikiResult> Invoke(GetApiWikiInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiWikiResult>("azure-native:apimanagement:getApiWiki", args ?? new GetApiWikiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiWikiArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetApiWikiArgs()
        {
        }
        public static new GetApiWikiArgs Empty => new GetApiWikiArgs();
    }

    public sealed class GetApiWikiInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetApiWikiInvokeArgs()
        {
        }
        public static new GetApiWikiInvokeArgs Empty => new GetApiWikiInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiWikiResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Collection wiki documents included into this wiki.
        /// </summary>
        public readonly ImmutableArray<Outputs.WikiDocumentationContractResponse> Documents;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApiWikiResult(
            string azureApiVersion,

            ImmutableArray<Outputs.WikiDocumentationContractResponse> documents,

            string id,

            string name,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Documents = documents;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
