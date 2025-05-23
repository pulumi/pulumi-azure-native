// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    [OutputType]
    public sealed class GalleryTargetExtendedLocationResponse
    {
        /// <summary>
        /// Optional. Allows users to provide customer managed keys for encrypting the OS and data disks in the gallery artifact.
        /// </summary>
        public readonly Outputs.EncryptionImagesResponse? Encryption;
        /// <summary>
        /// The name of the extended location.
        /// </summary>
        public readonly Outputs.GalleryExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// The number of replicas of the Image Version to be created per extended location. This property is updatable.
        /// </summary>
        public readonly int? ExtendedLocationReplicaCount;
        /// <summary>
        /// The name of the region.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Specifies the storage account type to be used to store the image. This property is not updatable.
        /// </summary>
        public readonly string? StorageAccountType;

        [OutputConstructor]
        private GalleryTargetExtendedLocationResponse(
            Outputs.EncryptionImagesResponse? encryption,

            Outputs.GalleryExtendedLocationResponse? extendedLocation,

            int? extendedLocationReplicaCount,

            string? name,

            string? storageAccountType)
        {
            Encryption = encryption;
            ExtendedLocation = extendedLocation;
            ExtendedLocationReplicaCount = extendedLocationReplicaCount;
            Name = name;
            StorageAccountType = storageAccountType;
        }
    }
}
