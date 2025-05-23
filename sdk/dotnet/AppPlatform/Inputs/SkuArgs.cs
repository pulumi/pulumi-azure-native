// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Inputs
{

    /// <summary>
    /// Sku of Azure Spring Apps
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current capacity of the target resource
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Name of the Sku
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Tier of the Sku
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SkuArgs()
        {
            Name = "S0";
            Tier = "Standard";
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
