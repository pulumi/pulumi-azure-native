// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Outputs
{

    /// <summary>
    /// The image update trigger that caused a build.
    /// </summary>
    [OutputType]
    public sealed class ImageUpdateTriggerResponse
    {
        /// <summary>
        /// The unique ID of the trigger.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The list of image updates that caused the build.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageDescriptorResponse> Images;
        /// <summary>
        /// The timestamp when the image update happened.
        /// </summary>
        public readonly string? Timestamp;

        [OutputConstructor]
        private ImageUpdateTriggerResponse(
            string? id,

            ImmutableArray<Outputs.ImageDescriptorResponse> images,

            string? timestamp)
        {
            Id = id;
            Images = images;
            Timestamp = timestamp;
        }
    }
}
