// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageDiscovery
{
    public static class GetStorageDiscoveryWorkspace
    {
        /// <summary>
        /// Get a StorageDiscoveryWorkspace
        /// 
        /// Uses Azure REST API version 2025-06-01-preview.
        /// </summary>
        public static Task<GetStorageDiscoveryWorkspaceResult> InvokeAsync(GetStorageDiscoveryWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageDiscoveryWorkspaceResult>("azure-native:storagediscovery:getStorageDiscoveryWorkspace", args ?? new GetStorageDiscoveryWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Get a StorageDiscoveryWorkspace
        /// 
        /// Uses Azure REST API version 2025-06-01-preview.
        /// </summary>
        public static Output<GetStorageDiscoveryWorkspaceResult> Invoke(GetStorageDiscoveryWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageDiscoveryWorkspaceResult>("azure-native:storagediscovery:getStorageDiscoveryWorkspace", args ?? new GetStorageDiscoveryWorkspaceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a StorageDiscoveryWorkspace
        /// 
        /// Uses Azure REST API version 2025-06-01-preview.
        /// </summary>
        public static Output<GetStorageDiscoveryWorkspaceResult> Invoke(GetStorageDiscoveryWorkspaceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageDiscoveryWorkspaceResult>("azure-native:storagediscovery:getStorageDiscoveryWorkspace", args ?? new GetStorageDiscoveryWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageDiscoveryWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the StorageDiscoveryWorkspace
        /// </summary>
        [Input("storageDiscoveryWorkspaceName", required: true)]
        public string StorageDiscoveryWorkspaceName { get; set; } = null!;

        public GetStorageDiscoveryWorkspaceArgs()
        {
        }
        public static new GetStorageDiscoveryWorkspaceArgs Empty => new GetStorageDiscoveryWorkspaceArgs();
    }

    public sealed class GetStorageDiscoveryWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the StorageDiscoveryWorkspace
        /// </summary>
        [Input("storageDiscoveryWorkspaceName", required: true)]
        public Input<string> StorageDiscoveryWorkspaceName { get; set; } = null!;

        public GetStorageDiscoveryWorkspaceInvokeArgs()
        {
        }
        public static new GetStorageDiscoveryWorkspaceInvokeArgs Empty => new GetStorageDiscoveryWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageDiscoveryWorkspaceResult
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
        public readonly Outputs.StorageDiscoveryWorkspacePropertiesResponse Properties;
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
        private GetStorageDiscoveryWorkspaceResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.StorageDiscoveryWorkspacePropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
