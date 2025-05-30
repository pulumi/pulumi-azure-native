// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StreamAnalytics.Inputs
{

    /// <summary>
    /// The properties that are associated with a SKU.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the SKU. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.StreamAnalytics.SkuName>? Name { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
