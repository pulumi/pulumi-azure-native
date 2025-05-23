// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// Single topology parameter declaration. Declared parameters can and must be referenced throughout the topology and can optionally have default values to be used when they are not defined in the pipelines.
    /// </summary>
    [OutputType]
    public sealed class ParameterDeclarationResponse
    {
        /// <summary>
        /// The default value for the parameter to be used if the pipeline does not specify a value.
        /// </summary>
        public readonly string? Default;
        /// <summary>
        /// Description of the parameter.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Name of the parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the parameter.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ParameterDeclarationResponse(
            string? @default,

            string? description,

            string name,

            string type)
        {
            Default = @default;
            Description = description;
            Name = name;
            Type = type;
        }
    }
}
