// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20231115Preview.Outputs
{

    /// <summary>
    /// LoadBalancerProfile - Profile of the cluster load balancer.
    /// </summary>
    [OutputType]
    public sealed class NetworkProfileResponseLoadBalancerProfile
    {
        /// <summary>
        /// Count - Number of load balancer VMs. The default value is 0.
        /// </summary>
        public readonly int? Count;

        [OutputConstructor]
        private NetworkProfileResponseLoadBalancerProfile(int? count)
        {
            Count = count;
        }
    }
}
