// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase.Outputs
{

    /// <summary>
    /// Properties of the definition of a gallery application used in Test Base package.
    /// </summary>
    [OutputType]
    public sealed class GalleryAppDefinitionResponse
    {
        /// <summary>
        /// Whether the disclaimer of the gallery application is accepted.
        /// </summary>
        public readonly bool? IsConsented;
        /// <summary>
        /// The SKU id of the gallery application.
        /// </summary>
        public readonly string SkuId;

        [OutputConstructor]
        private GalleryAppDefinitionResponse(
            bool? isConsented,

            string skuId)
        {
            IsConsented = isConsented;
            SkuId = skuId;
        }
    }
}
