// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// List of properties of the device.
    /// </summary>
    [OutputType]
    public sealed class DevicePropertiesResponse
    {
        /// <summary>
        /// Model of the device.
        /// </summary>
        public readonly string? DeviceModel;
        /// <summary>
        /// Name of the device Vendor.
        /// </summary>
        public readonly string? DeviceVendor;
        /// <summary>
        /// Link speed.
        /// </summary>
        public readonly int? LinkSpeedInMbps;

        [OutputConstructor]
        private DevicePropertiesResponse(
            string? deviceModel,

            string? deviceVendor,

            int? linkSpeedInMbps)
        {
            DeviceModel = deviceModel;
            DeviceVendor = deviceVendor;
            LinkSpeedInMbps = linkSpeedInMbps;
        }
    }
}
