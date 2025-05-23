// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of CapacityReservationTargetResponse
    /// </summary>
    [OutputType]
    public sealed class CapacityReservationTargetResponseResponse
    {
        /// <summary>
        /// &lt;p&gt;The ID of the targeted Capacity Reservation.&lt;/p&gt;
        /// </summary>
        public readonly string? CapacityReservationId;
        /// <summary>
        /// &lt;p&gt;The ARN of the targeted Capacity Reservation group.&lt;/p&gt;
        /// </summary>
        public readonly string? CapacityReservationResourceGroupArn;

        [OutputConstructor]
        private CapacityReservationTargetResponseResponse(
            string? capacityReservationId,

            string? capacityReservationResourceGroupArn)
        {
            CapacityReservationId = capacityReservationId;
            CapacityReservationResourceGroupArn = capacityReservationResourceGroupArn;
        }
    }
}
