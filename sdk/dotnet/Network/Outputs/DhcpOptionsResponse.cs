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
    /// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for a subnet overrides VNET DHCP options.
    /// </summary>
    [OutputType]
    public sealed class DhcpOptionsResponse
    {
        /// <summary>
        /// The list of DNS servers IP addresses.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;

        [OutputConstructor]
        private DhcpOptionsResponse(ImmutableArray<string> dnsServers)
        {
            DnsServers = dnsServers;
        }
    }
}
