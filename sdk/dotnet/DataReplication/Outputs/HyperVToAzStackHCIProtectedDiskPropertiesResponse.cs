// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Outputs
{

    /// <summary>
    /// HyperVToAzStackHCI protected disk properties.
    /// </summary>
    [OutputType]
    public sealed class HyperVToAzStackHCIProtectedDiskPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the disk capacity in bytes.
        /// </summary>
        public readonly double CapacityInBytes;
        /// <summary>
        /// Gets or sets the disk type.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// Gets or sets a value indicating whether dynamic sizing is enabled on the virtual hard
        /// disk.
        /// </summary>
        public readonly bool IsDynamic;
        /// <summary>
        /// Gets or sets a value indicating whether the disk is the OS disk.
        /// </summary>
        public readonly bool IsOsDisk;
        /// <summary>
        /// Gets or sets the failover clone disk.
        /// </summary>
        public readonly string MigrateDiskName;
        /// <summary>
        /// Gets or sets the seed disk name.
        /// </summary>
        public readonly string SeedDiskName;
        /// <summary>
        /// Gets or sets the source disk Id.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// Gets or sets the source disk Name.
        /// </summary>
        public readonly string SourceDiskName;
        /// <summary>
        /// Gets or sets the ARM Id of the storage container.
        /// </summary>
        public readonly string StorageContainerId;
        /// <summary>
        /// Gets or sets the local path of the storage container.
        /// </summary>
        public readonly string StorageContainerLocalPath;
        /// <summary>
        /// Gets or sets the test failover clone disk.
        /// </summary>
        public readonly string TestMigrateDiskName;

        [OutputConstructor]
        private HyperVToAzStackHCIProtectedDiskPropertiesResponse(
            double capacityInBytes,

            string diskType,

            bool isDynamic,

            bool isOsDisk,

            string migrateDiskName,

            string seedDiskName,

            string sourceDiskId,

            string sourceDiskName,

            string storageContainerId,

            string storageContainerLocalPath,

            string testMigrateDiskName)
        {
            CapacityInBytes = capacityInBytes;
            DiskType = diskType;
            IsDynamic = isDynamic;
            IsOsDisk = isOsDisk;
            MigrateDiskName = migrateDiskName;
            SeedDiskName = seedDiskName;
            SourceDiskId = sourceDiskId;
            SourceDiskName = sourceDiskName;
            StorageContainerId = storageContainerId;
            StorageContainerLocalPath = storageContainerLocalPath;
            TestMigrateDiskName = testMigrateDiskName;
        }
    }
}
