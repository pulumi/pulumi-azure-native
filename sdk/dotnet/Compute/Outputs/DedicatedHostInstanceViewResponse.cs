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
    /// The instance view of a dedicated host.
    /// </summary>
    [OutputType]
    public sealed class DedicatedHostInstanceViewResponse
    {
        /// <summary>
        /// Specifies the unique id of the dedicated physical machine on which the dedicated host resides.
        /// </summary>
        public readonly string AssetId;
        /// <summary>
        /// Unutilized capacity of the dedicated host.
        /// </summary>
        public readonly Outputs.DedicatedHostAvailableCapacityResponse? AvailableCapacity;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponse> Statuses;

        [OutputConstructor]
        private DedicatedHostInstanceViewResponse(
            string assetId,

            Outputs.DedicatedHostAvailableCapacityResponse? availableCapacity,

            ImmutableArray<Outputs.InstanceViewStatusResponse> statuses)
        {
            AssetId = assetId;
            AvailableCapacity = availableCapacity;
            Statuses = statuses;
        }
    }
}
