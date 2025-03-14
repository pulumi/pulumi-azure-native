// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WebPubSub.V20210901Preview.Inputs
{

    /// <summary>
    /// live trace category configuration of a Microsoft.SignalRService resource.
    /// </summary>
    public sealed class LiveTraceCategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or the log category is enabled.
        /// Available values: true, false.
        /// Case insensitive.
        /// </summary>
        [Input("enabled")]
        public Input<string>? Enabled { get; set; }

        /// <summary>
        /// Gets or sets the log category's name.
        /// Available values: ConnectivityLogs, MessagingLogs.
        /// Case insensitive.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LiveTraceCategoryArgs()
        {
        }
        public static new LiveTraceCategoryArgs Empty => new LiveTraceCategoryArgs();
    }
}
