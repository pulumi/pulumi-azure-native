// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevSpaces.Inputs
{

    /// <summary>
    /// Model representing SKU for Azure Dev Spaces Controller.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SKU for Azure Dev Spaces Controller.
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DevSpaces.SkuName> Name { get; set; } = null!;

        /// <summary>
        /// The tier of the SKU for Azure Dev Spaces Controller.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.DevSpaces.SkuTier>? Tier { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
