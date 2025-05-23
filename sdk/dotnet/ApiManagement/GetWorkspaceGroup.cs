// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    public static class GetWorkspaceGroup
    {
        /// <summary>
        /// Gets the details of the group specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWorkspaceGroupResult> InvokeAsync(GetWorkspaceGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceGroupResult>("azure-native:apimanagement:getWorkspaceGroup", args ?? new GetWorkspaceGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the group specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceGroupResult> Invoke(GetWorkspaceGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceGroupResult>("azure-native:apimanagement:getWorkspaceGroup", args ?? new GetWorkspaceGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the group specified by its identifier.
        /// 
        /// Uses Azure REST API version 2022-09-01-preview.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWorkspaceGroupResult> Invoke(GetWorkspaceGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceGroupResult>("azure-native:apimanagement:getWorkspaceGroup", args ?? new GetWorkspaceGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

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

        public GetWorkspaceGroupArgs()
        {
        }
        public static new GetWorkspaceGroupArgs Empty => new GetWorkspaceGroupArgs();
    }

    public sealed class GetWorkspaceGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

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

        public GetWorkspaceGroupInvokeArgs()
        {
        }
        public static new GetWorkspaceGroupInvokeArgs Empty => new GetWorkspaceGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceGroupResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
        /// </summary>
        public readonly bool BuiltIn;
        /// <summary>
        /// Group description. Can contain HTML formatting tags.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Group name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://&lt;tenant&gt;.onmicrosoft.com/groups/&lt;group object id&gt;`; otherwise the value is null.
        /// </summary>
        public readonly string? ExternalId;
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
        private GetWorkspaceGroupResult(
            string azureApiVersion,

            bool builtIn,

            string? description,

            string displayName,

            string? externalId,

            string id,

            string name,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            BuiltIn = builtIn;
            Description = description;
            DisplayName = displayName;
            ExternalId = externalId;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
