// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// A custom action that can be performed with a Gallery Application Version.
    /// </summary>
    [OutputType]
    public sealed class GalleryApplicationCustomActionResponse
    {
        /// <summary>
        /// Description to help the users understand what this custom action does.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the custom action.  Must be unique within the Gallery Application Version.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The parameters that this custom action uses
        /// </summary>
        public readonly ImmutableArray<Outputs.GalleryApplicationCustomActionParameterResponse> Parameters;
        /// <summary>
        /// The script to run when executing this custom action.
        /// </summary>
        public readonly string Script;

        [OutputConstructor]
        private GalleryApplicationCustomActionResponse(
            string? description,

            string name,

            ImmutableArray<Outputs.GalleryApplicationCustomActionParameterResponse> parameters,

            string script)
        {
            Description = description;
            Name = name;
            Parameters = parameters;
            Script = script;
        }
    }
}
