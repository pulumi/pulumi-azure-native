// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20210201.Outputs
{

    /// <summary>
    /// Contains FQDN of the DNS record associated with the public IP address.
    /// </summary>
    [OutputType]
    public sealed class PublicIPAddressDnsSettingsResponse
    {
        /// <summary>
        /// The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
        /// </summary>
        public readonly string? DomainNameLabel;
        /// <summary>
        /// The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
        /// </summary>
        public readonly string? Fqdn;
        /// <summary>
        /// The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
        /// </summary>
        public readonly string? ReverseFqdn;

        [OutputConstructor]
        private PublicIPAddressDnsSettingsResponse(
            string? domainNameLabel,

            string? fqdn,

            string? reverseFqdn)
        {
            DomainNameLabel = domainNameLabel;
            Fqdn = fqdn;
            ReverseFqdn = reverseFqdn;
        }
    }
}
