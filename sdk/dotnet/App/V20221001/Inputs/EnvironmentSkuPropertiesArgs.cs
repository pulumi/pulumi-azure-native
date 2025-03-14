// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20221001.Inputs
{

    /// <summary>
    /// Managed Environment resource SKU properties.
    /// </summary>
    public sealed class EnvironmentSkuPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Sku.
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNative.App.V20221001.SkuName> Name { get; set; } = null!;

        public EnvironmentSkuPropertiesArgs()
        {
        }
        public static new EnvironmentSkuPropertiesArgs Empty => new EnvironmentSkuPropertiesArgs();
    }
}
