// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20180601.Outputs
{

    /// <summary>
    /// Outbound NAT pool of the load balancer.
    /// </summary>
    [OutputType]
    public sealed class OutboundNatRuleResponse
    {
        /// <summary>
        /// The number of outbound ports to be used for NAT.
        /// </summary>
        public readonly int? AllocatedOutboundPorts;
        /// <summary>
        /// A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend IPs.
        /// </summary>
        public readonly Outputs.SubResourceResponse BackendAddressPool;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The Frontend IP addresses of the load balancer.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> FrontendIPConfigurations;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private OutboundNatRuleResponse(
            int? allocatedOutboundPorts,

            Outputs.SubResourceResponse backendAddressPool,

            string? etag,

            ImmutableArray<Outputs.SubResourceResponse> frontendIPConfigurations,

            string? id,

            string? name,

            string? provisioningState)
        {
            AllocatedOutboundPorts = allocatedOutboundPorts;
            BackendAddressPool = backendAddressPool;
            Etag = etag;
            FrontendIPConfigurations = frontendIPConfigurations;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
        }
    }
}
