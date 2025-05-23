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
    /// BGP settings details.
    /// </summary>
    public sealed class BgpSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The BGP speaker's ASN.
        /// </summary>
        [Input("asn")]
        public Input<double>? Asn { get; set; }

        /// <summary>
        /// The BGP peering address and BGP identifier of this BGP speaker.
        /// </summary>
        [Input("bgpPeeringAddress")]
        public Input<string>? BgpPeeringAddress { get; set; }

        [Input("bgpPeeringAddresses")]
        private InputList<Inputs.IPConfigurationBgpPeeringAddressArgs>? _bgpPeeringAddresses;

        /// <summary>
        /// BGP peering address with IP configuration ID for virtual network gateway.
        /// </summary>
        public InputList<Inputs.IPConfigurationBgpPeeringAddressArgs> BgpPeeringAddresses
        {
            get => _bgpPeeringAddresses ?? (_bgpPeeringAddresses = new InputList<Inputs.IPConfigurationBgpPeeringAddressArgs>());
            set => _bgpPeeringAddresses = value;
        }

        /// <summary>
        /// The weight added to routes learned from this BGP speaker.
        /// </summary>
        [Input("peerWeight")]
        public Input<int>? PeerWeight { get; set; }

        public BgpSettingsArgs()
        {
        }
        public static new BgpSettingsArgs Empty => new BgpSettingsArgs();
    }
}
