// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// Specifies the storage settings for the virtual machine disks.
    /// </summary>
    [OutputType]
    public sealed class StorageProfileResponse
    {
        /// <summary>
        /// Specifies whether the regional disks should be aligned/moved to the VM zone. This is applicable only for VMs with placement property set. Please note that this change is irreversible. Minimum api-version: 2024-11-01.
        /// </summary>
        public readonly bool? AlignRegionalDisksToVMZone;
        /// <summary>
        /// Specifies the parameters that are used to add a data disk to a virtual machine. For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
        /// </summary>
        public readonly ImmutableArray<Outputs.DataDiskResponse> DataDisks;
        /// <summary>
        /// Specifies the disk controller type configured for the VM. **Note:** This property will be set to the default disk controller type if not specified provided virtual machine is being created with 'hyperVGeneration' set to V2 based on the capabilities of the operating system disk and VM size from the the specified minimum api version. You need to deallocate the VM before updating its disk controller type unless you are updating the VM size in the VM configuration which implicitly deallocates and reallocates the VM. Minimum api-version: 2022-08-01.
        /// </summary>
        public readonly string? DiskControllerType;
        /// <summary>
        /// Specifies information about the image to use. You can specify information about platform images, marketplace images, or virtual machine images. This element is required when you want to use a platform image, marketplace image, or virtual machine image, but is not used in other creation operations.
        /// </summary>
        public readonly Outputs.ImageReferenceResponse? ImageReference;
        /// <summary>
        /// Specifies information about the operating system disk used by the virtual machine. For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/managed-disks-overview).
        /// </summary>
        public readonly Outputs.OSDiskResponse? OsDisk;

        [OutputConstructor]
        private StorageProfileResponse(
            bool? alignRegionalDisksToVMZone,

            ImmutableArray<Outputs.DataDiskResponse> dataDisks,

            string? diskControllerType,

            Outputs.ImageReferenceResponse? imageReference,

            Outputs.OSDiskResponse? osDisk)
        {
            AlignRegionalDisksToVMZone = alignRegionalDisksToVMZone;
            DataDisks = dataDisks;
            DiskControllerType = diskControllerType;
            ImageReference = imageReference;
            OsDisk = osDisk;
        }
    }
}
