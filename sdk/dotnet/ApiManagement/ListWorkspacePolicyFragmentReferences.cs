// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class ListWorkspacePolicyFragmentReferences
    {
        /// <summary>
        /// Lists policy resources that reference the policy fragment.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListWorkspacePolicyFragmentReferencesResult> InvokeAsync(ListWorkspacePolicyFragmentReferencesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWorkspacePolicyFragmentReferencesResult>("azure-native:apimanagement:listWorkspacePolicyFragmentReferences", args ?? new ListWorkspacePolicyFragmentReferencesArgs(), options.WithDefaults());

        /// <summary>
        /// Lists policy resources that reference the policy fragment.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWorkspacePolicyFragmentReferencesResult> Invoke(ListWorkspacePolicyFragmentReferencesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkspacePolicyFragmentReferencesResult>("azure-native:apimanagement:listWorkspacePolicyFragmentReferences", args ?? new ListWorkspacePolicyFragmentReferencesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists policy resources that reference the policy fragment.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWorkspacePolicyFragmentReferencesResult> Invoke(ListWorkspacePolicyFragmentReferencesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWorkspacePolicyFragmentReferencesResult>("azure-native:apimanagement:listWorkspacePolicyFragmentReferences", args ?? new ListWorkspacePolicyFragmentReferencesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWorkspacePolicyFragmentReferencesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A resource identifier.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

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

        /// <summary>
        /// Number of records to skip.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// Number of records to return.
        /// </summary>
        [Input("top")]
        public int? Top { get; set; }

        /// <summary>
        /// Workspace identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("workspaceId", required: true)]
        public string WorkspaceId { get; set; } = null!;

        public ListWorkspacePolicyFragmentReferencesArgs()
        {
        }
        public static new ListWorkspacePolicyFragmentReferencesArgs Empty => new ListWorkspacePolicyFragmentReferencesArgs();
    }

    public sealed class ListWorkspacePolicyFragmentReferencesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A resource identifier.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

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

        /// <summary>
        /// Number of records to skip.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// Number of records to return.
        /// </summary>
        [Input("top")]
        public Input<int>? Top { get; set; }

        /// <summary>
        /// Workspace identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public ListWorkspacePolicyFragmentReferencesInvokeArgs()
        {
        }
        public static new ListWorkspacePolicyFragmentReferencesInvokeArgs Empty => new ListWorkspacePolicyFragmentReferencesInvokeArgs();
    }


    [OutputType]
    public sealed class ListWorkspacePolicyFragmentReferencesResult
    {
        /// <summary>
        /// Total record count number.
        /// </summary>
        public readonly double? Count;
        /// <summary>
        /// Next page link if any.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// A collection of resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceCollectionResponseValue> Value;

        [OutputConstructor]
        private ListWorkspacePolicyFragmentReferencesResult(
            double? count,

            string? nextLink,

            ImmutableArray<Outputs.ResourceCollectionResponseValue> value)
        {
            Count = count;
            NextLink = nextLink;
            Value = value;
        }
    }
}
