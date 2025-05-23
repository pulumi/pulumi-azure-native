// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BotService.Inputs
{

    /// <summary>
    /// Email channel definition
    /// </summary>
    public sealed class EmailChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The channel name
        /// Expected value is 'EmailChannel'.
        /// </summary>
        [Input("channelName", required: true)]
        public Input<string> ChannelName { get; set; } = null!;

        /// <summary>
        /// Entity Tag of the resource
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The set of properties specific to email channel resource
        /// </summary>
        [Input("properties")]
        public Input<Inputs.EmailChannelPropertiesArgs>? Properties { get; set; }

        public EmailChannelArgs()
        {
            Location = "global";
        }
        public static new EmailChannelArgs Empty => new EmailChannelArgs();
    }
}
