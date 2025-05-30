// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of IpPermission
    /// </summary>
    public sealed class IpPermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).&lt;/p&gt;
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// &lt;p&gt;The IP protocol name (&lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, &lt;code&gt;icmp&lt;/code&gt;, &lt;code&gt;icmpv6&lt;/code&gt;) or number (see &lt;a href='http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml'&gt;Protocol Numbers&lt;/a&gt;).&lt;/p&gt; &lt;p&gt;Use &lt;code&gt;-1&lt;/code&gt; to specify all protocols. When authorizing security group rules, specifying &lt;code&gt;-1&lt;/code&gt; or a protocol number other than &lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, &lt;code&gt;icmp&lt;/code&gt;, or &lt;code&gt;icmpv6&lt;/code&gt; allows traffic on all ports, regardless of any port range you specify. For &lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, and &lt;code&gt;icmp&lt;/code&gt;, you must specify a port range. For &lt;code&gt;icmpv6&lt;/code&gt;, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.&lt;/p&gt;
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("ipRanges")]
        private InputList<Inputs.IpRangeArgs>? _ipRanges;

        /// <summary>
        /// &lt;p&gt;The IPv4 address ranges.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.IpRangeArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.IpRangeArgs>());
            set => _ipRanges = value;
        }

        [Input("ipv6Ranges")]
        private InputList<Inputs.Ipv6RangeArgs>? _ipv6Ranges;

        /// <summary>
        /// &lt;p&gt;The IPv6 address ranges.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.Ipv6RangeArgs> Ipv6Ranges
        {
            get => _ipv6Ranges ?? (_ipv6Ranges = new InputList<Inputs.Ipv6RangeArgs>());
            set => _ipv6Ranges = value;
        }

        [Input("prefixListIds")]
        private InputList<Inputs.PrefixListIdArgs>? _prefixListIds;

        /// <summary>
        /// &lt;p&gt;The prefix list IDs.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.PrefixListIdArgs> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<Inputs.PrefixListIdArgs>());
            set => _prefixListIds = value;
        }

        /// <summary>
        /// &lt;p&gt;If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).&lt;/p&gt;
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        [Input("userIdGroupPairs")]
        private InputList<Inputs.UserIdGroupPairArgs>? _userIdGroupPairs;

        /// <summary>
        /// &lt;p&gt;The security group and Amazon Web Services account ID pairs.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.UserIdGroupPairArgs> UserIdGroupPairs
        {
            get => _userIdGroupPairs ?? (_userIdGroupPairs = new InputList<Inputs.UserIdGroupPairArgs>());
            set => _userIdGroupPairs = value;
        }

        public IpPermissionArgs()
        {
        }
        public static new IpPermissionArgs Empty => new IpPermissionArgs();
    }
}
