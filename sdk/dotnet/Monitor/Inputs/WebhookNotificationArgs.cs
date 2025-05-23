// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Inputs
{

    /// <summary>
    /// Webhook notification of an autoscale event.
    /// </summary>
    public sealed class WebhookNotificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// a property bag of settings. This value can be empty.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// the service address to receive the notification.
        /// </summary>
        [Input("serviceUri")]
        public Input<string>? ServiceUri { get; set; }

        public WebhookNotificationArgs()
        {
        }
        public static new WebhookNotificationArgs Empty => new WebhookNotificationArgs();
    }
}
