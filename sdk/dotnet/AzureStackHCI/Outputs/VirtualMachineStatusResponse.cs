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
    /// The observed state of virtual machines
    /// </summary>
    [OutputType]
    public sealed class VirtualMachineStatusResponse
    {
        /// <summary>
        /// VirtualMachine provisioning error code
        /// </summary>
        public readonly string? ErrorCode;
        /// <summary>
        /// Descriptive error message
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// The power state of the virtual machine
        /// </summary>
        public readonly string? PowerState;
        public readonly Outputs.VirtualMachineStatusResponseProvisioningStatus? ProvisioningStatus;

        [OutputConstructor]
        private VirtualMachineStatusResponse(
            string? errorCode,

            string? errorMessage,

            string? powerState,

            Outputs.VirtualMachineStatusResponseProvisioningStatus? provisioningStatus)
        {
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            PowerState = powerState;
            ProvisioningStatus = provisioningStatus;
        }
    }
}
