// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ElasticSan
{
    public static class GetVolumeSnapshot
    {
        /// <summary>
        /// Get a Volume Snapshot.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVolumeSnapshotResult> InvokeAsync(GetVolumeSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeSnapshotResult>("azure-native:elasticsan:getVolumeSnapshot", args ?? new GetVolumeSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Volume Snapshot.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVolumeSnapshotResult> Invoke(GetVolumeSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeSnapshotResult>("azure-native:elasticsan:getVolumeSnapshot", args ?? new GetVolumeSnapshotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Volume Snapshot.
        /// 
        /// Uses Azure REST API version 2024-05-01.
        /// 
        /// Other available API versions: 2023-01-01, 2024-06-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native elasticsan [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVolumeSnapshotResult> Invoke(GetVolumeSnapshotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeSnapshotResult>("azure-native:elasticsan:getVolumeSnapshot", args ?? new GetVolumeSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ElasticSan.
        /// </summary>
        [Input("elasticSanName", required: true)]
        public string ElasticSanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volume snapshot within the given volume group.
        /// </summary>
        [Input("snapshotName", required: true)]
        public string SnapshotName { get; set; } = null!;

        /// <summary>
        /// The name of the VolumeGroup.
        /// </summary>
        [Input("volumeGroupName", required: true)]
        public string VolumeGroupName { get; set; } = null!;

        public GetVolumeSnapshotArgs()
        {
        }
        public static new GetVolumeSnapshotArgs Empty => new GetVolumeSnapshotArgs();
    }

    public sealed class GetVolumeSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ElasticSan.
        /// </summary>
        [Input("elasticSanName", required: true)]
        public Input<string> ElasticSanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volume snapshot within the given volume group.
        /// </summary>
        [Input("snapshotName", required: true)]
        public Input<string> SnapshotName { get; set; } = null!;

        /// <summary>
        /// The name of the VolumeGroup.
        /// </summary>
        [Input("volumeGroupName", required: true)]
        public Input<string> VolumeGroupName { get; set; } = null!;

        public GetVolumeSnapshotInvokeArgs()
        {
        }
        public static new GetVolumeSnapshotInvokeArgs Empty => new GetVolumeSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeSnapshotResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Data used when creating a volume snapshot.
        /// </summary>
        public readonly Outputs.SnapshotCreationDataResponse CreationData;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the operation on the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Size of Source Volume
        /// </summary>
        public readonly double SourceVolumeSizeGiB;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Source Volume Name of a snapshot
        /// </summary>
        public readonly string VolumeName;

        [OutputConstructor]
        private GetVolumeSnapshotResult(
            string azureApiVersion,

            Outputs.SnapshotCreationDataResponse creationData,

            string id,

            string name,

            string provisioningState,

            double sourceVolumeSizeGiB,

            Outputs.SystemDataResponse systemData,

            string type,

            string volumeName)
        {
            AzureApiVersion = azureApiVersion;
            CreationData = creationData;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            SourceVolumeSizeGiB = sourceVolumeSizeGiB;
            SystemData = systemData;
            Type = type;
            VolumeName = volumeName;
        }
    }
}
