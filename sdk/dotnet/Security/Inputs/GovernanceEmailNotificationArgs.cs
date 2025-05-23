// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Inputs
{

    /// <summary>
    /// The governance email weekly notification configuration.
    /// </summary>
    public sealed class GovernanceEmailNotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exclude manager from weekly email notification.
        /// </summary>
        [Input("disableManagerEmailNotification")]
        public Input<bool>? DisableManagerEmailNotification { get; set; }

        /// <summary>
        /// Exclude  owner from weekly email notification.
        /// </summary>
        [Input("disableOwnerEmailNotification")]
        public Input<bool>? DisableOwnerEmailNotification { get; set; }

        public GovernanceEmailNotificationArgs()
        {
        }
        public static new GovernanceEmailNotificationArgs Empty => new GovernanceEmailNotificationArgs();
    }
}
