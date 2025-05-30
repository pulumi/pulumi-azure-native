// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    public static class GetHuntComment
    {
        /// <summary>
        /// Gets a hunt comment
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetHuntCommentResult> InvokeAsync(GetHuntCommentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHuntCommentResult>("azure-native:securityinsights:getHuntComment", args ?? new GetHuntCommentArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a hunt comment
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetHuntCommentResult> Invoke(GetHuntCommentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHuntCommentResult>("azure-native:securityinsights:getHuntComment", args ?? new GetHuntCommentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a hunt comment
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetHuntCommentResult> Invoke(GetHuntCommentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHuntCommentResult>("azure-native:securityinsights:getHuntComment", args ?? new GetHuntCommentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHuntCommentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hunt comment id (GUID)
        /// </summary>
        [Input("huntCommentId", required: true)]
        public string HuntCommentId { get; set; } = null!;

        /// <summary>
        /// The hunt id (GUID)
        /// </summary>
        [Input("huntId", required: true)]
        public string HuntId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetHuntCommentArgs()
        {
        }
        public static new GetHuntCommentArgs Empty => new GetHuntCommentArgs();
    }

    public sealed class GetHuntCommentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hunt comment id (GUID)
        /// </summary>
        [Input("huntCommentId", required: true)]
        public Input<string> HuntCommentId { get; set; } = null!;

        /// <summary>
        /// The hunt id (GUID)
        /// </summary>
        [Input("huntId", required: true)]
        public Input<string> HuntId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetHuntCommentInvokeArgs()
        {
        }
        public static new GetHuntCommentInvokeArgs Empty => new GetHuntCommentInvokeArgs();
    }


    [OutputType]
    public sealed class GetHuntCommentResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The message for the comment
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetHuntCommentResult(
            string azureApiVersion,

            string? etag,

            string id,

            string message,

            string name,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Id = id;
            Message = message;
            Name = name;
            SystemData = systemData;
            Type = type;
        }
    }
}
