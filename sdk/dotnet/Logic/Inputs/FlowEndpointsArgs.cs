// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.Inputs
{

    /// <summary>
    /// The flow endpoints configuration.
    /// </summary>
    public sealed class FlowEndpointsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessEndpointIpAddresses")]
        private InputList<Inputs.IpAddressArgs>? _accessEndpointIpAddresses;

        /// <summary>
        /// The access endpoint ip address.
        /// </summary>
        public InputList<Inputs.IpAddressArgs> AccessEndpointIpAddresses
        {
            get => _accessEndpointIpAddresses ?? (_accessEndpointIpAddresses = new InputList<Inputs.IpAddressArgs>());
            set => _accessEndpointIpAddresses = value;
        }

        [Input("outgoingIpAddresses")]
        private InputList<Inputs.IpAddressArgs>? _outgoingIpAddresses;

        /// <summary>
        /// The outgoing ip address.
        /// </summary>
        public InputList<Inputs.IpAddressArgs> OutgoingIpAddresses
        {
            get => _outgoingIpAddresses ?? (_outgoingIpAddresses = new InputList<Inputs.IpAddressArgs>());
            set => _outgoingIpAddresses = value;
        }

        public FlowEndpointsArgs()
        {
        }
        public static new FlowEndpointsArgs Empty => new FlowEndpointsArgs();
    }
}
