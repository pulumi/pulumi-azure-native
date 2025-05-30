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
    /// This is the storage profile of a Gallery Image Version.
    /// </summary>
    [OutputType]
    public sealed class GalleryImageVersionStorageProfileResponse
    {
        /// <summary>
        /// A list of data disk images.
        /// </summary>
        public readonly ImmutableArray<Outputs.GalleryDataDiskImageResponse> DataDiskImages;
        /// <summary>
        /// This is the OS disk image.
        /// </summary>
        public readonly Outputs.GalleryOSDiskImageResponse? OsDiskImage;
        /// <summary>
        /// The source of the gallery artifact version.
        /// </summary>
        public readonly Outputs.GalleryArtifactVersionFullSourceResponse? Source;

        [OutputConstructor]
        private GalleryImageVersionStorageProfileResponse(
            ImmutableArray<Outputs.GalleryDataDiskImageResponse> dataDiskImages,

            Outputs.GalleryOSDiskImageResponse? osDiskImage,

            Outputs.GalleryArtifactVersionFullSourceResponse? source)
        {
            DataDiskImages = dataDiskImages;
            OsDiskImage = osDiskImage;
            Source = source;
        }
    }
}
