// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Outputs
{

    /// <summary>
    /// SKU for the resource.
    /// </summary>
    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// The SKU name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The SKU tier.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private SkuResponse(
            string name,

            string? tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
