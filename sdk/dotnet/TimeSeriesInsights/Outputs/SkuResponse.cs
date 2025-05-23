// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TimeSeriesInsights.Outputs
{

    /// <summary>
    /// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
    /// </summary>
    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// The capacity of the sku. For Gen1 environments, this value can be changed to support scale out of environments after they have been created.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// The name of this SKU.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private SkuResponse(
            int capacity,

            string name)
        {
            Capacity = capacity;
            Name = name;
        }
    }
}
