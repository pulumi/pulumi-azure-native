// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class CapacityReservationGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Offer used by this capacity reservation group.
        /// </summary>
        [Input("offer")]
        public Input<Inputs.ServerlessOfferArgs>? Offer { get; set; }

        /// <summary>
        /// [Required] Specifies the amount of capacity to reserve.
        /// </summary>
        [Input("reservedCapacity", required: true)]
        public Input<int> ReservedCapacity { get; set; } = null!;

        public CapacityReservationGroupArgs()
        {
        }
        public static new CapacityReservationGroupArgs Empty => new CapacityReservationGroupArgs();
    }
}
