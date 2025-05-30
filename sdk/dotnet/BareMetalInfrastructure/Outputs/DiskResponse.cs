// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BareMetalInfrastructure.Outputs
{

    /// <summary>
    /// Specifies the disk information fo the Azure Bare Metal Instance
    /// </summary>
    [OutputType]
    public sealed class DiskResponse
    {
        /// <summary>
        /// Specifies the size of an empty data disk in gigabytes.
        /// </summary>
        public readonly int? DiskSizeGB;
        /// <summary>
        /// Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
        /// </summary>
        public readonly int Lun;
        /// <summary>
        /// The disk name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private DiskResponse(
            int? diskSizeGB,

            int lun,

            string? name)
        {
            DiskSizeGB = diskSizeGB;
            Lun = lun;
            Name = name;
        }
    }
}
