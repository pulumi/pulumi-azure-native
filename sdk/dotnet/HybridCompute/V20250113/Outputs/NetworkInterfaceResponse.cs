// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20250113.Outputs
{

    /// <summary>
    /// Describes a network interface.
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceResponse
    {
        /// <summary>
        /// Represents the ID of the network interface.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The list of IP addresses in this interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpAddressResponse> IpAddresses;
        /// <summary>
        /// Represents MAC address of the network interface.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// Represents the name of the network interface.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private NetworkInterfaceResponse(
            string? id,

            ImmutableArray<Outputs.IpAddressResponse> ipAddresses,

            string? macAddress,

            string? name)
        {
            Id = id;
            IpAddresses = ipAddresses;
            MacAddress = macAddress;
            Name = name;
        }
    }
}
