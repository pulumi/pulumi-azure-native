// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.V20241101Preview.Inputs
{

    /// <summary>
    /// A network profile for network settings of a ContainerGroupProfile. Used to manage load balancer and application gateway backend pools, specifically updating the IP addresses of CGs within the backend pool.
    /// </summary>
    public sealed class NetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Gateway the CG profile will use to interact with CGs in a backend pool
        /// </summary>
        [Input("applicationGateway")]
        public Input<Inputs.ApplicationGatewayArgs>? ApplicationGateway { get; set; }

        /// <summary>
        /// LoadBalancer the CG profile will use to interact with CGs in a backend pool
        /// </summary>
        [Input("loadBalancer")]
        public Input<Inputs.LoadBalancerArgs>? LoadBalancer { get; set; }

        public NetworkProfileArgs()
        {
        }
        public static new NetworkProfileArgs Empty => new NetworkProfileArgs();
    }
}
