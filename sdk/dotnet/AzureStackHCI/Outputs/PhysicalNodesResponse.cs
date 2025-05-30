// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// The PhysicalNodes of a cluster.
    /// </summary>
    [OutputType]
    public sealed class PhysicalNodesResponse
    {
        /// <summary>
        /// The IPv4 address assigned to each physical server on your Azure Stack HCI cluster.
        /// </summary>
        public readonly string? Ipv4Address;
        /// <summary>
        /// NETBIOS name of each physical server on your Azure Stack HCI cluster.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private PhysicalNodesResponse(
            string? ipv4Address,

            string? name)
        {
            Ipv4Address = ipv4Address;
            Name = name;
        }
    }
}
