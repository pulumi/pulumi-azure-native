// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20230601.Outputs
{

    /// <summary>
    /// Network Virtual Appliance NIC properties.
    /// </summary>
    [OutputType]
    public sealed class VirtualApplianceNicPropertiesResponse
    {
        /// <summary>
        /// Instance on which nic is attached.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// NIC name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private IP address.
        /// </summary>
        public readonly string PrivateIpAddress;
        /// <summary>
        /// Public IP address.
        /// </summary>
        public readonly string PublicIpAddress;

        [OutputConstructor]
        private VirtualApplianceNicPropertiesResponse(
            string instanceName,

            string name,

            string privateIpAddress,

            string publicIpAddress)
        {
            InstanceName = instanceName;
            Name = name;
            PrivateIpAddress = privateIpAddress;
            PublicIpAddress = publicIpAddress;
        }
    }
}
