// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// Details on the Estimated External Storage for AVS Assessment.
    /// </summary>
    [OutputType]
    public sealed class AvsEstimatedExternalStorageResponse
    {
        /// <summary>
        /// Total monthly cost for type of storage.
        /// </summary>
        public readonly double? MonthlyPrice;
        /// <summary>
        /// Recommended External Storage.
        /// </summary>
        public readonly string? StorageType;
        /// <summary>
        /// Predicted storage utilization.
        /// </summary>
        public readonly double? StorageUtilization;
        /// <summary>
        /// Predicted total Storage used in GB.
        /// </summary>
        public readonly double? TotalStorageInGB;

        [OutputConstructor]
        private AvsEstimatedExternalStorageResponse(
            double? monthlyPrice,

            string? storageType,

            double? storageUtilization,

            double? totalStorageInGB)
        {
            MonthlyPrice = monthlyPrice;
            StorageType = storageType;
            StorageUtilization = storageUtilization;
            TotalStorageInGB = totalStorageInGB;
        }
    }
}
