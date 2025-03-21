// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.V20240701
{
    public static class ListVolumeReplications
    {
        /// <summary>
        /// List all replications for a specified volume
        /// </summary>
        public static Task<ListVolumeReplicationsResult> InvokeAsync(ListVolumeReplicationsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListVolumeReplicationsResult>("azure-native:netapp/v20240701:listVolumeReplications", args ?? new ListVolumeReplicationsArgs(), options.WithDefaults());

        /// <summary>
        /// List all replications for a specified volume
        /// </summary>
        public static Output<ListVolumeReplicationsResult> Invoke(ListVolumeReplicationsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListVolumeReplicationsResult>("azure-native:netapp/v20240701:listVolumeReplications", args ?? new ListVolumeReplicationsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// List all replications for a specified volume
        /// </summary>
        public static Output<ListVolumeReplicationsResult> Invoke(ListVolumeReplicationsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListVolumeReplicationsResult>("azure-native:netapp/v20240701:listVolumeReplications", args ?? new ListVolumeReplicationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVolumeReplicationsArgs : global::Pulumi.InvokeArgs
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
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public string VolumeName { get; set; } = null!;

        public ListVolumeReplicationsArgs()
        {
        }
        public static new ListVolumeReplicationsArgs Empty => new ListVolumeReplicationsArgs();
    }

    public sealed class ListVolumeReplicationsInvokeArgs : global::Pulumi.InvokeArgs
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
        /// The name of the volume
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        public ListVolumeReplicationsInvokeArgs()
        {
        }
        public static new ListVolumeReplicationsInvokeArgs Empty => new ListVolumeReplicationsInvokeArgs();
    }


    [OutputType]
    public sealed class ListVolumeReplicationsResult
    {
        /// <summary>
        /// A list of replications
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicationResponse> Value;

        [OutputConstructor]
        private ListVolumeReplicationsResult(ImmutableArray<Outputs.ReplicationResponse> value)
        {
            Value = value;
        }
    }
}
