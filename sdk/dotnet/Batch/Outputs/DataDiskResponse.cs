// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    /// <summary>
    /// Settings which will be used by the data disks associated to Compute Nodes in the Pool. When using attached data disks, you need to mount and format the disks from within a VM to use them.
    /// </summary>
    [OutputType]
    public sealed class DataDiskResponse
    {
        /// <summary>
        /// Values are:
        /// 
        ///  none - The caching mode for the disk is not enabled.
        ///  readOnly - The caching mode for the disk is read only.
        ///  readWrite - The caching mode for the disk is read and write.
        /// 
        ///  The default value for caching is none. For information about the caching options see: https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
        /// </summary>
        public readonly string? Caching;
        public readonly int DiskSizeGB;
        /// <summary>
        /// The lun is used to uniquely identify each data disk. If attaching multiple disks, each should have a distinct lun. The value must be between 0 and 63, inclusive.
        /// </summary>
        public readonly int Lun;
        /// <summary>
        /// If omitted, the default is "Standard_LRS". Values are:
        /// 
        ///  Standard_LRS - The data disk should use standard locally redundant storage.
        ///  Premium_LRS - The data disk should use premium locally redundant storage.
        /// </summary>
        public readonly string? StorageAccountType;

        [OutputConstructor]
        private DataDiskResponse(
            string? caching,

            int diskSizeGB,

            int lun,

            string? storageAccountType)
        {
            Caching = caching;
            DiskSizeGB = diskSizeGB;
            Lun = lun;
            StorageAccountType = storageAccountType;
        }
    }
}
