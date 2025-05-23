// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.Outputs
{

    /// <summary>
    /// Distribute as a Managed Disk Image.
    /// </summary>
    [OutputType]
    public sealed class ImageTemplateManagedImageDistributorResponse
    {
        /// <summary>
        /// Tags that will be applied to the artifact once it has been created/updated by the distributor.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ArtifactTags;
        /// <summary>
        /// Resource Id of the Managed Disk Image
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Azure location for the image, should match if image already exists
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name to be used for the associated RunOutput.
        /// </summary>
        public readonly string RunOutputName;
        /// <summary>
        /// Type of distribution.
        /// Expected value is 'ManagedImage'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ImageTemplateManagedImageDistributorResponse(
            ImmutableDictionary<string, string>? artifactTags,

            string imageId,

            string location,

            string runOutputName,

            string type)
        {
            ArtifactTags = artifactTags;
            ImageId = imageId;
            Location = location;
            RunOutputName = runOutputName;
            Type = type;
        }
    }
}
