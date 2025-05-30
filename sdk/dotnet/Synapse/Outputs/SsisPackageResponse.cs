// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.Outputs
{

    /// <summary>
    /// Ssis Package.
    /// </summary>
    [OutputType]
    public sealed class SsisPackageResponse
    {
        /// <summary>
        /// Metadata description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Folder id which contains package.
        /// </summary>
        public readonly double? FolderId;
        /// <summary>
        /// Metadata id.
        /// </summary>
        public readonly double? Id;
        /// <summary>
        /// Metadata name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Parameters in package
        /// </summary>
        public readonly ImmutableArray<Outputs.SsisParameterResponse> Parameters;
        /// <summary>
        /// Project id which contains package.
        /// </summary>
        public readonly double? ProjectId;
        /// <summary>
        /// Project version which contains package.
        /// </summary>
        public readonly double? ProjectVersion;
        /// <summary>
        /// The type of SSIS object metadata.
        /// Expected value is 'Package'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SsisPackageResponse(
            string? description,

            double? folderId,

            double? id,

            string? name,

            ImmutableArray<Outputs.SsisParameterResponse> parameters,

            double? projectId,

            double? projectVersion,

            string type)
        {
            Description = description;
            FolderId = folderId;
            Id = id;
            Name = name;
            Parameters = parameters;
            ProjectId = projectId;
            ProjectVersion = projectVersion;
            Type = type;
        }
    }
}
