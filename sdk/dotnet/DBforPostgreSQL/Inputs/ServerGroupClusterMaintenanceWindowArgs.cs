// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Inputs
{

    /// <summary>
    /// Schedule settings for regular cluster updates.
    /// </summary>
    public sealed class ServerGroupClusterMaintenanceWindowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether custom maintenance window is enabled or not.
        /// </summary>
        [Input("customWindow")]
        public Input<string>? CustomWindow { get; set; }

        /// <summary>
        /// Preferred day of the week for maintenance window.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<int>? DayOfWeek { get; set; }

        /// <summary>
        /// Start hour within preferred day of the week for maintenance window.
        /// </summary>
        [Input("startHour")]
        public Input<int>? StartHour { get; set; }

        /// <summary>
        /// Start minute within the start hour for maintenance window.
        /// </summary>
        [Input("startMinute")]
        public Input<int>? StartMinute { get; set; }

        public ServerGroupClusterMaintenanceWindowArgs()
        {
        }
        public static new ServerGroupClusterMaintenanceWindowArgs Empty => new ServerGroupClusterMaintenanceWindowArgs();
    }
}
