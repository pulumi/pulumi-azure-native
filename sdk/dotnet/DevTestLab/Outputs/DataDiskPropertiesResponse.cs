// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Outputs
{

    /// <summary>
    /// Request body for adding a new or existing data disk to a virtual machine.
    /// </summary>
    [OutputType]
    public sealed class DataDiskPropertiesResponse
    {
        /// <summary>
        /// Specifies options to attach a new disk to the virtual machine.
        /// </summary>
        public readonly Outputs.AttachNewDataDiskOptionsResponse? AttachNewDataDiskOptions;
        /// <summary>
        /// Specifies the existing lab disk id to attach to virtual machine.
        /// </summary>
        public readonly string? ExistingLabDiskId;
        /// <summary>
        /// Caching option for a data disk (i.e. None, ReadOnly, ReadWrite).
        /// </summary>
        public readonly string? HostCaching;

        [OutputConstructor]
        private DataDiskPropertiesResponse(
            Outputs.AttachNewDataDiskOptionsResponse? attachNewDataDiskOptions,

            string? existingLabDiskId,

            string? hostCaching)
        {
            AttachNewDataDiskOptions = attachNewDataDiskOptions;
            ExistingLabDiskId = existingLabDiskId;
            HostCaching = hostCaching;
        }
    }
}
