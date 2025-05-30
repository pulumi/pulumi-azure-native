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
    public sealed class ContainerConfigurationResponse
    {
        /// <summary>
        /// This is the full image reference, as would be specified to "docker pull". An image will be sourced from the default Docker registry unless the image is fully qualified with an alternative registry.
        /// </summary>
        public readonly ImmutableArray<string> ContainerImageNames;
        /// <summary>
        /// If any images must be downloaded from a private registry which requires credentials, then those credentials must be provided here.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerRegistryResponse> ContainerRegistries;
        public readonly string Type;

        [OutputConstructor]
        private ContainerConfigurationResponse(
            ImmutableArray<string> containerImageNames,

            ImmutableArray<Outputs.ContainerRegistryResponse> containerRegistries,

            string type)
        {
            ContainerImageNames = containerImageNames;
            ContainerRegistries = containerRegistries;
            Type = type;
        }
    }
}
