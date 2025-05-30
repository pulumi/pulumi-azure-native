// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// option A properties object
    /// </summary>
    public sealed class ExternalNetworkPropertiesOptionAPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BFD configuration properties
        /// </summary>
        [Input("bfdConfiguration")]
        public Input<Inputs.BfdConfigurationArgs>? BfdConfiguration { get; set; }

        /// <summary>
        /// Egress Acl. ARM resource ID of Access Control Lists.
        /// </summary>
        [Input("egressAclId")]
        public Input<string>? EgressAclId { get; set; }

        /// <summary>
        /// Ingress Acl. ARM resource ID of Access Control Lists.
        /// </summary>
        [Input("ingressAclId")]
        public Input<string>? IngressAclId { get; set; }

        /// <summary>
        /// MTU to use for option A peering.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// Peer ASN number.Example : 28
        /// </summary>
        [Input("peerASN", required: true)]
        public Input<double> PeerASN { get; set; } = null!;

        /// <summary>
        /// IPv4 Address Prefix.
        /// </summary>
        [Input("primaryIpv4Prefix")]
        public Input<string>? PrimaryIpv4Prefix { get; set; }

        /// <summary>
        /// IPv6 Address Prefix.
        /// </summary>
        [Input("primaryIpv6Prefix")]
        public Input<string>? PrimaryIpv6Prefix { get; set; }

        /// <summary>
        /// Secondary IPv4 Address Prefix.
        /// </summary>
        [Input("secondaryIpv4Prefix")]
        public Input<string>? SecondaryIpv4Prefix { get; set; }

        /// <summary>
        /// Secondary IPv6 Address Prefix.
        /// </summary>
        [Input("secondaryIpv6Prefix")]
        public Input<string>? SecondaryIpv6Prefix { get; set; }

        /// <summary>
        /// Vlan identifier. Example : 501
        /// </summary>
        [Input("vlanId", required: true)]
        public Input<int> VlanId { get; set; } = null!;

        public ExternalNetworkPropertiesOptionAPropertiesArgs()
        {
            Mtu = 1500;
        }
        public static new ExternalNetworkPropertiesOptionAPropertiesArgs Empty => new ExternalNetworkPropertiesOptionAPropertiesArgs();
    }
}
