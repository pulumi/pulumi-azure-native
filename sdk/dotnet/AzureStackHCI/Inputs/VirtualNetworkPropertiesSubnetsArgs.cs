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
    /// Subnet subnet in a virtual network resource.
    /// </summary>
    public sealed class VirtualNetworkPropertiesSubnetsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cidr for this subnet - IPv4, IPv6
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        [Input("addressPrefixes")]
        private InputList<string>? _addressPrefixes;

        /// <summary>
        /// AddressPrefixes - List of address prefixes for the subnet.
        /// </summary>
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        /// <summary>
        /// IPAllocationMethod - The IP address allocation method. Possible values include: 'Static', 'Dynamic'
        /// </summary>
        [Input("ipAllocationMethod")]
        public InputUnion<string, Pulumi.AzureNative.AzureStackHCI.IpAllocationMethodEnum>? IpAllocationMethod { get; set; }

        [Input("ipConfigurationReferences")]
        private InputList<Inputs.VirtualNetworkPropertiesIpConfigurationReferencesArgs>? _ipConfigurationReferences;

        /// <summary>
        /// IPConfigurationReferences - list of IPConfigurationReferences
        /// </summary>
        public InputList<Inputs.VirtualNetworkPropertiesIpConfigurationReferencesArgs> IpConfigurationReferences
        {
            get => _ipConfigurationReferences ?? (_ipConfigurationReferences = new InputList<Inputs.VirtualNetworkPropertiesIpConfigurationReferencesArgs>());
            set => _ipConfigurationReferences = value;
        }

        /// <summary>
        /// Name - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// RouteTable for the subnet
        /// </summary>
        [Input("routeTable")]
        public Input<Inputs.VirtualNetworkPropertiesRouteTableArgs>? RouteTable { get; set; }

        /// <summary>
        /// Vlan to use for the subnet
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public VirtualNetworkPropertiesSubnetsArgs()
        {
        }
        public static new VirtualNetworkPropertiesSubnetsArgs Empty => new VirtualNetworkPropertiesSubnetsArgs();
    }
}
