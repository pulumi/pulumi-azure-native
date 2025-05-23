// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Outputs
{

    /// <summary>
    /// Property overrides on a subnet of a virtual network.
    /// </summary>
    [OutputType]
    public sealed class SubnetOverrideResponse
    {
        /// <summary>
        /// The name given to the subnet within the lab.
        /// </summary>
        public readonly string? LabSubnetName;
        /// <summary>
        /// The resource ID of the subnet.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// Properties that virtual machines on this subnet will share.
        /// </summary>
        public readonly Outputs.SubnetSharedPublicIpAddressConfigurationResponse? SharedPublicIpAddressConfiguration;
        /// <summary>
        /// Indicates whether this subnet can be used during virtual machine creation (i.e. Allow, Deny).
        /// </summary>
        public readonly string? UseInVmCreationPermission;
        /// <summary>
        /// Indicates whether public IP addresses can be assigned to virtual machines on this subnet (i.e. Allow, Deny).
        /// </summary>
        public readonly string? UsePublicIpAddressPermission;
        /// <summary>
        /// The virtual network pool associated with this subnet.
        /// </summary>
        public readonly string? VirtualNetworkPoolName;

        [OutputConstructor]
        private SubnetOverrideResponse(
            string? labSubnetName,

            string? resourceId,

            Outputs.SubnetSharedPublicIpAddressConfigurationResponse? sharedPublicIpAddressConfiguration,

            string? useInVmCreationPermission,

            string? usePublicIpAddressPermission,

            string? virtualNetworkPoolName)
        {
            LabSubnetName = labSubnetName;
            ResourceId = resourceId;
            SharedPublicIpAddressConfiguration = sharedPublicIpAddressConfiguration;
            UseInVmCreationPermission = useInVmCreationPermission;
            UsePublicIpAddressPermission = usePublicIpAddressPermission;
            VirtualNetworkPoolName = virtualNetworkPoolName;
        }
    }
}
