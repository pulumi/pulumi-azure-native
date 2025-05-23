// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Outputs
{

    /// <summary>
    /// Properties of the destination info for event subscription supporting push.
    /// </summary>
    [OutputType]
    public sealed class PushInfoResponse
    {
        /// <summary>
        /// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
        /// Uses the managed identity setup on the parent resource (namely, namespace) to acquire the authentication tokens being used during dead-lettering.
        /// </summary>
        public readonly Outputs.DeadLetterWithResourceIdentityResponse? DeadLetterDestinationWithResourceIdentity;
        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery.
        /// </summary>
        public readonly Outputs.DeliveryWithResourceIdentityResponse? DeliveryWithResourceIdentity;
        /// <summary>
        /// Information about the destination where events have to be delivered for the event subscription.
        /// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery.
        /// </summary>
        public readonly object? Destination;
        /// <summary>
        /// Time span duration in ISO 8601 format that determines how long messages are available to the subscription from the time the message was published.
        /// This duration value is expressed using the following format: \'P(n)Y(n)M(n)DT(n)H(n)M(n)S\', where:
        ///     - (n) is replaced by the value of each time element that follows the (n).
        ///     - P is the duration (or Period) designator and is always placed at the beginning of the duration.
        ///     - Y is the year designator, and it follows the value for the number of years.
        ///     - M is the month designator, and it follows the value for the number of months.
        ///     - W is the week designator, and it follows the value for the number of weeks.
        ///     - D is the day designator, and it follows the value for the number of days.
        ///     - T is the time designator, and it precedes the time components.
        ///     - H is the hour designator, and it follows the value for the number of hours.
        ///     - M is the minute designator, and it follows the value for the number of minutes.
        ///     - S is the second designator, and it follows the value for the number of seconds.
        /// This duration value cannot be set greater than the topic’s EventRetentionInDays. It is is an optional field where its minimum value is 1 minute, and its maximum is determined
        /// by topic’s EventRetentionInDays value. The followings are examples of valid values:
        ///     - \'P0DT23H12M\' or \'PT23H12M\': for duration of 23 hours and 12 minutes.
        ///     - \'P1D\' or \'P1DT0H0M0S\': for duration of 1 day.
        /// </summary>
        public readonly string? EventTimeToLive;
        /// <summary>
        /// The maximum delivery count of the events.
        /// </summary>
        public readonly int? MaxDeliveryCount;

        [OutputConstructor]
        private PushInfoResponse(
            Outputs.DeadLetterWithResourceIdentityResponse? deadLetterDestinationWithResourceIdentity,

            Outputs.DeliveryWithResourceIdentityResponse? deliveryWithResourceIdentity,

            object? destination,

            string? eventTimeToLive,

            int? maxDeliveryCount)
        {
            DeadLetterDestinationWithResourceIdentity = deadLetterDestinationWithResourceIdentity;
            DeliveryWithResourceIdentity = deliveryWithResourceIdentity;
            Destination = destination;
            EventTimeToLive = eventTimeToLive;
            MaxDeliveryCount = maxDeliveryCount;
        }
    }
}
