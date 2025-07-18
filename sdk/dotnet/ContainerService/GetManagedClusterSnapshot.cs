// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService
{
    public static class GetManagedClusterSnapshot
    {
        /// <summary>
        /// A managed cluster snapshot resource.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2022-02-02-preview, 2022-03-02-preview, 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-02-preview, 2023-10-02-preview, 2023-11-02-preview, 2024-01-02-preview, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetManagedClusterSnapshotResult> InvokeAsync(GetManagedClusterSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetManagedClusterSnapshotResult>("azure-native:containerservice:getManagedClusterSnapshot", args ?? new GetManagedClusterSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// A managed cluster snapshot resource.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2022-02-02-preview, 2022-03-02-preview, 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-02-preview, 2023-10-02-preview, 2023-11-02-preview, 2024-01-02-preview, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedClusterSnapshotResult> Invoke(GetManagedClusterSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedClusterSnapshotResult>("azure-native:containerservice:getManagedClusterSnapshot", args ?? new GetManagedClusterSnapshotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A managed cluster snapshot resource.
        /// 
        /// Uses Azure REST API version 2024-10-02-preview.
        /// 
        /// Other available API versions: 2022-02-02-preview, 2022-03-02-preview, 2022-04-02-preview, 2022-05-02-preview, 2022-06-02-preview, 2022-07-02-preview, 2022-08-02-preview, 2022-08-03-preview, 2022-09-02-preview, 2022-10-02-preview, 2022-11-02-preview, 2023-01-02-preview, 2023-02-02-preview, 2023-03-02-preview, 2023-04-02-preview, 2023-05-02-preview, 2023-06-02-preview, 2023-07-02-preview, 2023-08-02-preview, 2023-09-02-preview, 2023-10-02-preview, 2023-11-02-preview, 2024-01-02-preview, 2024-02-02-preview, 2024-03-02-preview, 2024-04-02-preview, 2024-05-02-preview, 2024-06-02-preview, 2024-07-02-preview, 2024-09-02-preview, 2025-01-02-preview, 2025-02-02-preview, 2025-03-02-preview, 2025-04-02-preview, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetManagedClusterSnapshotResult> Invoke(GetManagedClusterSnapshotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetManagedClusterSnapshotResult>("azure-native:containerservice:getManagedClusterSnapshot", args ?? new GetManagedClusterSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetManagedClusterSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetManagedClusterSnapshotArgs()
        {
        }
        public static new GetManagedClusterSnapshotArgs Empty => new GetManagedClusterSnapshotArgs();
    }

    public sealed class GetManagedClusterSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetManagedClusterSnapshotInvokeArgs()
        {
        }
        public static new GetManagedClusterSnapshotInvokeArgs Empty => new GetManagedClusterSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetManagedClusterSnapshotResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// CreationData to be used to specify the source resource ID to create this snapshot.
        /// </summary>
        public readonly Outputs.CreationDataResponse? CreationData;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// What the properties will be showed when getting managed cluster snapshot. Those properties are read-only.
        /// </summary>
        public readonly Outputs.ManagedClusterPropertiesForSnapshotResponse ManagedClusterPropertiesReadOnly;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of a snapshot. The default is NodePool.
        /// </summary>
        public readonly string? SnapshotType;
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
        private GetManagedClusterSnapshotResult(
            string azureApiVersion,

            Outputs.CreationDataResponse? creationData,

            string id,

            string location,

            Outputs.ManagedClusterPropertiesForSnapshotResponse managedClusterPropertiesReadOnly,

            string name,

            string? snapshotType,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CreationData = creationData;
            Id = id;
            Location = location;
            ManagedClusterPropertiesReadOnly = managedClusterPropertiesReadOnly;
            Name = name;
            SnapshotType = snapshotType;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
