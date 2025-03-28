// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.V20250301Preview.Inputs
{

    /// <summary>
    /// The properties of a static IP Arc Sql availability group listener
    /// </summary>
    public sealed class SqlAvailabilityGroupStaticIPListenerPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the DNS name for the listener.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        [Input("ipV4AddressesAndMasks")]
        private InputList<Inputs.SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs>? _ipV4AddressesAndMasks;

        /// <summary>
        /// IP V4 Addresses and masks for the listener.
        /// </summary>
        public InputList<Inputs.SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs> IpV4AddressesAndMasks
        {
            get => _ipV4AddressesAndMasks ?? (_ipV4AddressesAndMasks = new InputList<Inputs.SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs>());
            set => _ipV4AddressesAndMasks = value;
        }

        [Input("ipV6Addresses")]
        private InputList<string>? _ipV6Addresses;

        /// <summary>
        /// IP V6 Addresses for the listener
        /// </summary>
        public InputList<string> IpV6Addresses
        {
            get => _ipV6Addresses ?? (_ipV6Addresses = new InputList<string>());
            set => _ipV6Addresses = value;
        }

        /// <summary>
        /// Network port for the listener. Default is 1433.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public SqlAvailabilityGroupStaticIPListenerPropertiesArgs()
        {
        }
        public static new SqlAvailabilityGroupStaticIPListenerPropertiesArgs Empty => new SqlAvailabilityGroupStaticIPListenerPropertiesArgs();
    }
}
