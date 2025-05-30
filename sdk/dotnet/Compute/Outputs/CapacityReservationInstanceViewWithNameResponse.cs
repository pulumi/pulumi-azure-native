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
    /// The instance view of a capacity reservation that includes the name of the capacity reservation. It is used for the response to the instance view of a capacity reservation group.
    /// </summary>
    [OutputType]
    public sealed class CapacityReservationInstanceViewWithNameResponse
    {
        /// <summary>
        /// The name of the capacity reservation.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponse> Statuses;
        /// <summary>
        /// Unutilized capacity of the capacity reservation.
        /// </summary>
        public readonly Outputs.CapacityReservationUtilizationResponse? UtilizationInfo;

        [OutputConstructor]
        private CapacityReservationInstanceViewWithNameResponse(
            string name,

            ImmutableArray<Outputs.InstanceViewStatusResponse> statuses,

            Outputs.CapacityReservationUtilizationResponse? utilizationInfo)
        {
            Name = name;
            Statuses = statuses;
            UtilizationInfo = utilizationInfo;
        }
    }
}
