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
    /// Load balancer backend addresses.
    /// </summary>
    [OutputType]
    public sealed class LoadBalancerBackendAddressResponse
    {
        /// <summary>
        /// A list of administrative states which once set can override health probe so that Load Balancer will always forward new connections to backend, or deny new connections and reset existing connections.
        /// </summary>
        public readonly string? AdminState;
        /// <summary>
        /// Collection of inbound NAT rule port mappings.
        /// </summary>
        public readonly ImmutableArray<Outputs.NatRulePortMappingResponse> InboundNatRulesPortMapping;
        /// <summary>
        /// IP Address belonging to the referenced virtual network.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// Reference to the frontend ip address configuration defined in regional loadbalancer.
        /// </summary>
        public readonly Outputs.SubResourceResponse? LoadBalancerFrontendIPConfiguration;
        /// <summary>
        /// Name of the backend address.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Reference to IP address defined in network interfaces.
        /// </summary>
        public readonly Outputs.SubResourceResponse NetworkInterfaceIPConfiguration;
        /// <summary>
        /// Reference to an existing subnet.
        /// </summary>
        public readonly Outputs.SubResourceResponse? Subnet;
        /// <summary>
        /// Reference to an existing virtual network.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualNetwork;

        [OutputConstructor]
        private LoadBalancerBackendAddressResponse(
            string? adminState,

            ImmutableArray<Outputs.NatRulePortMappingResponse> inboundNatRulesPortMapping,

            string? ipAddress,

            Outputs.SubResourceResponse? loadBalancerFrontendIPConfiguration,

            string? name,

            Outputs.SubResourceResponse networkInterfaceIPConfiguration,

            Outputs.SubResourceResponse? subnet,

            Outputs.SubResourceResponse? virtualNetwork)
        {
            AdminState = adminState;
            InboundNatRulesPortMapping = inboundNatRulesPortMapping;
            IpAddress = ipAddress;
            LoadBalancerFrontendIPConfiguration = loadBalancerFrontendIPConfiguration;
            Name = name;
            NetworkInterfaceIPConfiguration = networkInterfaceIPConfiguration;
            Subnet = subnet;
            VirtualNetwork = virtualNetwork;
        }
    }
}
