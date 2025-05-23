// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WebPubSub.Inputs
{

    /// <summary>
    /// Properties of event handler.
    /// </summary>
    public sealed class EventHandlerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Upstream auth settings. If not set, no auth is used for upstream messages.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.UpstreamAuthSettingsArgs>? Auth { get; set; }

        [Input("systemEvents")]
        private InputList<string>? _systemEvents;

        /// <summary>
        /// Gets or sets the list of system events.
        /// </summary>
        public InputList<string> SystemEvents
        {
            get => _systemEvents ?? (_systemEvents = new InputList<string>());
            set => _systemEvents = value;
        }

        /// <summary>
        /// Gets or sets the EventHandler URL template. You can use a predefined parameter {hub} and {event} inside the template, the value of the EventHandler URL is dynamically calculated when the client request comes in.
        /// For example, UrlTemplate can be `http://example.com/api/{hub}/{event}`. The host part can't contains parameters.
        /// </summary>
        [Input("urlTemplate", required: true)]
        public Input<string> UrlTemplate { get; set; } = null!;

        /// <summary>
        /// Gets or sets the matching pattern for event names.
        /// There are 3 kinds of patterns supported:
        ///     1. "*", it matches any event name
        ///     2. Combine multiple events with ",", for example "event1,event2", it matches event "event1" and "event2"
        ///     3. A single event name, for example, "event1", it matches "event1"
        /// </summary>
        [Input("userEventPattern")]
        public Input<string>? UserEventPattern { get; set; }

        public EventHandlerArgs()
        {
        }
        public static new EventHandlerArgs Empty => new EventHandlerArgs();
    }
}
