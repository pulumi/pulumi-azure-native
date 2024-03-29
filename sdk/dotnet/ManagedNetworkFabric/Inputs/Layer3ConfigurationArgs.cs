// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// layer3Configuration
    /// </summary>
    public sealed class Layer3ConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// exportRoutePolicyId
        /// </summary>
        [Input("exportRoutePolicyId")]
        public Input<string>? ExportRoutePolicyId { get; set; }

        /// <summary>
        /// importRoutePolicyId
        /// </summary>
        [Input("importRoutePolicyId")]
        public Input<string>? ImportRoutePolicyId { get; set; }

        /// <summary>
        /// ASN of PE devices for CE/PE connectivity.Example : 28
        /// </summary>
        [Input("peerASN")]
        public Input<int>? PeerASN { get; set; }

        /// <summary>
        /// IPv4 Address Prefix of CE-PE interconnect links. Example: 172.31.0.0/31. The values can be specified at the time of creation or can be updated afterwards. Any update to the values post-provisioning may disrupt traffic. The 1st and 3rd IPs are to be configured on CE1 and CE2 for Option B interfaces. The 2nd and 4th IPs are to be configured on PE1 and PE2 for Option B interfaces.
        /// </summary>
        [Input("primaryIpv4Prefix")]
        public Input<string>? PrimaryIpv4Prefix { get; set; }

        /// <summary>
        /// IPv6 Address Prefix of CE-PE interconnect links. Example: 3FFE:FFFF:0:CD30::a0/126. The values can be specified at the time of creation or can be updated afterwards. Any update to the values post-provisioning may disrupt traffic. The 1st and 3rd IPs are to be configured on CE1 and CE2 for Option B interfaces. The 2nd and 4th IPs are to be configured on PE1 and PE2 for Option B interfaces.
        /// </summary>
        [Input("primaryIpv6Prefix")]
        public Input<string>? PrimaryIpv6Prefix { get; set; }

        /// <summary>
        /// Secondary IPv4 Address Prefix of CE-PE interconnect links. Example: 172.31.0.20/31. The values can be specified at the time of creation or can be updated afterwards. Any update to the values post-provisioning may disrupt traffic. The 1st and 3rd IPs are to be configured on CE1 and CE2 for Option B interfaces. The 2nd and 4th IPs are to be configured on PE1 and PE2 for Option B interfaces.
        /// </summary>
        [Input("secondaryIpv4Prefix")]
        public Input<string>? SecondaryIpv4Prefix { get; set; }

        /// <summary>
        /// Secondary IPv6 Address Prefix of CE-PE interconnect links. Example: 3FFE:FFFF:0:CD30::a4/126. The values can be specified at the time of creation or can be updated afterwards. Any update to the values post-provisioning may disrupt traffic. The 1st and 3rd IPs are to be configured on CE1 and CE2 for Option B interfaces. The 2nd and 4th IPs are to be configured on PE1 and PE2 for Option B interfaces.
        /// </summary>
        [Input("secondaryIpv6Prefix")]
        public Input<string>? SecondaryIpv6Prefix { get; set; }

        /// <summary>
        /// VLAN for CE/PE Layer 3 connectivity.Example : 501
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public Layer3ConfigurationArgs()
        {
        }
        public static new Layer3ConfigurationArgs Empty => new Layer3ConfigurationArgs();
    }
}
