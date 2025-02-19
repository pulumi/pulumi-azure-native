// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20210901Preview.Outputs
{

    /// <summary>
    /// VirtualHardDiskStatus defines the observed state of virtualharddisks
    /// </summary>
    [OutputType]
    public sealed class VirtualHardDiskStatusResponse
    {
        /// <summary>
        /// VirtualHardDisk provisioning error code
        /// </summary>
        public readonly string? ErrorCode;
        /// <summary>
        /// Descriptive error message
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// The provisioning status of the virtual hard disk
        /// </summary>
        public readonly Outputs.VirtualHardDiskStatusResponseProvisioningStatus? ProvisioningStatus;

        [OutputConstructor]
        private VirtualHardDiskStatusResponse(
            string? errorCode,

            string? errorMessage,

            Outputs.VirtualHardDiskStatusResponseProvisioningStatus? provisioningStatus)
        {
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            ProvisioningStatus = provisioningStatus;
        }
    }
}
