// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Inputs
{

    /// <summary>
    /// Route is associated with a subnet.
    /// </summary>
    public sealed class VirtualNetworkPropertiesRoutesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AddressPrefix - The destination CIDR to which the route applies.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// Name - name of the subnet
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NextHopIPAddress - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Input("nextHopIpAddress")]
        public Input<string>? NextHopIpAddress { get; set; }

        public VirtualNetworkPropertiesRoutesArgs()
        {
        }
        public static new VirtualNetworkPropertiesRoutesArgs Empty => new VirtualNetworkPropertiesRoutesArgs();
    }
}
