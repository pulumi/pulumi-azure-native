// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.Inputs
{

    /// <summary>
    /// Microsoft.Elastic SKU.
    /// </summary>
    public sealed class ResourceSkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the SKU.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ResourceSkuArgs()
        {
        }
        public static new ResourceSkuArgs Empty => new ResourceSkuArgs();
    }
}
