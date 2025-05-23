// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetworkCloud.Inputs
{

    public sealed class L2ServiceLoadBalancerConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipAddressPools")]
        private InputList<Inputs.IpAddressPoolArgs>? _ipAddressPools;

        /// <summary>
        /// The list of pools of IP addresses that can be allocated to load balancer services.
        /// </summary>
        public InputList<Inputs.IpAddressPoolArgs> IpAddressPools
        {
            get => _ipAddressPools ?? (_ipAddressPools = new InputList<Inputs.IpAddressPoolArgs>());
            set => _ipAddressPools = value;
        }

        public L2ServiceLoadBalancerConfigurationArgs()
        {
        }
        public static new L2ServiceLoadBalancerConfigurationArgs Empty => new L2ServiceLoadBalancerConfigurationArgs();
    }
}
