// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Portal.V20250401Preview.Outputs
{

    /// <summary>
    /// Markdown part settings.
    /// </summary>
    [OutputType]
    public sealed class MarkdownPartMetadataSettingsResponse
    {
        /// <summary>
        /// The content of markdown part.
        /// </summary>
        public readonly Outputs.MarkdownPartMetadataSettingsContentResponse? Content;

        [OutputConstructor]
        private MarkdownPartMetadataSettingsResponse(Outputs.MarkdownPartMetadataSettingsContentResponse? content)
        {
            Content = content;
        }
    }
}
