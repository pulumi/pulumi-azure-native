// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.Inputs
{

    /// <summary>
    /// Live trace category configuration of a Microsoft.SignalRService resource.
    /// </summary>
    public sealed class LiveTraceCategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or the live trace category is enabled.
        /// Available values: true, false.
        /// Case insensitive.
        /// </summary>
        [Input("enabled")]
        public Input<string>? Enabled { get; set; }

        /// <summary>
        /// Gets or sets the live trace category's name.
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
