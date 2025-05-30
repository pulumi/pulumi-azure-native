// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.Outputs
{

    /// <summary>
    /// LoadBalancerProfile represents the profile of the cluster public load balancer.
    /// </summary>
    [OutputType]
    public sealed class LoadBalancerProfileResponse
    {
        /// <summary>
        /// The list of effective outbound IP addresses of the public load balancer.
        /// </summary>
        public readonly ImmutableArray<Outputs.EffectiveOutboundIPResponse> EffectiveOutboundIps;
        /// <summary>
        /// The desired managed outbound IPs for the cluster public load balancer.
        /// </summary>
        public readonly Outputs.ManagedOutboundIPsResponse? ManagedOutboundIps;

        [OutputConstructor]
        private LoadBalancerProfileResponse(
            ImmutableArray<Outputs.EffectiveOutboundIPResponse> effectiveOutboundIps,

            Outputs.ManagedOutboundIPsResponse? managedOutboundIps)
        {
            EffectiveOutboundIps = effectiveOutboundIps;
            ManagedOutboundIps = managedOutboundIps;
        }
    }
}
