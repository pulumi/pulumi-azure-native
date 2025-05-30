// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Outputs
{

    /// <summary>
    /// Additional information about every inline event.
    /// </summary>
    [OutputType]
    public sealed class InlineEventPropertiesResponse
    {
        /// <summary>
        /// The dataSchemaUrl for the inline event.
        /// </summary>
        public readonly string? DataSchemaUrl;
        /// <summary>
        /// The description for the inline event.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The displayName for the inline event.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The documentationUrl for the inline event.
        /// </summary>
        public readonly string? DocumentationUrl;

        [OutputConstructor]
        private InlineEventPropertiesResponse(
            string? dataSchemaUrl,

            string? description,

            string? displayName,

            string? documentationUrl)
        {
            DataSchemaUrl = dataSchemaUrl;
            Description = description;
            DisplayName = displayName;
            DocumentationUrl = documentationUrl;
        }
    }
}
