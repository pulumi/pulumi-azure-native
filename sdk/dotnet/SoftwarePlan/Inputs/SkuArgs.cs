// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SoftwarePlan.Inputs
{

    /// <summary>
    /// The SKU to be applied for this resource
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the SKU to be applied
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
