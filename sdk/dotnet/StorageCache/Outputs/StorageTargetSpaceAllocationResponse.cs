// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorageCache.Outputs
{

    /// <summary>
    /// Storage Target space allocation properties.
    /// </summary>
    [OutputType]
    public sealed class StorageTargetSpaceAllocationResponse
    {
        /// <summary>
        /// The percentage of cache space allocated for this storage target
        /// </summary>
        public readonly int? AllocationPercentage;
        /// <summary>
        /// Name of the storage target.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private StorageTargetSpaceAllocationResponse(
            int? allocationPercentage,

            string? name)
        {
            AllocationPercentage = allocationPercentage;
            Name = name;
        }
    }
}
