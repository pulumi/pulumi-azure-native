// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Storage profile for the directory on the target container.
    /// </summary>
    public sealed class TargetStorageProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure file share profile for hydration of application folders not mounted on
        /// the container file system.
        /// </summary>
        [Input("azureFileShareProfile")]
        public Input<Inputs.AzureFileShareHydrationProfileArgs>? AzureFileShareProfile { get; set; }

        /// <summary>
        /// Gets or sets the storage provider type on the target.
        /// Applicable when StorageProjectionType is not ContainerFileSystem.
        /// </summary>
        [Input("hydrationStorageProviderType")]
        public InputUnion<string, Pulumi.AzureNative.Migrate.TargetHydrationStorageProviderType>? HydrationStorageProviderType { get; set; }

        /// <summary>
        /// Gets or sets the target persistent volume id.
        /// Applicable when StorageProjectionType is PersistentVolume and on using an
        /// existing PersistentVolume.
        /// </summary>
        [Input("persistentVolumeId")]
        public Input<string>? PersistentVolumeId { get; set; }

        /// <summary>
        /// Gets or sets the target storage access type.
        /// </summary>
        [Input("storageAccessType")]
        public InputUnion<string, Pulumi.AzureNative.Migrate.TargetStorageAccessType>? StorageAccessType { get; set; }

        /// <summary>
        /// Gets or sets the target projection type.
        /// </summary>
        [Input("storageProjectionType")]
        public InputUnion<string, Pulumi.AzureNative.Migrate.TargetStorageProjectionType>? StorageProjectionType { get; set; }

        /// <summary>
        /// Gets or sets the name of the projected volume on the target environment.
        /// </summary>
        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        /// <summary>
        /// Gets or sets the storage size on the target.
        /// Applicable when StorageProjectionType is PersistentVolume and on creating a new
        /// PersistentVolume.
        /// </summary>
        [Input("targetSize")]
        public Input<string>? TargetSize { get; set; }

        public TargetStorageProfileArgs()
        {
        }
        public static new TargetStorageProfileArgs Empty => new TargetStorageProfileArgs();
    }
}
