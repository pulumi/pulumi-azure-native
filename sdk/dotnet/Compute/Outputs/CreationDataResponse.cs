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
    /// Data used when creating a disk.
    /// </summary>
    [OutputType]
    public sealed class CreationDataResponse
    {
        /// <summary>
        /// This enumerates the possible sources of a disk's creation.
        /// </summary>
        public readonly string CreateOption;
        /// <summary>
        /// Required if createOption is CopyFromSanSnapshot. This is the ARM id of the source elastic san volume snapshot.
        /// </summary>
        public readonly string? ElasticSanResourceId;
        /// <summary>
        /// Required if creating from a Gallery Image. The id/sharedGalleryImageId/communityGalleryImageId of the ImageDiskReference will be the ARM id of the shared galley image version from which to create a disk.
        /// </summary>
        public readonly Outputs.ImageDiskReferenceResponse? GalleryImageReference;
        /// <summary>
        /// Disk source information for PIR or user images.
        /// </summary>
        public readonly Outputs.ImageDiskReferenceResponse? ImageReference;
        /// <summary>
        /// Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
        /// </summary>
        public readonly int? LogicalSectorSize;
        /// <summary>
        /// Set this flag to true to get a boost on the performance target of the disk deployed, see here on the respective performance target. This flag can only be set on disk creation time and cannot be disabled after enabled.
        /// </summary>
        public readonly bool? PerformancePlus;
        /// <summary>
        /// If this field is set on a snapshot and createOption is CopyStart, the snapshot will be copied at a quicker speed.
        /// </summary>
        public readonly string? ProvisionedBandwidthCopySpeed;
        /// <summary>
        /// If createOption is ImportSecure, this is the URI of a blob to be imported into VM guest state.
        /// </summary>
        public readonly string? SecurityDataUri;
        /// <summary>
        /// If createOption is Copy, this is the ARM id of the source snapshot or disk.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// If this field is set, this is the unique id identifying the source of this resource.
        /// </summary>
        public readonly string SourceUniqueId;
        /// <summary>
        /// If createOption is Import, this is the URI of a blob to be imported into a managed disk.
        /// </summary>
        public readonly string? SourceUri;
        /// <summary>
        /// Required if createOption is Import. The Azure Resource Manager identifier of the storage account containing the blob to import as a disk.
        /// </summary>
        public readonly string? StorageAccountId;
        /// <summary>
        /// If createOption is Upload, this is the size of the contents of the upload including the VHD footer. This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer).
        /// </summary>
        public readonly double? UploadSizeBytes;

        [OutputConstructor]
        private CreationDataResponse(
            string createOption,

            string? elasticSanResourceId,

            Outputs.ImageDiskReferenceResponse? galleryImageReference,

            Outputs.ImageDiskReferenceResponse? imageReference,

            int? logicalSectorSize,

            bool? performancePlus,

            string? provisionedBandwidthCopySpeed,

            string? securityDataUri,

            string? sourceResourceId,

            string sourceUniqueId,

            string? sourceUri,

            string? storageAccountId,

            double? uploadSizeBytes)
        {
            CreateOption = createOption;
            ElasticSanResourceId = elasticSanResourceId;
            GalleryImageReference = galleryImageReference;
            ImageReference = imageReference;
            LogicalSectorSize = logicalSectorSize;
            PerformancePlus = performancePlus;
            ProvisionedBandwidthCopySpeed = provisionedBandwidthCopySpeed;
            SecurityDataUri = securityDataUri;
            SourceResourceId = sourceResourceId;
            SourceUniqueId = sourceUniqueId;
            SourceUri = sourceUri;
            StorageAccountId = storageAccountId;
            UploadSizeBytes = uploadSizeBytes;
        }
    }
}
