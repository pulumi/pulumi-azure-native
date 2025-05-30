// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridNetwork.Inputs
{

    /// <summary>
    /// Manifest artifact properties.
    /// </summary>
    public sealed class ManifestArtifactFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The artifact name
        /// </summary>
        [Input("artifactName")]
        public Input<string>? ArtifactName { get; set; }

        /// <summary>
        /// The artifact type.
        /// </summary>
        [Input("artifactType")]
        public InputUnion<string, Pulumi.AzureNative.HybridNetwork.ArtifactType>? ArtifactType { get; set; }

        /// <summary>
        /// The artifact version.
        /// </summary>
        [Input("artifactVersion")]
        public Input<string>? ArtifactVersion { get; set; }

        public ManifestArtifactFormatArgs()
        {
        }
        public static new ManifestArtifactFormatArgs Empty => new ManifestArtifactFormatArgs();
    }
}
