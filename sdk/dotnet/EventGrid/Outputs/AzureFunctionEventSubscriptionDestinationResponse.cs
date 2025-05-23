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
    /// Information about the azure function destination for an event subscription.
    /// </summary>
    [OutputType]
    public sealed class AzureFunctionEventSubscriptionDestinationResponse
    {
        /// <summary>
        /// Delivery attribute details.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.DynamicDeliveryAttributeMappingResponse, Outputs.StaticDeliveryAttributeMappingResponse>> DeliveryAttributeMappings;
        /// <summary>
        /// Type of the endpoint for the event subscription destination.
        /// Expected value is 'AzureFunction'.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// Maximum number of events per batch.
        /// </summary>
        public readonly int? MaxEventsPerBatch;
        /// <summary>
        /// Preferred batch size in Kilobytes.
        /// </summary>
        public readonly int? PreferredBatchSizeInKilobytes;
        /// <summary>
        /// The Azure Resource Id that represents the endpoint of the Azure Function destination of an event subscription.
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private AzureFunctionEventSubscriptionDestinationResponse(
            ImmutableArray<Union<Outputs.DynamicDeliveryAttributeMappingResponse, Outputs.StaticDeliveryAttributeMappingResponse>> deliveryAttributeMappings,

            string endpointType,

            int? maxEventsPerBatch,

            int? preferredBatchSizeInKilobytes,

            string? resourceId)
        {
            DeliveryAttributeMappings = deliveryAttributeMappings;
            EndpointType = endpointType;
            MaxEventsPerBatch = maxEventsPerBatch;
            PreferredBatchSizeInKilobytes = preferredBatchSizeInKilobytes;
            ResourceId = resourceId;
        }
    }
}
