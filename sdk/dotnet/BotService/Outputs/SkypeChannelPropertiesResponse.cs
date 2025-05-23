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
    /// The parameters to provide for the Microsoft Teams channel.
    /// </summary>
    [OutputType]
    public sealed class SkypeChannelPropertiesResponse
    {
        /// <summary>
        /// Calling web hook for Skype channel
        /// </summary>
        public readonly string? CallingWebHook;
        /// <summary>
        /// Enable calling for Skype channel
        /// </summary>
        public readonly bool? EnableCalling;
        /// <summary>
        /// Enable groups for Skype channel
        /// </summary>
        public readonly bool? EnableGroups;
        /// <summary>
        /// Enable media cards for Skype channel
        /// </summary>
        public readonly bool? EnableMediaCards;
        /// <summary>
        /// Enable messaging for Skype channel
        /// </summary>
        public readonly bool? EnableMessaging;
        /// <summary>
        /// Enable screen sharing for Skype channel
        /// </summary>
        public readonly bool? EnableScreenSharing;
        /// <summary>
        /// Enable video for Skype channel
        /// </summary>
        public readonly bool? EnableVideo;
        /// <summary>
        /// Group mode for Skype channel
        /// </summary>
        public readonly string? GroupsMode;
        /// <summary>
        /// Incoming call route for Skype channel
        /// </summary>
        public readonly string? IncomingCallRoute;
        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        public readonly bool IsEnabled;

        [OutputConstructor]
        private SkypeChannelPropertiesResponse(
            string? callingWebHook,

            bool? enableCalling,

            bool? enableGroups,

            bool? enableMediaCards,

            bool? enableMessaging,

            bool? enableScreenSharing,

            bool? enableVideo,

            string? groupsMode,

            string? incomingCallRoute,

            bool isEnabled)
        {
            CallingWebHook = callingWebHook;
            EnableCalling = enableCalling;
            EnableGroups = enableGroups;
            EnableMediaCards = enableMediaCards;
            EnableMessaging = enableMessaging;
            EnableScreenSharing = enableScreenSharing;
            EnableVideo = enableVideo;
            GroupsMode = groupsMode;
            IncomingCallRoute = incomingCallRoute;
            IsEnabled = isEnabled;
        }
    }
}
