// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20230901Preview.Outputs
{

    /// <summary>
    /// The observed state of gallery images
    /// </summary>
    [OutputType]
    public sealed class GalleryImageStatusResponse
    {
        /// <summary>
        /// The download status of the gallery image
        /// </summary>
        public readonly Outputs.GalleryImageStatusResponseDownloadStatus? DownloadStatus;
        /// <summary>
        /// GalleryImage provisioning error code
        /// </summary>
        public readonly string? ErrorCode;
        /// <summary>
        /// Descriptive error message
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// The progress of the operation in percentage
        /// </summary>
        public readonly double? ProgressPercentage;
        public readonly Outputs.GalleryImageStatusResponseProvisioningStatus? ProvisioningStatus;

        [OutputConstructor]
        private GalleryImageStatusResponse(
            Outputs.GalleryImageStatusResponseDownloadStatus? downloadStatus,

            string? errorCode,

            string? errorMessage,

            double? progressPercentage,

            Outputs.GalleryImageStatusResponseProvisioningStatus? provisioningStatus)
        {
            DownloadStatus = downloadStatus;
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            ProgressPercentage = progressPercentage;
            ProvisioningStatus = provisioningStatus;
        }
    }
}
