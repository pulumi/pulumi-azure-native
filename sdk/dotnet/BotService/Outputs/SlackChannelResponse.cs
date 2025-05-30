// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BotService.Outputs
{

    /// <summary>
    /// Slack channel definition
    /// </summary>
    [OutputType]
    public sealed class SlackChannelResponse
    {
        /// <summary>
        /// The channel name
        /// Expected value is 'SlackChannel'.
        /// </summary>
        public readonly string ChannelName;
        /// <summary>
        /// Entity Tag of the resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Specifies the location of the resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The set of properties specific to Slack channel resource
        /// </summary>
        public readonly Outputs.SlackChannelPropertiesResponse? Properties;
        /// <summary>
        /// Provisioning state of the resource
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private SlackChannelResponse(
            string channelName,

            string? etag,

            string? location,

            Outputs.SlackChannelPropertiesResponse? properties,

            string provisioningState)
        {
            ChannelName = channelName;
            Etag = etag;
            Location = location;
            Properties = properties;
            ProvisioningState = provisioningState;
        }
    }
}
