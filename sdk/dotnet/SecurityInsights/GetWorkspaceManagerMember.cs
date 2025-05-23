// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    public static class GetWorkspaceManagerMember
    {
        /// <summary>
        /// Gets a workspace manager member
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWorkspaceManagerMemberResult> InvokeAsync(GetWorkspaceManagerMemberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceManagerMemberResult>("azure-native:securityinsights:getWorkspaceManagerMember", args ?? new GetWorkspaceManagerMemberArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a workspace manager member
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceManagerMemberResult> Invoke(GetWorkspaceManagerMemberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceManagerMemberResult>("azure-native:securityinsights:getWorkspaceManagerMember", args ?? new GetWorkspaceManagerMemberInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a workspace manager member
        /// 
        /// Uses Azure REST API version 2025-01-01-preview.
        /// 
        /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceManagerMemberResult> Invoke(GetWorkspaceManagerMemberInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceManagerMemberResult>("azure-native:securityinsights:getWorkspaceManagerMember", args ?? new GetWorkspaceManagerMemberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceManagerMemberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace manager member
        /// </summary>
        [Input("workspaceManagerMemberName", required: true)]
        public string WorkspaceManagerMemberName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceManagerMemberArgs()
        {
        }
        public static new GetWorkspaceManagerMemberArgs Empty => new GetWorkspaceManagerMemberArgs();
    }

    public sealed class GetWorkspaceManagerMemberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace manager member
        /// </summary>
        [Input("workspaceManagerMemberName", required: true)]
        public Input<string> WorkspaceManagerMemberName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetWorkspaceManagerMemberInvokeArgs()
        {
        }
        public static new GetWorkspaceManagerMemberInvokeArgs Empty => new GetWorkspaceManagerMemberInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceManagerMemberResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Resource Etag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
        /// </summary>
        public readonly string TargetWorkspaceResourceId;
        /// <summary>
        /// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
        /// </summary>
        public readonly string TargetWorkspaceTenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWorkspaceManagerMemberResult(
            string azureApiVersion,

            string etag,

            string id,

            string name,

            Outputs.SystemDataResponse systemData,

            string targetWorkspaceResourceId,

            string targetWorkspaceTenantId,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Etag = etag;
            Id = id;
            Name = name;
            SystemData = systemData;
            TargetWorkspaceResourceId = targetWorkspaceResourceId;
            TargetWorkspaceTenantId = targetWorkspaceTenantId;
            Type = type;
        }
    }
}
