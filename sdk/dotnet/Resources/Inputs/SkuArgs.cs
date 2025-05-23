// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Resources.Inputs
{

    /// <summary>
    /// SKU for the resource.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SKU capacity.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// The SKU family.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The SKU model.
        /// </summary>
        [Input("model")]
        public Input<string>? Model { get; set; }

        /// <summary>
        /// The SKU name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The SKU size.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The SKU tier.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
