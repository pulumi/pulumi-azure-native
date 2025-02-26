// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20210901Preview.Inputs
{

    /// <summary>
    /// NetworkProfile - describes the network configuration the virtual machine
    /// </summary>
    public sealed class VirtualmachinesPropertiesNetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("networkInterfaces")]
        private InputList<Inputs.VirtualmachinesPropertiesNetworkInterfacesArgs>? _networkInterfaces;

        /// <summary>
        /// NetworkInterfaces - list of network interfaces to be attached to the virtual machine
        /// </summary>
        public InputList<Inputs.VirtualmachinesPropertiesNetworkInterfacesArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.VirtualmachinesPropertiesNetworkInterfacesArgs>());
            set => _networkInterfaces = value;
        }

        public VirtualmachinesPropertiesNetworkProfileArgs()
        {
        }
        public static new VirtualmachinesPropertiesNetworkProfileArgs Empty => new VirtualmachinesPropertiesNetworkProfileArgs();
    }
}
