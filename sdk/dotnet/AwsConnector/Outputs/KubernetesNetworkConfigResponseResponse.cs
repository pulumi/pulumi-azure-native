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
    /// Definition of KubernetesNetworkConfigResponse
    /// </summary>
    [OutputType]
    public sealed class KubernetesNetworkConfigResponseResponse
    {
        /// <summary>
        /// &lt;p&gt;The IP family used to assign Kubernetes &lt;code&gt;Pod&lt;/code&gt; and &lt;code&gt;Service&lt;/code&gt; objects IP addresses. The IP family is always &lt;code&gt;ipv4&lt;/code&gt;, unless you have a &lt;code&gt;1.21&lt;/code&gt; or later cluster running version &lt;code&gt;1.10.1&lt;/code&gt; or later of the Amazon VPC CNI plugin for Kubernetes and specified &lt;code&gt;ipv6&lt;/code&gt; when you created the cluster. &lt;/p&gt;
        /// </summary>
        public readonly Outputs.IpFamilyEnumValueResponse? IpFamily;
        /// <summary>
        /// &lt;p&gt;The CIDR block that Kubernetes &lt;code&gt;Pod&lt;/code&gt; and &lt;code&gt;Service&lt;/code&gt; object IP addresses are assigned from. Kubernetes assigns addresses from an &lt;code&gt;IPv4&lt;/code&gt; CIDR block assigned to a subnet that the node is in. If you didn't specify a CIDR block when you created the cluster, then Kubernetes assigns addresses from either the &lt;code&gt;10.100.0.0/16&lt;/code&gt; or &lt;code&gt;172.20.0.0/16&lt;/code&gt; CIDR blocks. If this was specified, then it was specified when the cluster was created and it can't be changed.&lt;/p&gt;
        /// </summary>
        public readonly string? ServiceIpv4Cidr;
        /// <summary>
        /// &lt;p&gt;The CIDR block that Kubernetes pod and service IP addresses are assigned from if you created a 1.21 or later cluster with version 1.10.1 or later of the Amazon VPC CNI add-on and specified &lt;code&gt;ipv6&lt;/code&gt; for &lt;b&gt;ipFamily&lt;/b&gt; when you created the cluster. Kubernetes assigns service addresses from the unique local address range (&lt;code&gt;fc00::/7&lt;/code&gt;) because you can't specify a custom IPv6 CIDR block when you create the cluster.&lt;/p&gt;
        /// </summary>
        public readonly string? ServiceIpv6Cidr;

        [OutputConstructor]
        private KubernetesNetworkConfigResponseResponse(
            Outputs.IpFamilyEnumValueResponse? ipFamily,

            string? serviceIpv4Cidr,

            string? serviceIpv6Cidr)
        {
            IpFamily = ipFamily;
            ServiceIpv4Cidr = serviceIpv4Cidr;
            ServiceIpv6Cidr = serviceIpv6Cidr;
        }
    }
}
