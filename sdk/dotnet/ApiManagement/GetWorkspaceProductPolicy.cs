// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetWorkspaceProductPolicy
    {
        /// <summary>
        /// Get the policy configuration at the Product level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWorkspaceProductPolicyResult> InvokeAsync(GetWorkspaceProductPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceProductPolicyResult>("azure-native:apimanagement:getWorkspaceProductPolicy", args ?? new GetWorkspaceProductPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Get the policy configuration at the Product level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceProductPolicyResult> Invoke(GetWorkspaceProductPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceProductPolicyResult>("azure-native:apimanagement:getWorkspaceProductPolicy", args ?? new GetWorkspaceProductPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the policy configuration at the Product level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceProductPolicyResult> Invoke(GetWorkspaceProductPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceProductPolicyResult>("azure-native:apimanagement:getWorkspaceProductPolicy", args ?? new GetWorkspaceProductPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceProductPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Policy Export Format.
        /// </summary>
        [Input("format")]
        public string? Format { get; set; }

        /// <summary>
        /// The identifier of the Policy.
        /// </summary>
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

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
        /// Workspace identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("workspaceId", required: true)]
        public string WorkspaceId { get; set; } = null!;

        public GetWorkspaceProductPolicyArgs()
        {
        }
        public static new GetWorkspaceProductPolicyArgs Empty => new GetWorkspaceProductPolicyArgs();
    }

    public sealed class GetWorkspaceProductPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Policy Export Format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// The identifier of the Policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

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
        /// Workspace identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public GetWorkspaceProductPolicyInvokeArgs()
        {
        }
        public static new GetWorkspaceProductPolicyInvokeArgs Empty => new GetWorkspaceProductPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceProductPolicyResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Format of the policyContent.
        /// </summary>
        public readonly string? Format;
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
        /// <summary>
        /// Contents of the Policy as defined by the format.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetWorkspaceProductPolicyResult(
            string azureApiVersion,

            string? format,

            string id,

            string name,

            string type,

            string value)
        {
            AzureApiVersion = azureApiVersion;
            Format = format;
            Id = id;
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
