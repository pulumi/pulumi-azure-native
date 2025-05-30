// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    [OutputType]
    public sealed class TaskContainerSettingsResponse
    {
        /// <summary>
        /// If this array is null or be not present, container task will mount entire temporary disk drive in windows (or AZ_BATCH_NODE_ROOT_DIR in Linux). It won't' mount any data paths into container if this array is set as empty.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerHostBatchBindMountEntryResponse> ContainerHostBatchBindMounts;
        /// <summary>
        /// These additional options are supplied as arguments to the "docker create" command, in addition to those controlled by the Batch Service.
        /// </summary>
        public readonly string? ContainerRunOptions;
        /// <summary>
        /// This is the full image reference, as would be specified to "docker pull". If no tag is provided as part of the image name, the tag ":latest" is used as a default.
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// This setting can be omitted if was already provided at pool creation.
        /// </summary>
        public readonly Outputs.ContainerRegistryResponse? Registry;
        public readonly string? WorkingDirectory;

        [OutputConstructor]
        private TaskContainerSettingsResponse(
            ImmutableArray<Outputs.ContainerHostBatchBindMountEntryResponse> containerHostBatchBindMounts,

            string? containerRunOptions,

            string imageName,

            Outputs.ContainerRegistryResponse? registry,

            string? workingDirectory)
        {
            ContainerHostBatchBindMounts = containerHostBatchBindMounts;
            ContainerRunOptions = containerRunOptions;
            ImageName = imageName;
            Registry = registry;
            WorkingDirectory = workingDirectory;
        }
    }
}
