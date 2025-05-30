// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Outputs
{

    /// <summary>
    /// Properties of an artifact parameter.
    /// </summary>
    [OutputType]
    public sealed class ArtifactParameterPropertiesResponse
    {
        /// <summary>
        /// The name of the artifact parameter.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value of the artifact parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ArtifactParameterPropertiesResponse(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
