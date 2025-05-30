// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cloudngfw.Outputs
{

    /// <summary>
    /// VwanInfo for Firewall Networking
    /// </summary>
    [OutputType]
    public sealed class VwanConfigurationResponse
    {
        /// <summary>
        /// IP of trust subnet for UDR
        /// </summary>
        public readonly Outputs.IPAddressResponse? IpOfTrustSubnetForUdr;
        /// <summary>
        /// Network Virtual Appliance resource ID 
        /// </summary>
        public readonly string? NetworkVirtualApplianceId;
        /// <summary>
        /// Trust Subnet
        /// </summary>
        public readonly Outputs.IPAddressSpaceResponse? TrustSubnet;
        /// <summary>
        /// Untrust Subnet
        /// </summary>
        public readonly Outputs.IPAddressSpaceResponse? UnTrustSubnet;
        /// <summary>
        /// vHub Address
        /// </summary>
        public readonly Outputs.IPAddressSpaceResponse VHub;

        [OutputConstructor]
        private VwanConfigurationResponse(
            Outputs.IPAddressResponse? ipOfTrustSubnetForUdr,

            string? networkVirtualApplianceId,

            Outputs.IPAddressSpaceResponse? trustSubnet,

            Outputs.IPAddressSpaceResponse? unTrustSubnet,

            Outputs.IPAddressSpaceResponse vHub)
        {
            IpOfTrustSubnetForUdr = ipOfTrustSubnetForUdr;
            NetworkVirtualApplianceId = networkVirtualApplianceId;
            TrustSubnet = trustSubnet;
            UnTrustSubnet = unTrustSubnet;
            VHub = vHub;
        }
    }
}
