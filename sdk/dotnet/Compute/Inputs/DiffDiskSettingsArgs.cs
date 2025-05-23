// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// Describes the parameters of ephemeral disk settings that can be specified for operating system disk. **Note:** The ephemeral disk settings can only be specified for managed disk.
    /// </summary>
    public sealed class DiffDiskSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ephemeral disk settings for operating system disk.
        /// </summary>
        [Input("option")]
        public InputUnion<string, Pulumi.AzureNative.Compute.DiffDiskOptions>? Option { get; set; }

        /// <summary>
        /// Specifies the ephemeral disk placement for operating system disk. Possible values are: **CacheDisk,** **ResourceDisk,** **NvmeDisk.** The defaulting behavior is: **CacheDisk** if one is configured for the VM size otherwise **ResourceDisk** or **NvmeDisk** is used. Refer to the VM size documentation for Windows VM at https://docs.microsoft.com/azure/virtual-machines/windows/sizes and Linux VM at https://docs.microsoft.com/azure/virtual-machines/linux/sizes to check which VM sizes exposes a cache disk. Minimum api-version for NvmeDisk: 2024-03-01.
        /// </summary>
        [Input("placement")]
        public InputUnion<string, Pulumi.AzureNative.Compute.DiffDiskPlacement>? Placement { get; set; }

        public DiffDiskSettingsArgs()
        {
        }
        public static new DiffDiskSettingsArgs Empty => new DiffDiskSettingsArgs();
    }
}
