// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Outputs
{

    /// <summary>
    /// Category related properties of a child configuration.
    /// </summary>
    [OutputType]
    public sealed class CategoryInformationResponse
    {
        /// <summary>
        /// Category display name of the child configuration.
        /// </summary>
        public readonly string? CategoryDisplayName;
        /// <summary>
        /// Category name of the child configuration.
        /// </summary>
        public readonly string? CategoryName;
        /// <summary>
        /// Description text for the category.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Links for the category.
        /// </summary>
        public readonly ImmutableArray<Outputs.LinkResponse> Links;

        [OutputConstructor]
        private CategoryInformationResponse(
            string? categoryDisplayName,

            string? categoryName,

            string? description,

            ImmutableArray<Outputs.LinkResponse> links)
        {
            CategoryDisplayName = categoryDisplayName;
            CategoryName = categoryName;
            Description = description;
            Links = links;
        }
    }
}
