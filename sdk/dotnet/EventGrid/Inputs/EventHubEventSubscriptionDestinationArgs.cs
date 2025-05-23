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
    /// Information about the event hub destination for an event subscription.
    /// </summary>
    public sealed class EventHubEventSubscriptionDestinationArgs : global::Pulumi.ResourceArgs
    {
        [Input("deliveryAttributeMappings")]
        private InputList<Union<Inputs.DynamicDeliveryAttributeMappingArgs, Inputs.StaticDeliveryAttributeMappingArgs>>? _deliveryAttributeMappings;

        /// <summary>
        /// Delivery attribute details.
        /// </summary>
        public InputList<Union<Inputs.DynamicDeliveryAttributeMappingArgs, Inputs.StaticDeliveryAttributeMappingArgs>> DeliveryAttributeMappings
        {
            get => _deliveryAttributeMappings ?? (_deliveryAttributeMappings = new InputList<Union<Inputs.DynamicDeliveryAttributeMappingArgs, Inputs.StaticDeliveryAttributeMappingArgs>>());
            set => _deliveryAttributeMappings = value;
        }

        /// <summary>
        /// Type of the endpoint for the event subscription destination.
        /// Expected value is 'EventHub'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public EventHubEventSubscriptionDestinationArgs()
        {
        }
        public static new EventHubEventSubscriptionDestinationArgs Empty => new EventHubEventSubscriptionDestinationArgs();
    }
}
