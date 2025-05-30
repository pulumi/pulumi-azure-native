// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// Sku of Azure Spring Apps
    /// </summary>
    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// Current capacity of the target resource
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// Name of the Sku
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tier of the Sku
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private SkuResponse(
            int? capacity,

            string? name,

            string? tier)
        {
            Capacity = capacity;
            Name = name;
            Tier = tier;
        }
    }
}
