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
    /// Copy Log Details for a disk
    /// </summary>
    [OutputType]
    public sealed class DataBoxDiskCopyLogDetailsResponse
    {
        /// <summary>
        /// Indicates the type of job details.
        /// Expected value is 'DataBoxDisk'.
        /// </summary>
        public readonly string CopyLogDetailsType;
        /// <summary>
        /// Disk Serial Number.
        /// </summary>
        public readonly string DiskSerialNumber;
        /// <summary>
        /// Link for copy error logs.
        /// </summary>
        public readonly string ErrorLogLink;
        /// <summary>
        /// Link for copy verbose logs.
        /// </summary>
        public readonly string VerboseLogLink;

        [OutputConstructor]
        private DataBoxDiskCopyLogDetailsResponse(
            string copyLogDetailsType,

            string diskSerialNumber,

            string errorLogLink,

            string verboseLogLink)
        {
            CopyLogDetailsType = copyLogDetailsType;
            DiskSerialNumber = diskSerialNumber;
            ErrorLogLink = errorLogLink;
            VerboseLogLink = verboseLogLink;
        }
    }
}
