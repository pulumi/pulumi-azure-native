// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240501Preview.Outputs
{

    /// <summary>
    /// The upload status of the virtual hard disk
    /// </summary>
    [OutputType]
    public sealed class VirtualHardDiskUploadStatusResponse
    {
        /// <summary>
        /// The status of Uploading virtual hard disk [Succeeded, Failed, InProgress]
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The uploaded sized of the virtual hard disk in MB
        /// </summary>
        public readonly double? UploadedSizeInMB;

        [OutputConstructor]
        private VirtualHardDiskUploadStatusResponse(
            string status,

            double? uploadedSizeInMB)
        {
            Status = status;
            UploadedSizeInMB = uploadedSizeInMB;
        }
    }
}
