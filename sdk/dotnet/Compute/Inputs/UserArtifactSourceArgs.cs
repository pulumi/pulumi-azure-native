// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Inputs
{

    /// <summary>
    /// The source image from which the Image Version is going to be created.
    /// </summary>
    public sealed class UserArtifactSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The defaultConfigurationLink of the artifact, must be a readable storage page blob.
        /// </summary>
        [Input("defaultConfigurationLink")]
        public Input<string>? DefaultConfigurationLink { get; set; }

        /// <summary>
        /// Required. The mediaLink of the artifact, must be a readable storage page blob.
        /// </summary>
        [Input("mediaLink", required: true)]
        public Input<string> MediaLink { get; set; } = null!;

        public UserArtifactSourceArgs()
        {
        }
        public static new UserArtifactSourceArgs Empty => new UserArtifactSourceArgs();
    }
}
