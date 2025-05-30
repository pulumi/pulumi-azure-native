// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// The StorageAdapter physical nodes of a cluster.
    /// </summary>
    public sealed class StorageAdapterIPInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv4 address assigned to each storage adapter physical node on your Azure Stack HCI cluster.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        /// <summary>
        /// storage adapter physical node name.
        /// </summary>
        [Input("physicalNode")]
        public Input<string>? PhysicalNode { get; set; }

        /// <summary>
        /// The SubnetMask address assigned to each storage adapter physical node on your Azure Stack HCI cluster.
        /// </summary>
        [Input("subnetMask")]
        public Input<string>? SubnetMask { get; set; }

        public StorageAdapterIPInfoArgs()
        {
        }
        public static new StorageAdapterIPInfoArgs Empty => new StorageAdapterIPInfoArgs();
    }
}
