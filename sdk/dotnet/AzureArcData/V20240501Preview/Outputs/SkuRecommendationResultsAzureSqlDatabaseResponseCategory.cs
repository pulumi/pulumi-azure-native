// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.V20240501Preview.Outputs
{

    [OutputType]
    public sealed class SkuRecommendationResultsAzureSqlDatabaseResponseCategory
    {
        /// <summary>
        /// The compute tier of the target SKU.
        /// </summary>
        public readonly string? ComputeTier;
        /// <summary>
        /// The hardware type of the target SKU.
        /// </summary>
        public readonly string? HardwareType;
        /// <summary>
        /// The SQL purchasing model of the target SKU.
        /// </summary>
        public readonly string? SqlPurchasingModel;
        /// <summary>
        /// The SQL service tier of the target SKU.
        /// </summary>
        public readonly string? SqlServiceTier;
        /// <summary>
        /// Indicates if zone redundancy is available for the target SKU.
        /// </summary>
        public readonly bool? ZoneRedundancyAvailable;

        [OutputConstructor]
        private SkuRecommendationResultsAzureSqlDatabaseResponseCategory(
            string? computeTier,

            string? hardwareType,

            string? sqlPurchasingModel,

            string? sqlServiceTier,

            bool? zoneRedundancyAvailable)
        {
            ComputeTier = computeTier;
            HardwareType = hardwareType;
            SqlPurchasingModel = sqlPurchasingModel;
            SqlServiceTier = sqlServiceTier;
            ZoneRedundancyAvailable = zoneRedundancyAvailable;
        }
    }
}
