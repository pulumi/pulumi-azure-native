// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240101.Outputs
{

    /// <summary>
    /// NetworkProfile - describes the network configuration the virtual machine instance
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineInstancePropertiesResponseNetworkProfile
    {
        /// <summary>
        /// NetworkInterfaces - list of network interfaces to be attached to the virtual machine instance
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineInstancePropertiesResponseNetworkInterfaces> NetworkInterfaces;

        [OutputConstructor]
        private VirtualMachineInstancePropertiesResponseNetworkProfile(ImmutableArray<Outputs.VirtualMachineInstancePropertiesResponseNetworkInterfaces> networkInterfaces)
        {
            NetworkInterfaces = networkInterfaces;
        }
    }
}
