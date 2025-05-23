// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StoragePool.Outputs
{

    /// <summary>
    /// LUN to expose the Azure Managed Disk.
    /// </summary>
    [OutputType]
    public sealed class IscsiLunResponse
    {
        /// <summary>
        /// Specifies the Logical Unit Number of the iSCSI LUN.
        /// </summary>
        public readonly int Lun;
        /// <summary>
        /// Azure Resource ID of the Managed Disk.
        /// </summary>
        public readonly string ManagedDiskAzureResourceId;
        /// <summary>
        /// User defined name for iSCSI LUN; example: "lun0"
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private IscsiLunResponse(
            int lun,

            string managedDiskAzureResourceId,

            string name)
        {
            Lun = lun;
            ManagedDiskAzureResourceId = managedDiskAzureResourceId;
            Name = name;
        }
    }
}
