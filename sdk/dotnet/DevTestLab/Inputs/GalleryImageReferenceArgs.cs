// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// The reference information for an Azure Marketplace image.
    /// </summary>
    public sealed class GalleryImageReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The offer of the gallery image.
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        /// <summary>
        /// The OS type of the gallery image.
        /// </summary>
        [Input("osType")]
        public Input<string>? OsType { get; set; }

        /// <summary>
        /// The publisher of the gallery image.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        /// <summary>
        /// The SKU of the gallery image.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        /// <summary>
        /// The version of the gallery image.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GalleryImageReferenceArgs()
        {
        }
        public static new GalleryImageReferenceArgs Empty => new GalleryImageReferenceArgs();
    }
}
