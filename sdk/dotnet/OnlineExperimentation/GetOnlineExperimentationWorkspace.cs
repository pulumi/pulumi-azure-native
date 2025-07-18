// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OnlineExperimentation
{
    public static class GetOnlineExperimentationWorkspace
    {
        /// <summary>
        /// Gets an online experimentation workspace.
        /// 
        /// Uses Azure REST API version 2025-05-31-preview.
        /// 
        /// Other available API versions: 2025-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native onlineexperimentation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetOnlineExperimentationWorkspaceResult> InvokeAsync(GetOnlineExperimentationWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOnlineExperimentationWorkspaceResult>("azure-native:onlineexperimentation:getOnlineExperimentationWorkspace", args ?? new GetOnlineExperimentationWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an online experimentation workspace.
        /// 
        /// Uses Azure REST API version 2025-05-31-preview.
        /// 
        /// Other available API versions: 2025-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native onlineexperimentation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOnlineExperimentationWorkspaceResult> Invoke(GetOnlineExperimentationWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOnlineExperimentationWorkspaceResult>("azure-native:onlineexperimentation:getOnlineExperimentationWorkspace", args ?? new GetOnlineExperimentationWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an online experimentation workspace.
        /// 
        /// Uses Azure REST API version 2025-05-31-preview.
        /// 
        /// Other available API versions: 2025-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native onlineexperimentation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOnlineExperimentationWorkspaceResult> Invoke(GetOnlineExperimentationWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOnlineExperimentationWorkspaceResult>("azure-native:onlineexperimentation:getOnlineExperimentationWorkspace", args ?? new GetOnlineExperimentationWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOnlineExperimentationWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OnlineExperimentationWorkspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetOnlineExperimentationWorkspaceArgs()
        {
        }
        public static new GetOnlineExperimentationWorkspaceArgs Empty => new GetOnlineExperimentationWorkspaceArgs();
    }

    public sealed class GetOnlineExperimentationWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the OnlineExperimentationWorkspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetOnlineExperimentationWorkspaceInvokeArgs()
        {
        }
        public static new GetOnlineExperimentationWorkspaceInvokeArgs Empty => new GetOnlineExperimentationWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetOnlineExperimentationWorkspaceResult
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
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.OnlineExperimentationWorkspacePropertiesResponse Properties;
        /// <summary>
        /// The SKU (Stock Keeping Unit) assigned to this resource.
        /// </summary>
        public readonly Outputs.OnlineExperimentationWorkspaceSkuResponse? Sku;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetOnlineExperimentationWorkspaceResult(
            string azureApiVersion,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.OnlineExperimentationWorkspacePropertiesResponse properties,

            Outputs.OnlineExperimentationWorkspaceSkuResponse? sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
