// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BotService.Inputs
{

    /// <summary>
    /// The SKU of the cognitive services account.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sku name
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNative.BotService.SkuName> Name { get; set; } = null!;

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
