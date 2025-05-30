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
    /// The operation for configuring intrusion detection.
    /// </summary>
    [OutputType]
    public sealed class FirewallPolicyIntrusionDetectionConfigurationResponse
    {
        /// <summary>
        /// List of rules for traffic to bypass.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyIntrusionDetectionBypassTrafficSpecificationsResponse> BypassTrafficSettings;
        /// <summary>
        /// IDPS Private IP address ranges are used to identify traffic direction (i.e. inbound, outbound, etc.). By default, only ranges defined by IANA RFC 1918 are considered private IP addresses. To modify default ranges, specify your Private IP address ranges with this property
        /// </summary>
        public readonly ImmutableArray<string> PrivateRanges;
        /// <summary>
        /// List of specific signatures states.
        /// </summary>
        public readonly ImmutableArray<Outputs.FirewallPolicyIntrusionDetectionSignatureSpecificationResponse> SignatureOverrides;

        [OutputConstructor]
        private FirewallPolicyIntrusionDetectionConfigurationResponse(
            ImmutableArray<Outputs.FirewallPolicyIntrusionDetectionBypassTrafficSpecificationsResponse> bypassTrafficSettings,

            ImmutableArray<string> privateRanges,

            ImmutableArray<Outputs.FirewallPolicyIntrusionDetectionSignatureSpecificationResponse> signatureOverrides)
        {
            BypassTrafficSettings = bypassTrafficSettings;
            PrivateRanges = privateRanges;
            SignatureOverrides = signatureOverrides;
        }
    }
}
