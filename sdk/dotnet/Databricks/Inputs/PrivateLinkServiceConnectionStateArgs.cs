// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Inputs
{

    /// <summary>
    /// The current state of a private endpoint connection
    /// </summary>
    public sealed class PrivateLinkServiceConnectionStateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actions required for a private endpoint connection
        /// </summary>
        [Input("actionsRequired")]
        public Input<string>? ActionsRequired { get; set; }

        /// <summary>
        /// The description for the current state of a private endpoint connection
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The status of a private endpoint connection
        /// </summary>
        [Input("status", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Databricks.PrivateLinkServiceConnectionStatus> Status { get; set; } = null!;

        public PrivateLinkServiceConnectionStateArgs()
        {
        }
        public static new PrivateLinkServiceConnectionStateArgs Empty => new PrivateLinkServiceConnectionStateArgs();
    }
}
