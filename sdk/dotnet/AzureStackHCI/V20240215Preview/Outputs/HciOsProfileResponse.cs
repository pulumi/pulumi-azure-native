// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240215Preview.Outputs
{

    /// <summary>
    /// OS configurations for HCI device.
    /// </summary>
    [OutputType]
    public sealed class HciOsProfileResponse
    {
        /// <summary>
        /// The boot type of the device. e.g. UEFI, Legacy etc
        /// </summary>
        public readonly string BootType;

        [OutputConstructor]
        private HciOsProfileResponse(string bootType)
        {
            BootType = bootType;
        }
    }
}
