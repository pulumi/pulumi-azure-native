// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Outputs
{

    /// <summary>
    /// An array of destination IPv4 Addresses or IPv6 Addresses.
    /// </summary>
    [OutputType]
    public sealed class NeighborGroupDestinationResponse
    {
        /// <summary>
        /// Array of IPv4 Addresses.
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Addresses;
        /// <summary>
        /// Array of IPv6 Addresses.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Addresses;

        [OutputConstructor]
        private NeighborGroupDestinationResponse(
            ImmutableArray<string> ipv4Addresses,

            ImmutableArray<string> ipv6Addresses)
        {
            Ipv4Addresses = ipv4Addresses;
            Ipv6Addresses = ipv6Addresses;
        }
    }
}
