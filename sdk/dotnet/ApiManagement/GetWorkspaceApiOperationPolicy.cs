// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetWorkspaceApiOperationPolicy
    {
        /// <summary>
        /// Get the policy configuration at the API Operation level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWorkspaceApiOperationPolicyResult> InvokeAsync(GetWorkspaceApiOperationPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceApiOperationPolicyResult>("azure-native:apimanagement:getWorkspaceApiOperationPolicy", args ?? new GetWorkspaceApiOperationPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Get the policy configuration at the API Operation level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceApiOperationPolicyResult> Invoke(GetWorkspaceApiOperationPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceApiOperationPolicyResult>("azure-native:apimanagement:getWorkspaceApiOperationPolicy", args ?? new GetWorkspaceApiOperationPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the policy configuration at the API Operation level.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceApiOperationPolicyResult> Invoke(GetWorkspaceApiOperationPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceApiOperationPolicyResult>("azure-native:apimanagement:getWorkspaceApiOperationPolicy", args ?? new GetWorkspaceApiOperationPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceApiOperationPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Policy Export Format.
        /// </summary>
        [Input("format")]
        public string? Format { get; set; }

        /// <summary>
        /// Operation identifier within an API. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationId", required: true)]
        public string OperationId { get; set; } = null!;

        /// <summary>
        /// The identifier of the Policy.
        /// </summary>
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

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

        public GetWorkspaceApiOperationPolicyArgs()
        {
        }
        public static new GetWorkspaceApiOperationPolicyArgs Empty => new GetWorkspaceApiOperationPolicyArgs();
    }

    public sealed class GetWorkspaceApiOperationPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Policy Export Format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Operation identifier within an API. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("operationId", required: true)]
        public Input<string> OperationId { get; set; } = null!;

        /// <summary>
        /// The identifier of the Policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

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

        public GetWorkspaceApiOperationPolicyInvokeArgs()
        {
        }
        public static new GetWorkspaceApiOperationPolicyInvokeArgs Empty => new GetWorkspaceApiOperationPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceApiOperationPolicyResult
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
        private GetWorkspaceApiOperationPolicyResult(
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
