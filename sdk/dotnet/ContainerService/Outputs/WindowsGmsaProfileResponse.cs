// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// Windows gMSA Profile in the managed cluster.
    /// </summary>
    [OutputType]
    public sealed class WindowsGmsaProfileResponse
    {
        /// <summary>
        /// Specifies the DNS server for Windows gMSA. &lt;br&gt;&lt;br&gt; Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.
        /// </summary>
        public readonly string? DnsServer;
        /// <summary>
        /// Specifies whether to enable Windows gMSA in the managed cluster.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Specifies the root domain name for Windows gMSA. &lt;br&gt;&lt;br&gt; Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.
        /// </summary>
        public readonly string? RootDomainName;

        [OutputConstructor]
        private WindowsGmsaProfileResponse(
            string? dnsServer,

            bool? enabled,

            string? rootDomainName)
        {
            DnsServer = dnsServer;
            Enabled = enabled;
            RootDomainName = rootDomainName;
        }
    }
}
