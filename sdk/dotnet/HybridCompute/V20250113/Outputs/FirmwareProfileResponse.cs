// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20250113.Outputs
{

    /// <summary>
    /// Describes the firmware of the machine
    /// </summary>
    [OutputType]
    public sealed class FirmwareProfileResponse
    {
        /// <summary>
        /// The serial number of the firmware
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The type of the firmware
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FirmwareProfileResponse(
            string serialNumber,

            string type)
        {
            SerialNumber = serialNumber;
            Type = type;
        }
    }
}
