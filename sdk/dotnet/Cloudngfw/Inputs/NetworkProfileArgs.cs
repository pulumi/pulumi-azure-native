// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Inputs
{

    /// <summary>
    /// Network settings for Firewall
    /// </summary>
    public sealed class NetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("egressNatIp")]
        private InputList<Inputs.IPAddressArgs>? _egressNatIp;

        /// <summary>
        /// Egress nat IP to use
        /// </summary>
        public InputList<Inputs.IPAddressArgs> EgressNatIp
        {
            get => _egressNatIp ?? (_egressNatIp = new InputList<Inputs.IPAddressArgs>());
            set => _egressNatIp = value;
        }

        /// <summary>
        /// Enable egress NAT, enabled by default
        /// </summary>
        [Input("enableEgressNat", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cloudngfw.EgressNat> EnableEgressNat { get; set; } = null!;

        /// <summary>
        /// vnet or vwan, cannot be updated
        /// </summary>
        [Input("networkType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cloudngfw.NetworkType> NetworkType { get; set; } = null!;

        [Input("privateSourceNatRulesDestination")]
        private InputList<string>? _privateSourceNatRulesDestination;

        /// <summary>
        /// Array of ipv4 destination address for which source NAT is to be performed
        /// </summary>
        public InputList<string> PrivateSourceNatRulesDestination
        {
            get => _privateSourceNatRulesDestination ?? (_privateSourceNatRulesDestination = new InputList<string>());
            set => _privateSourceNatRulesDestination = value;
        }

        [Input("publicIps", required: true)]
        private InputList<Inputs.IPAddressArgs>? _publicIps;

        /// <summary>
        /// List of IPs associated with the Firewall
        /// </summary>
        public InputList<Inputs.IPAddressArgs> PublicIps
        {
            get => _publicIps ?? (_publicIps = new InputList<Inputs.IPAddressArgs>());
            set => _publicIps = value;
        }

        [Input("trustedRanges")]
        private InputList<string>? _trustedRanges;

        /// <summary>
        /// Non-RFC 1918 address
        /// </summary>
        public InputList<string> TrustedRanges
        {
            get => _trustedRanges ?? (_trustedRanges = new InputList<string>());
            set => _trustedRanges = value;
        }

        /// <summary>
        /// Vnet configurations
        /// </summary>
        [Input("vnetConfiguration")]
        public Input<Inputs.VnetConfigurationArgs>? VnetConfiguration { get; set; }

        /// <summary>
        /// Vwan configurations
        /// </summary>
        [Input("vwanConfiguration")]
        public Input<Inputs.VwanConfigurationArgs>? VwanConfiguration { get; set; }

        public NetworkProfileArgs()
        {
        }
        public static new NetworkProfileArgs Empty => new NetworkProfileArgs();
    }
}
