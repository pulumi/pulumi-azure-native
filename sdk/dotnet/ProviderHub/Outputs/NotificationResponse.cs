// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Outputs
{

    [OutputType]
    public sealed class NotificationResponse
    {
        /// <summary>
        /// The notification type.
        /// </summary>
        public readonly string? NotificationType;
        /// <summary>
        /// Whether notifications should be skipped.
        /// </summary>
        public readonly string? SkipNotifications;

        [OutputConstructor]
        private NotificationResponse(
            string? notificationType,

            string? skipNotifications)
        {
            NotificationType = notificationType;
            SkipNotifications = skipNotifications;
        }
    }
}
