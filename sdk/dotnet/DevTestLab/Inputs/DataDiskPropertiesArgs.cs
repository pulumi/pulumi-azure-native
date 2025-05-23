// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// Request body for adding a new or existing data disk to a virtual machine.
    /// </summary>
    public sealed class DataDiskPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies options to attach a new disk to the virtual machine.
        /// </summary>
        [Input("attachNewDataDiskOptions")]
        public Input<Inputs.AttachNewDataDiskOptionsArgs>? AttachNewDataDiskOptions { get; set; }

        /// <summary>
        /// Specifies the existing lab disk id to attach to virtual machine.
        /// </summary>
        [Input("existingLabDiskId")]
        public Input<string>? ExistingLabDiskId { get; set; }

        /// <summary>
        /// Caching option for a data disk (i.e. None, ReadOnly, ReadWrite).
        /// </summary>
        [Input("hostCaching")]
        public InputUnion<string, Pulumi.AzureNative.DevTestLab.HostCachingOptions>? HostCaching { get; set; }

        public DataDiskPropertiesArgs()
        {
        }
        public static new DataDiskPropertiesArgs Empty => new DataDiskPropertiesArgs();
    }
}
