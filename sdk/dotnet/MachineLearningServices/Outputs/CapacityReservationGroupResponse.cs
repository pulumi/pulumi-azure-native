// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class CapacityReservationGroupResponse
    {
        /// <summary>
        /// Offer used by this capacity reservation group.
        /// </summary>
        public readonly Outputs.ServerlessOfferResponse? Offer;
        /// <summary>
        /// [Required] Specifies the amount of capacity to reserve.
        /// </summary>
        public readonly int ReservedCapacity;

        [OutputConstructor]
        private CapacityReservationGroupResponse(
            Outputs.ServerlessOfferResponse? offer,

            int reservedCapacity)
        {
            Offer = offer;
            ReservedCapacity = reservedCapacity;
        }
    }
}
