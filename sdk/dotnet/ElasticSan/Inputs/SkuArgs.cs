// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ElasticSan.Inputs
{

    /// <summary>
    /// The SKU name. Required for account creation; optional for update.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ElasticSan.SkuName> Name { get; set; } = null!;

        /// <summary>
        /// The sku tier.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.ElasticSan.SkuTier>? Tier { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
