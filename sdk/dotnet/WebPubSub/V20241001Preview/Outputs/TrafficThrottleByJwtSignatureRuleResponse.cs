// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WebPubSub.V20241001Preview.Outputs
{

    /// <summary>
    /// Throttle the client traffic by the JWT signature
    /// </summary>
    [OutputType]
    public sealed class TrafficThrottleByJwtSignatureRuleResponse
    {
        /// <summary>
        /// The aggregation window for the message bytes. The message bytes will be aggregated in this window and be reset after the window. Default value is 60 seconds.
        /// </summary>
        public readonly int? AggregationWindowInSeconds;
        /// <summary>
        /// Maximum accumulated inbound message bytes allowed for the same JWT signature within a time window. Clients with the same JWT signature will get disconnected if the message bytes exceeds this value. Default value is 1GB.
        /// </summary>
        public readonly double? MaxInboundMessageBytes;
        /// <summary>
        /// 
        /// Expected value is 'TrafficThrottleByJwtSignatureRule'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TrafficThrottleByJwtSignatureRuleResponse(
            int? aggregationWindowInSeconds,

            double? maxInboundMessageBytes,

            string type)
        {
            AggregationWindowInSeconds = aggregationWindowInSeconds;
            MaxInboundMessageBytes = maxInboundMessageBytes;
            Type = type;
        }
    }
}
