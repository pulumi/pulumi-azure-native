// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of IpPermission
    /// </summary>
    [OutputType]
    public sealed class IpPermissionResponse
    {
        /// <summary>
        /// &lt;p&gt;If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).&lt;/p&gt;
        /// </summary>
        public readonly int? FromPort;
        /// <summary>
        /// &lt;p&gt;The IP protocol name (&lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, &lt;code&gt;icmp&lt;/code&gt;, &lt;code&gt;icmpv6&lt;/code&gt;) or number (see &lt;a href='http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml'&gt;Protocol Numbers&lt;/a&gt;).&lt;/p&gt; &lt;p&gt;Use &lt;code&gt;-1&lt;/code&gt; to specify all protocols. When authorizing security group rules, specifying &lt;code&gt;-1&lt;/code&gt; or a protocol number other than &lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, &lt;code&gt;icmp&lt;/code&gt;, or &lt;code&gt;icmpv6&lt;/code&gt; allows traffic on all ports, regardless of any port range you specify. For &lt;code&gt;tcp&lt;/code&gt;, &lt;code&gt;udp&lt;/code&gt;, and &lt;code&gt;icmp&lt;/code&gt;, you must specify a port range. For &lt;code&gt;icmpv6&lt;/code&gt;, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.&lt;/p&gt;
        /// </summary>
        public readonly string? IpProtocol;
        /// <summary>
        /// &lt;p&gt;The IPv4 address ranges.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.IpRangeResponse> IpRanges;
        /// <summary>
        /// &lt;p&gt;The IPv6 address ranges.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.Ipv6RangeResponse> Ipv6Ranges;
        /// <summary>
        /// &lt;p&gt;The prefix list IDs.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.PrefixListIdResponse> PrefixListIds;
        /// <summary>
        /// &lt;p&gt;If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).&lt;/p&gt;
        /// </summary>
        public readonly int? ToPort;
        /// <summary>
        /// &lt;p&gt;The security group and Amazon Web Services account ID pairs.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.UserIdGroupPairResponse> UserIdGroupPairs;

        [OutputConstructor]
        private IpPermissionResponse(
            int? fromPort,

            string? ipProtocol,

            ImmutableArray<Outputs.IpRangeResponse> ipRanges,

            ImmutableArray<Outputs.Ipv6RangeResponse> ipv6Ranges,

            ImmutableArray<Outputs.PrefixListIdResponse> prefixListIds,

            int? toPort,

            ImmutableArray<Outputs.UserIdGroupPairResponse> userIdGroupPairs)
        {
            FromPort = fromPort;
            IpProtocol = ipProtocol;
            IpRanges = ipRanges;
            Ipv6Ranges = ipv6Ranges;
            PrefixListIds = prefixListIds;
            ToPort = toPort;
            UserIdGroupPairs = userIdGroupPairs;
        }
    }
}
