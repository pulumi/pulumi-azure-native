// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase.Inputs
{

    /// <summary>
    /// Properties of the definition of a gallery application used in Test Base package.
    /// </summary>
    public sealed class GalleryAppDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the disclaimer of the gallery application is accepted.
        /// </summary>
        [Input("isConsented")]
        public Input<bool>? IsConsented { get; set; }

        /// <summary>
        /// The SKU id of the gallery application.
        /// </summary>
        [Input("skuId", required: true)]
        public Input<string> SkuId { get; set; } = null!;

        public GalleryAppDefinitionArgs()
        {
            IsConsented = false;
        }
        public static new GalleryAppDefinitionArgs Empty => new GalleryAppDefinitionArgs();
    }
}
