// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.PowerBIDedicated.Outputs
{

    /// <summary>
    /// Represents the SKU name and Azure pricing tier for PowerBI Dedicated capacity resource.
    /// </summary>
    [OutputType]
    public sealed class CapacitySkuResponse
    {
        /// <summary>
        /// The capacity of the SKU.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// Name of the SKU level.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the Azure pricing tier to which the SKU applies.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private CapacitySkuResponse(
            int? capacity,

            string name,

            string? tier)
        {
            Capacity = capacity;
            Name = name;
            Tier = tier;
        }
    }
}
