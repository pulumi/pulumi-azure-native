// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BotService.V20230915Preview.Inputs
{

    /// <summary>
    /// The parameters to provide for the Alexa channel.
    /// </summary>
    public sealed class AlexaChannelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Alexa skill Id
        /// </summary>
        [Input("alexaSkillId", required: true)]
        public Input<string> AlexaSkillId { get; set; } = null!;

        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        public AlexaChannelPropertiesArgs()
        {
        }
        public static new AlexaChannelPropertiesArgs Empty => new AlexaChannelPropertiesArgs();
    }
}
