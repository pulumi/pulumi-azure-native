// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer.Inputs
{

    public sealed class SubscriberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email of the subscriber
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Number specifying what notifications to receive
        /// </summary>
        [Input("notifications")]
        public Input<double>? Notifications { get; set; }

        public SubscriberArgs()
        {
        }
        public static new SubscriberArgs Empty => new SubscriberArgs();
    }
}
