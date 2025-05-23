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
    /// The event type information for Channels.
    /// </summary>
    [OutputType]
    public sealed class EventTypeInfoResponse
    {
        /// <summary>
        /// A collection of inline event types for the resource. The inline event type keys are of type string which represents the name of the event.
        /// An example of a valid inline event name is "Contoso.OrderCreated".
        /// The inline event type values are of type InlineEventProperties and will contain additional information for every inline event type.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.InlineEventPropertiesResponse>? InlineEventTypes;
        /// <summary>
        /// The kind of event type used.
        /// </summary>
        public readonly string? Kind;

        [OutputConstructor]
        private EventTypeInfoResponse(
            ImmutableDictionary<string, Outputs.InlineEventPropertiesResponse>? inlineEventTypes,

            string? kind)
        {
            InlineEventTypes = inlineEventTypes;
            Kind = kind;
        }
    }
}
