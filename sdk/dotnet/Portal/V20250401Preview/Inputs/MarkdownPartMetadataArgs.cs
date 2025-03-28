// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Portal.V20250401Preview.Inputs
{

    /// <summary>
    /// Markdown part metadata.
    /// </summary>
    public sealed class MarkdownPartMetadataArgs : global::Pulumi.ResourceArgs
    {
        [Input("inputs")]
        private InputList<object>? _inputs;

        /// <summary>
        /// Input to dashboard part.
        /// </summary>
        public InputList<object> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<object>());
            set => _inputs = value;
        }

        /// <summary>
        /// Markdown part settings.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.MarkdownPartMetadataSettingsArgs>? Settings { get; set; }

        /// <summary>
        /// The dashboard part metadata type.
        /// Expected value is 'Extension/HubsExtension/PartType/MarkdownPart'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MarkdownPartMetadataArgs()
        {
        }
        public static new MarkdownPartMetadataArgs Empty => new MarkdownPartMetadataArgs();
    }
}
