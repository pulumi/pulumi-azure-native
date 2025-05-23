// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Inputs
{

    /// <summary>
    /// Information about the retry policy for an event subscription.
    /// </summary>
    public sealed class RetryPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time To Live (in minutes) for events.
        /// </summary>
        [Input("eventTimeToLiveInMinutes")]
        public Input<int>? EventTimeToLiveInMinutes { get; set; }

        /// <summary>
        /// Maximum number of delivery retry attempts for events.
        /// </summary>
        [Input("maxDeliveryAttempts")]
        public Input<int>? MaxDeliveryAttempts { get; set; }

        public RetryPolicyArgs()
        {
            EventTimeToLiveInMinutes = 1440;
            MaxDeliveryAttempts = 30;
        }
        public static new RetryPolicyArgs Empty => new RetryPolicyArgs();
    }
}
