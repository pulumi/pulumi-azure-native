// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// An event to be notified for.
    /// </summary>
    public sealed class EventArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The event type for which this notification is enabled (i.e. AutoShutdown, Cost)
        /// </summary>
        [Input("eventName")]
        public InputUnion<string, Pulumi.AzureNative.DevTestLab.NotificationChannelEventType>? EventName { get; set; }

        public EventArgs()
        {
        }
        public static new EventArgs Empty => new EventArgs();
    }
}
