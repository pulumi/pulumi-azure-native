// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.Inputs
{

    /// <summary>
    /// SKU for the resource.
    /// </summary>
    public sealed class SkuNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SKU name for this resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SkuNameArgs()
        {
        }
        public static new SkuNameArgs Empty => new SkuNameArgs();
    }
}
