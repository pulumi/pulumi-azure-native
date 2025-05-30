// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Inputs
{

    /// <summary>
    /// VnetInfo for Firewall Networking
    /// </summary>
    public sealed class VnetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP of trust subnet for UDR
        /// </summary>
        [Input("ipOfTrustSubnetForUdr")]
        public Input<Inputs.IPAddressArgs>? IpOfTrustSubnetForUdr { get; set; }

        /// <summary>
        /// Trust Subnet
        /// </summary>
        [Input("trustSubnet", required: true)]
        public Input<Inputs.IPAddressSpaceArgs> TrustSubnet { get; set; } = null!;

        /// <summary>
        /// Untrust Subnet
        /// </summary>
        [Input("unTrustSubnet", required: true)]
        public Input<Inputs.IPAddressSpaceArgs> UnTrustSubnet { get; set; } = null!;

        /// <summary>
        /// Azure Virtual Network
        /// </summary>
        [Input("vnet", required: true)]
        public Input<Inputs.IPAddressSpaceArgs> Vnet { get; set; } = null!;

        public VnetConfigurationArgs()
        {
        }
        public static new VnetConfigurationArgs Empty => new VnetConfigurationArgs();
    }
}
