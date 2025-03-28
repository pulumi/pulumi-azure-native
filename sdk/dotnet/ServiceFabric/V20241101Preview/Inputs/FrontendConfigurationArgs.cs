// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20241101Preview.Inputs
{

    /// <summary>
    /// Describes the frontend configurations for the node type.
    /// </summary>
    public sealed class FrontendConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource Id of application gateway backend address pool. The format of the resource Id is '/subscriptions/&lt;subscriptionId&gt;/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendAddressPools/{backendAddressPoolName}'.
        /// </summary>
        [Input("applicationGatewayBackendAddressPoolId")]
        public Input<string>? ApplicationGatewayBackendAddressPoolId { get; set; }

        /// <summary>
        /// The IP address type of this frontend configuration. If omitted the default value is IPv4.
        /// </summary>
        [Input("ipAddressType")]
        public InputUnion<string, Pulumi.AzureNative.ServiceFabric.V20241101Preview.IPAddressType>? IpAddressType { get; set; }

        /// <summary>
        /// The resource Id of the Load Balancer backend address pool that the VM instances of the node type are associated with. The format of the resource Id is '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}'.
        /// </summary>
        [Input("loadBalancerBackendAddressPoolId")]
        public Input<string>? LoadBalancerBackendAddressPoolId { get; set; }

        /// <summary>
        /// The resource Id of the Load Balancer inbound NAT pool that the VM instances of the node type are associated with. The format of the resource Id is '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatPools/{inboundNatPoolName}'.
        /// </summary>
        [Input("loadBalancerInboundNatPoolId")]
        public Input<string>? LoadBalancerInboundNatPoolId { get; set; }

        public FrontendConfigurationArgs()
        {
            IpAddressType = "IPv4";
        }
        public static new FrontendConfigurationArgs Empty => new FrontendConfigurationArgs();
    }
}
