// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Outputs
{

    /// <summary>
    /// Image artifact profile.
    /// </summary>
    [OutputType]
    public sealed class ImageArtifactProfileResponse
    {
        /// <summary>
        /// Image name.
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// Image version.
        /// </summary>
        public readonly string? ImageVersion;

        [OutputConstructor]
        private ImageArtifactProfileResponse(
            string? imageName,

            string? imageVersion)
        {
            ImageName = imageName;
            ImageVersion = imageVersion;
        }
    }
}
