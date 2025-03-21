// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureLargeInstance.V20240801Preview.Outputs
{

    /// <summary>
    /// Specifies the hardware settings for the Azure Large Instance.
    /// </summary>
    [OutputType]
    public sealed class HardwareProfileResponse
    {
        /// <summary>
        /// Specifies the Azure Large Instance SKU.
        /// </summary>
        public readonly string? AzureLargeInstanceSize;
        /// <summary>
        /// Name of the hardware type (vendor and/or their product name)
        /// </summary>
        public readonly string? HardwareType;

        [OutputConstructor]
        private HardwareProfileResponse(
            string? azureLargeInstanceSize,

            string? hardwareType)
        {
            AzureLargeInstanceSize = azureLargeInstanceSize;
            HardwareType = hardwareType;
        }
    }
}
