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
    /// Desired managed outbound IPs for the cluster load balancer.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs
    {
        /// <summary>
        /// The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack. 
        /// </summary>
        public readonly int? CountIPv6;

        [OutputConstructor]
        private ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs(
            int? count,

            int? countIPv6)
        {
            Count = count;
            CountIPv6 = countIPv6;
        }
    }
}
