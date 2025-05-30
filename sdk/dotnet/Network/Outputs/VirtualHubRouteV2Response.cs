// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// VirtualHubRouteTableV2 route.
    /// </summary>
    [OutputType]
    public sealed class VirtualHubRouteV2Response
    {
        /// <summary>
        /// The type of destinations.
        /// </summary>
        public readonly string? DestinationType;
        /// <summary>
        /// List of all destinations.
        /// </summary>
        public readonly ImmutableArray<string> Destinations;
        /// <summary>
        /// The type of next hops.
        /// </summary>
        public readonly string? NextHopType;
        /// <summary>
        /// NextHops ip address.
        /// </summary>
        public readonly ImmutableArray<string> NextHops;

        [OutputConstructor]
        private VirtualHubRouteV2Response(
            string? destinationType,

            ImmutableArray<string> destinations,

            string? nextHopType,

            ImmutableArray<string> nextHops)
        {
            DestinationType = destinationType;
            Destinations = destinations;
            NextHopType = nextHopType;
            NextHops = nextHops;
        }
    }
}
