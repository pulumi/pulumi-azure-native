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
    /// The parameters to provide for the Sms channel.
    /// </summary>
    [OutputType]
    public sealed class SmsChannelPropertiesResponse
    {
        /// <summary>
        /// The Sms account SID. Value only returned through POST to the action Channel List API, otherwise empty.
        /// </summary>
        public readonly string AccountSID;
        /// <summary>
        /// The Sms auth token. Value only returned through POST to the action Channel List API, otherwise empty.
        /// </summary>
        public readonly string? AuthToken;
        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// Whether this channel is validated for the bot
        /// </summary>
        public readonly bool? IsValidated;
        /// <summary>
        /// The Sms phone
        /// </summary>
        public readonly string Phone;

        [OutputConstructor]
        private SmsChannelPropertiesResponse(
            string accountSID,

            string? authToken,

            bool isEnabled,

            bool? isValidated,

            string phone)
        {
            AccountSID = accountSID;
            AuthToken = authToken;
            IsEnabled = isEnabled;
            IsValidated = isValidated;
            Phone = phone;
        }
    }
}
