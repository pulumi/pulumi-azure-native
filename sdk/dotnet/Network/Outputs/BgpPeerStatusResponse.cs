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
    /// BGP peer status details.
    /// </summary>
    [OutputType]
    public sealed class BgpPeerStatusResponse
    {
        /// <summary>
        /// The autonomous system number of the remote BGP peer.
        /// </summary>
        public readonly double Asn;
        /// <summary>
        /// For how long the peering has been up.
        /// </summary>
        public readonly string ConnectedDuration;
        /// <summary>
        /// The virtual network gateway's local address.
        /// </summary>
        public readonly string LocalAddress;
        /// <summary>
        /// The number of BGP messages received.
        /// </summary>
        public readonly double MessagesReceived;
        /// <summary>
        /// The number of BGP messages sent.
        /// </summary>
        public readonly double MessagesSent;
        /// <summary>
        /// The remote BGP peer.
        /// </summary>
        public readonly string Neighbor;
        /// <summary>
        /// The number of routes learned from this peer.
        /// </summary>
        public readonly double RoutesReceived;
        /// <summary>
        /// The BGP peer state.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private BgpPeerStatusResponse(
            double asn,

            string connectedDuration,

            string localAddress,

            double messagesReceived,

            double messagesSent,

            string neighbor,

            double routesReceived,

            string state)
        {
            Asn = asn;
            ConnectedDuration = connectedDuration;
            LocalAddress = localAddress;
            MessagesReceived = messagesReceived;
            MessagesSent = messagesSent;
            Neighbor = neighbor;
            RoutesReceived = routesReceived;
            State = state;
        }
    }
}
