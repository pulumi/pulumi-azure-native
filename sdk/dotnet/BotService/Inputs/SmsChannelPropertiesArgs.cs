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
    /// The parameters to provide for the Sms channel.
    /// </summary>
    public sealed class SmsChannelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Sms account SID. Value only returned through POST to the action Channel List API, otherwise empty.
        /// </summary>
        [Input("accountSID", required: true)]
        public Input<string> AccountSID { get; set; } = null!;

        /// <summary>
        /// The Sms auth token. Value only returned through POST to the action Channel List API, otherwise empty.
        /// </summary>
        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// Whether this channel is validated for the bot
        /// </summary>
        [Input("isValidated")]
        public Input<bool>? IsValidated { get; set; }

        /// <summary>
        /// The Sms phone
        /// </summary>
        [Input("phone", required: true)]
        public Input<string> Phone { get; set; } = null!;

        public SmsChannelPropertiesArgs()
        {
        }
        public static new SmsChannelPropertiesArgs Empty => new SmsChannelPropertiesArgs();
    }
}
