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
    /// The parameters to provide for the Email channel.
    /// </summary>
    [OutputType]
    public sealed class EmailChannelPropertiesResponse
    {
        /// <summary>
        /// Email channel auth method. 0 Password (Default); 1 Graph.
        /// </summary>
        public readonly double? AuthMethod;
        /// <summary>
        /// The email address
        /// </summary>
        public readonly string EmailAddress;
        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The magic code for setting up the modern authentication.
        /// </summary>
        public readonly string? MagicCode;
        /// <summary>
        /// The password for the email address. Value only returned through POST to the action Channel List API, otherwise empty.
        /// </summary>
        public readonly string? Password;

        [OutputConstructor]
        private EmailChannelPropertiesResponse(
            double? authMethod,

            string emailAddress,

            bool isEnabled,

            string? magicCode,

            string? password)
        {
            AuthMethod = authMethod;
            EmailAddress = emailAddress;
            IsEnabled = isEnabled;
            MagicCode = magicCode;
            Password = password;
        }
    }
}
