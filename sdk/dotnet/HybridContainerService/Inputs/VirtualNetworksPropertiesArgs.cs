// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Inputs
{

    /// <summary>
    /// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
    /// </summary>
    public sealed class VirtualNetworksPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// Address of the DNS servers associated with the network
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Address of the Gateway associated with the network
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("infraVnetProfile")]
        public Input<Inputs.VirtualNetworksPropertiesInfraVnetProfileArgs>? InfraVnetProfile { get; set; }

        /// <summary>
        /// IP Address Prefix of the network
        /// </summary>
        [Input("ipAddressPrefix")]
        public Input<string>? IpAddressPrefix { get; set; }

        [Input("vipPool")]
        private InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs>? _vipPool;

        /// <summary>
        /// Virtual IP Pool for Kubernetes
        /// </summary>
        public InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs> VipPool
        {
            get => _vipPool ?? (_vipPool = new InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs>());
            set => _vipPool = value;
        }

        [Input("vmipPool")]
        private InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs>? _vmipPool;

        /// <summary>
        /// IP Pool for Virtual Machines
        /// </summary>
        public InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs> VmipPool
        {
            get => _vmipPool ?? (_vmipPool = new InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs>());
            set => _vmipPool = value;
        }

        public VirtualNetworksPropertiesArgs()
        {
        }
        public static new VirtualNetworksPropertiesArgs Empty => new VirtualNetworksPropertiesArgs();
    }
}
