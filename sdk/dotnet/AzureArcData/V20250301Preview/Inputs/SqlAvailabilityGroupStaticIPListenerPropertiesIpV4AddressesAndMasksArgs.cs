// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.V20250301Preview.Inputs
{

    public sealed class SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPV4 address
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// IPV4 netmask
        /// </summary>
        [Input("mask")]
        public Input<string>? Mask { get; set; }

        public SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs()
        {
        }
        public static new SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs Empty => new SqlAvailabilityGroupStaticIPListenerPropertiesIpV4AddressesAndMasksArgs();
    }
}
