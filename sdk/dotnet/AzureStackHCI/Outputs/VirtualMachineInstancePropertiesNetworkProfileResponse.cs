// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.Outputs
{

    /// <summary>
    /// NetworkProfile - describes the network configuration the virtual machine instance
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineInstancePropertiesNetworkProfileResponse
    {
        /// <summary>
        /// NetworkInterfaces - list of network interfaces to be attached to the virtual machine instance
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceArmReferenceResponse> NetworkInterfaces;

        [OutputConstructor]
        private VirtualMachineInstancePropertiesNetworkProfileResponse(ImmutableArray<Outputs.NetworkInterfaceArmReferenceResponse> networkInterfaces)
        {
            NetworkInterfaces = networkInterfaces;
        }
    }
}
