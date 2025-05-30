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
    /// The arm template RE.
    /// </summary>
    [OutputType]
    public sealed class ArmResourceDefinitionResourceElementTemplateResponse
    {
        /// <summary>
        /// Artifact profile properties.
        /// </summary>
        public readonly Outputs.NSDArtifactProfileResponse? ArtifactProfile;
        /// <summary>
        /// Name and value pairs that define the parameter values. It can be  a well formed escaped JSON string.
        /// </summary>
        public readonly string? ParameterValues;
        /// <summary>
        /// The template type.
        /// </summary>
        public readonly string? TemplateType;

        [OutputConstructor]
        private ArmResourceDefinitionResourceElementTemplateResponse(
            Outputs.NSDArtifactProfileResponse? artifactProfile,

            string? parameterValues,

            string? templateType)
        {
            ArtifactProfile = artifactProfile;
            ParameterValues = parameterValues;
            TemplateType = templateType;
        }
    }
}
