// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab
{
    public static class GetDisk
    {
        /// <summary>
        /// Get disk.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Task<GetDiskResult> InvokeAsync(GetDiskArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiskResult>("azure-native:devtestlab:getDisk", args ?? new GetDiskArgs(), options.WithDefaults());

        /// <summary>
        /// Get disk.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetDiskResult> Invoke(GetDiskInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskResult>("azure-native:devtestlab:getDisk", args ?? new GetDiskInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get disk.
        /// 
        /// Uses Azure REST API version 2018-09-15.
        /// </summary>
        public static Output<GetDiskResult> Invoke(GetDiskInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiskResult>("azure-native:devtestlab:getDisk", args ?? new GetDiskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiskArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($select=diskType)'
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public string LabName { get; set; } = null!;

        /// <summary>
        /// The name of the Disk
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the user profile.
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetDiskArgs()
        {
        }
        public static new GetDiskArgs Empty => new GetDiskArgs();
    }

    public sealed class GetDiskInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the $expand query. Example: 'properties($select=diskType)'
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// The name of the lab.
        /// </summary>
        [Input("labName", required: true)]
        public Input<string> LabName { get; set; } = null!;

        /// <summary>
        /// The name of the Disk
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the user profile.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetDiskInvokeArgs()
        {
        }
        public static new GetDiskInvokeArgs Empty => new GetDiskInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiskResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The creation date of the disk.
        /// </summary>
        public readonly string CreatedDate;
        /// <summary>
        /// When backed by a blob, the name of the VHD blob without extension.
        /// </summary>
        public readonly string? DiskBlobName;
        /// <summary>
        /// The size of the disk in Gibibytes.
        /// </summary>
        public readonly int? DiskSizeGiB;
        /// <summary>
        /// The storage type for the disk (i.e. Standard, Premium).
        /// </summary>
        public readonly string? DiskType;
        /// <summary>
        /// When backed by a blob, the URI of underlying blob.
        /// </summary>
        public readonly string? DiskUri;
        /// <summary>
        /// The host caching policy of the disk (i.e. None, ReadOnly, ReadWrite).
        /// </summary>
        public readonly string? HostCaching;
        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource ID of the VM to which this disk is leased.
        /// </summary>
        public readonly string? LeasedByLabVmId;
        /// <summary>
        /// The location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// When backed by managed disk, this is the ID of the compute disk resource.
        /// </summary>
        public readonly string? ManagedDiskId;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// When backed by a blob, the storage account where the blob is.
        /// </summary>
        public readonly string? StorageAccountId;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        public readonly string UniqueIdentifier;

        [OutputConstructor]
        private GetDiskResult(
            string azureApiVersion,

            string createdDate,

            string? diskBlobName,

            int? diskSizeGiB,

            string? diskType,

            string? diskUri,

            string? hostCaching,

            string id,

            string? leasedByLabVmId,

            string? location,

            string? managedDiskId,

            string name,

            string provisioningState,

            string? storageAccountId,

            ImmutableDictionary<string, string>? tags,

            string type,

            string uniqueIdentifier)
        {
            AzureApiVersion = azureApiVersion;
            CreatedDate = createdDate;
            DiskBlobName = diskBlobName;
            DiskSizeGiB = diskSizeGiB;
            DiskType = diskType;
            DiskUri = diskUri;
            HostCaching = hostCaching;
            Id = id;
            LeasedByLabVmId = leasedByLabVmId;
            Location = location;
            ManagedDiskId = managedDiskId;
            Name = name;
            ProvisioningState = provisioningState;
            StorageAccountId = storageAccountId;
            Tags = tags;
            Type = type;
            UniqueIdentifier = uniqueIdentifier;
        }
    }
}
