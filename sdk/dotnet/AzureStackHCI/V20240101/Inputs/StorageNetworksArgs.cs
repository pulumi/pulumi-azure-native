// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240101.Inputs
{

    /// <summary>
    /// The StorageNetworks of a cluster.
    /// </summary>
    public sealed class StorageNetworksArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the storage network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the storage network adapter.
        /// </summary>
        [Input("networkAdapterName")]
        public Input<string>? NetworkAdapterName { get; set; }

        /// <summary>
        /// ID specified for the VLAN storage network. This setting is applied to the network interfaces that route the storage and VM migration traffic. 
        /// </summary>
        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        public StorageNetworksArgs()
        {
        }
        public static new StorageNetworksArgs Empty => new StorageNetworksArgs();
    }
}
