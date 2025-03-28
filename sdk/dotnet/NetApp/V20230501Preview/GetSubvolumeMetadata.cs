// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.V20230501Preview
{
    public static class GetSubvolumeMetadata
    {
        /// <summary>
        /// Get details of the specified subvolume
        /// </summary>
        public static Task<GetSubvolumeMetadataResult> InvokeAsync(GetSubvolumeMetadataArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubvolumeMetadataResult>("azure-native:netapp/v20230501preview:getSubvolumeMetadata", args ?? new GetSubvolumeMetadataArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of the specified subvolume
        /// </summary>
        public static Output<GetSubvolumeMetadataResult> Invoke(GetSubvolumeMetadataInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubvolumeMetadataResult>("azure-native:netapp/v20230501preview:getSubvolumeMetadata", args ?? new GetSubvolumeMetadataInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of the specified subvolume
        /// </summary>
        public static Output<GetSubvolumeMetadataResult> Invoke(GetSubvolumeMetadataInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubvolumeMetadataResult>("azure-native:netapp/v20230501preview:getSubvolumeMetadata", args ?? new GetSubvolumeMetadataInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubvolumeMetadataArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the subvolume.
        /// </summary>
        [Input("subvolumeName", required: true)]
        public string SubvolumeName { get; set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public string VolumeName { get; set; } = null!;

        public GetSubvolumeMetadataArgs()
        {
        }
        public static new GetSubvolumeMetadataArgs Empty => new GetSubvolumeMetadataArgs();
    }

    public sealed class GetSubvolumeMetadataInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the capacity pool
        /// </summary>
        [Input("poolName", required: true)]
        public Input<string> PoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the subvolume.
        /// </summary>
        [Input("subvolumeName", required: true)]
        public Input<string> SubvolumeName { get; set; } = null!;

        /// <summary>
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        public GetSubvolumeMetadataInvokeArgs()
        {
        }
        public static new GetSubvolumeMetadataInvokeArgs Empty => new GetSubvolumeMetadataInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubvolumeMetadataResult
    {
        /// <summary>
        /// Most recent access time and date
        /// </summary>
        public readonly string? AccessedTimeStamp;
        /// <summary>
        /// Bytes used
        /// </summary>
        public readonly double? BytesUsed;
        /// <summary>
        /// Most recent change time and date
        /// </summary>
        public readonly string? ChangedTimeStamp;
        /// <summary>
        /// Creation time and date
        /// </summary>
        public readonly string? CreationTimeStamp;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Most recent modification time and date
        /// </summary>
        public readonly string? ModifiedTimeStamp;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Path to the parent subvolume
        /// </summary>
        public readonly string? ParentPath;
        /// <summary>
        /// Path to the subvolume
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Permissions of the subvolume
        /// </summary>
        public readonly string? Permissions;
        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Size of subvolume
        /// </summary>
        public readonly double? Size;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSubvolumeMetadataResult(
            string? accessedTimeStamp,

            double? bytesUsed,

            string? changedTimeStamp,

            string? creationTimeStamp,

            string id,

            string? modifiedTimeStamp,

            string name,

            string? parentPath,

            string? path,

            string? permissions,

            string? provisioningState,

            double? size,

            string type)
        {
            AccessedTimeStamp = accessedTimeStamp;
            BytesUsed = bytesUsed;
            ChangedTimeStamp = changedTimeStamp;
            CreationTimeStamp = creationTimeStamp;
            Id = id;
            ModifiedTimeStamp = modifiedTimeStamp;
            Name = name;
            ParentPath = parentPath;
            Path = path;
            Permissions = permissions;
            ProvisioningState = provisioningState;
            Size = size;
            Type = type;
        }
    }
}
