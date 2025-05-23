// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Outputs
{

    /// <summary>
    /// DataBox Disk Copy Progress
    /// </summary>
    [OutputType]
    public sealed class DataBoxDiskCopyProgressResponse
    {
        /// <summary>
        /// Available actions on the job.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        /// <summary>
        /// Bytes copied during the copy of disk.
        /// </summary>
        public readonly double BytesCopied;
        /// <summary>
        /// Error, if any, in the stage
        /// </summary>
        public readonly Outputs.CloudErrorResponse Error;
        /// <summary>
        /// Indicates the percentage completed for the copy of the disk.
        /// </summary>
        public readonly int PercentComplete;
        /// <summary>
        /// The serial number of the disk
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The Status of the copy
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private DataBoxDiskCopyProgressResponse(
            ImmutableArray<string> actions,

            double bytesCopied,

            Outputs.CloudErrorResponse error,

            int percentComplete,

            string serialNumber,

            string status)
        {
            Actions = actions;
            BytesCopied = bytesCopied;
            Error = error;
            PercentComplete = percentComplete;
            SerialNumber = serialNumber;
            Status = status;
        }
    }
}
