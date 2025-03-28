// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.V20220801.Outputs
{

    /// <summary>
    /// The additional capabilities for a lab VM.
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineAdditionalCapabilitiesResponse
    {
        /// <summary>
        /// Flag to pre-install dedicated GPU drivers.
        /// </summary>
        public readonly string? InstallGpuDrivers;

        [OutputConstructor]
        private VirtualMachineAdditionalCapabilitiesResponse(string? installGpuDrivers)
        {
            InstallGpuDrivers = installGpuDrivers;
        }
    }
}
