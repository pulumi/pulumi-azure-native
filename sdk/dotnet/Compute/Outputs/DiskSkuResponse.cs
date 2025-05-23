// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    /// <summary>
    /// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
    /// </summary>
    [OutputType]
    public sealed class DiskSkuResponse
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The sku tier.
        /// </summary>
        public readonly string Tier;

        [OutputConstructor]
        private DiskSkuResponse(
            string? name,

            string tier)
        {
            Name = name;
            Tier = tier;
        }
    }
}
