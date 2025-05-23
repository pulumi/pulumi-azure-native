// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of EventBridgeConfiguration
    /// </summary>
    public sealed class EventBridgeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables delivery of events to Amazon EventBridge.
        /// </summary>
        [Input("eventBridgeEnabled")]
        public Input<bool>? EventBridgeEnabled { get; set; }

        public EventBridgeConfigurationArgs()
        {
            EventBridgeEnabled = true;
        }
        public static new EventBridgeConfigurationArgs Empty => new EventBridgeConfigurationArgs();
    }
}
