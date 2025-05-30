// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Automation.Outputs
{

    /// <summary>
    /// Definition of the content link.
    /// </summary>
    [OutputType]
    public sealed class ContentLinkResponse
    {
        /// <summary>
        /// Gets or sets the hash.
        /// </summary>
        public readonly Outputs.ContentHashResponse? ContentHash;
        /// <summary>
        /// Gets or sets the uri of content.
        /// </summary>
        public readonly string? Uri;
        /// <summary>
        /// Gets or sets the version of the content.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ContentLinkResponse(
            Outputs.ContentHashResponse? contentHash,

            string? uri,

            string? version)
        {
            ContentHash = contentHash;
            Uri = uri;
            Version = version;
        }
    }
}
