// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Outputs
{

    /// <summary>
    /// security services
    /// </summary>
    [OutputType]
    public sealed class SecurityServicesResponse
    {
        /// <summary>
        /// Anti spyware Profile data
        /// </summary>
        public readonly string? AntiSpywareProfile;
        /// <summary>
        /// anti virus profile data
        /// </summary>
        public readonly string? AntiVirusProfile;
        /// <summary>
        /// DNS Subscription profile data
        /// </summary>
        public readonly string? DnsSubscription;
        /// <summary>
        /// File blocking profile data
        /// </summary>
        public readonly string? FileBlockingProfile;
        /// <summary>
        /// Trusted Egress Decryption profile data
        /// </summary>
        public readonly string? OutboundTrustCertificate;
        /// <summary>
        /// Untrusted Egress Decryption profile data
        /// </summary>
        public readonly string? OutboundUnTrustCertificate;
        /// <summary>
        /// URL filtering profile data
        /// </summary>
        public readonly string? UrlFilteringProfile;
        /// <summary>
        /// IPs Vulnerability Profile Data
        /// </summary>
        public readonly string? VulnerabilityProfile;

        [OutputConstructor]
        private SecurityServicesResponse(
            string? antiSpywareProfile,

            string? antiVirusProfile,

            string? dnsSubscription,

            string? fileBlockingProfile,

            string? outboundTrustCertificate,

            string? outboundUnTrustCertificate,

            string? urlFilteringProfile,

            string? vulnerabilityProfile)
        {
            AntiSpywareProfile = antiSpywareProfile;
            AntiVirusProfile = antiVirusProfile;
            DnsSubscription = dnsSubscription;
            FileBlockingProfile = fileBlockingProfile;
            OutboundTrustCertificate = outboundTrustCertificate;
            OutboundUnTrustCertificate = outboundUnTrustCertificate;
            UrlFilteringProfile = urlFilteringProfile;
            VulnerabilityProfile = vulnerabilityProfile;
        }
    }
}
