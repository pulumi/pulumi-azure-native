// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp
{
    public static class GetVolumeGroup
    {
        /// <summary>
        /// Get details of the specified volume group
        /// 
        /// Uses Azure REST API version 2022-11-01.
        /// 
        /// Other available API versions: 2021-10-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01, 2024-09-01-preview.
        /// </summary>
        public static Task<GetVolumeGroupResult> InvokeAsync(GetVolumeGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeGroupResult>("azure-native:netapp:getVolumeGroup", args ?? new GetVolumeGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of the specified volume group
        /// 
        /// Uses Azure REST API version 2022-11-01.
        /// 
        /// Other available API versions: 2021-10-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01, 2024-09-01-preview.
        /// </summary>
        public static Output<GetVolumeGroupResult> Invoke(GetVolumeGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeGroupResult>("azure-native:netapp:getVolumeGroup", args ?? new GetVolumeGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of the specified volume group
        /// 
        /// Uses Azure REST API version 2022-11-01.
        /// 
        /// Other available API versions: 2021-10-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01, 2024-09-01-preview.
        /// </summary>
        public static Output<GetVolumeGroupResult> Invoke(GetVolumeGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeGroupResult>("azure-native:netapp:getVolumeGroup", args ?? new GetVolumeGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volumeGroup
        /// </summary>
        [Input("volumeGroupName", required: true)]
        public string VolumeGroupName { get; set; } = null!;

        public GetVolumeGroupArgs()
        {
        }
        public static new GetVolumeGroupArgs Empty => new GetVolumeGroupArgs();
    }

    public sealed class GetVolumeGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp account
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volumeGroup
        /// </summary>
        [Input("volumeGroupName", required: true)]
        public Input<string> VolumeGroupName { get; set; } = null!;

        public GetVolumeGroupInvokeArgs()
        {
        }
        public static new GetVolumeGroupInvokeArgs Empty => new GetVolumeGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeGroupResult
    {
        /// <summary>
        /// Volume group details
        /// </summary>
        public readonly Outputs.VolumeGroupMetaDataResponse? GroupMetaData;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure lifecycle management
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// List of volumes from group
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeGroupVolumePropertiesResponse> Volumes;

        [OutputConstructor]
        private GetVolumeGroupResult(
            Outputs.VolumeGroupMetaDataResponse? groupMetaData,

            string id,

            string? location,

            string name,

            string provisioningState,

            string type,

            ImmutableArray<Outputs.VolumeGroupVolumePropertiesResponse> volumes)
        {
            GroupMetaData = groupMetaData;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Type = type;
            Volumes = volumes;
        }
    }
}
