// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// DNS server details
    /// </summary>
    [OutputType]
    public sealed class ActiveDirectoryConnectorDNSDetailsResponse
    {
        /// <summary>
        /// DNS domain name for which DNS lookups should be forwarded to the Active Directory DNS servers.
        /// </summary>
        public readonly string? DomainName;
        /// <summary>
        /// List of Active Directory DNS server IP addresses.
        /// </summary>
        public readonly ImmutableArray<string> NameserverIPAddresses;
        /// <summary>
        /// Flag indicating whether to prefer Kubernetes DNS server response over AD DNS server response for IP address lookups.
        /// </summary>
        public readonly bool? PreferK8sDnsForPtrLookups;
        /// <summary>
        /// Replica count for DNS proxy service. Default value is 1.
        /// </summary>
        public readonly double? Replicas;

        [OutputConstructor]
        private ActiveDirectoryConnectorDNSDetailsResponse(
            string? domainName,

            ImmutableArray<string> nameserverIPAddresses,

            bool? preferK8sDnsForPtrLookups,

            double? replicas)
        {
            DomainName = domainName;
            NameserverIPAddresses = nameserverIPAddresses;
            PreferK8sDnsForPtrLookups = preferK8sDnsForPtrLookups;
            Replicas = replicas;
        }
    }
}
