// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Specifies the peering configuration.
    /// </summary>
    public sealed class ExpressRouteCircuitPeeringConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("advertisedCommunities")]
        private InputList<string>? _advertisedCommunities;

        /// <summary>
        /// The communities of bgp peering. Specified for microsoft peering.
        /// </summary>
        public InputList<string> AdvertisedCommunities
        {
            get => _advertisedCommunities ?? (_advertisedCommunities = new InputList<string>());
            set => _advertisedCommunities = value;
        }

        [Input("advertisedPublicPrefixes")]
        private InputList<string>? _advertisedPublicPrefixes;

        /// <summary>
        /// The reference to AdvertisedPublicPrefixes.
        /// </summary>
        public InputList<string> AdvertisedPublicPrefixes
        {
            get => _advertisedPublicPrefixes ?? (_advertisedPublicPrefixes = new InputList<string>());
            set => _advertisedPublicPrefixes = value;
        }

        /// <summary>
        /// The CustomerASN of the peering.
        /// </summary>
        [Input("customerASN")]
        public Input<int>? CustomerASN { get; set; }

        /// <summary>
        /// The legacy mode of the peering.
        /// </summary>
        [Input("legacyMode")]
        public Input<int>? LegacyMode { get; set; }

        /// <summary>
        /// The RoutingRegistryName of the configuration.
        /// </summary>
        [Input("routingRegistryName")]
        public Input<string>? RoutingRegistryName { get; set; }

        public ExpressRouteCircuitPeeringConfigArgs()
        {
        }
        public static new ExpressRouteCircuitPeeringConfigArgs Empty => new ExpressRouteCircuitPeeringConfigArgs();
    }
}
