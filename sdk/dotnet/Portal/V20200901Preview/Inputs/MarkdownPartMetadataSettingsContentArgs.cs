// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Portal.V20200901Preview.Inputs
{

    /// <summary>
    /// The content of markdown part.
    /// </summary>
    public sealed class MarkdownPartMetadataSettingsContentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The setting of the content of markdown part.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.MarkdownPartMetadataSettingsContentSettingsArgs>? Settings { get; set; }

        public MarkdownPartMetadataSettingsContentArgs()
        {
        }
        public static new MarkdownPartMetadataSettingsContentArgs Empty => new MarkdownPartMetadataSettingsContentArgs();
    }
}
