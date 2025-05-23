// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud.Outputs
{

    [OutputType]
    public sealed class L2ServiceLoadBalancerConfigurationResponse
    {
        /// <summary>
        /// The list of pools of IP addresses that can be allocated to load balancer services.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpAddressPoolResponse> IpAddressPools;

        [OutputConstructor]
        private L2ServiceLoadBalancerConfigurationResponse(ImmutableArray<Outputs.IpAddressPoolResponse> ipAddressPools)
        {
            IpAddressPools = ipAddressPools;
        }
    }
}
