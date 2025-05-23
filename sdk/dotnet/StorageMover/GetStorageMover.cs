// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageMover
{
    public static class GetStorageMover
    {
        /// <summary>
        /// Gets a Storage Mover resource.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-07-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagemover [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetStorageMoverResult> InvokeAsync(GetStorageMoverArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageMoverResult>("azure-native:storagemover:getStorageMover", args ?? new GetStorageMoverArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Storage Mover resource.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-07-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagemover [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetStorageMoverResult> Invoke(GetStorageMoverInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageMoverResult>("azure-native:storagemover:getStorageMover", args ?? new GetStorageMoverInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Storage Mover resource.
        /// 
        /// Uses Azure REST API version 2024-07-01.
        /// 
        /// Other available API versions: 2023-03-01, 2023-07-01-preview, 2023-10-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagemover [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetStorageMoverResult> Invoke(GetStorageMoverInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageMoverResult>("azure-native:storagemover:getStorageMover", args ?? new GetStorageMoverInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageMoverArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Mover resource.
        /// </summary>
        [Input("storageMoverName", required: true)]
        public string StorageMoverName { get; set; } = null!;

        public GetStorageMoverArgs()
        {
        }
        public static new GetStorageMoverArgs Empty => new GetStorageMoverArgs();
    }

    public sealed class GetStorageMoverInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Mover resource.
        /// </summary>
        [Input("storageMoverName", required: true)]
        public Input<string> StorageMoverName { get; set; } = null!;

        public GetStorageMoverInvokeArgs()
        {
        }
        public static new GetStorageMoverInvokeArgs Empty => new GetStorageMoverInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageMoverResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// A description for the Storage Mover.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
        /// The provisioning state of this resource.
        /// </summary>
        public readonly string ProvisioningState;
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
        private GetStorageMoverResult(
            string azureApiVersion,

            string? description,

            string id,

            string location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Description = description;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
