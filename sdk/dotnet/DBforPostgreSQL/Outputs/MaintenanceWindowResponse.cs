// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Outputs
{

    /// <summary>
    /// Maintenance window properties of a flexible server.
    /// </summary>
    [OutputType]
    public sealed class MaintenanceWindowResponse
    {
        /// <summary>
        /// Indicates whether custom window is enabled or disabled.
        /// </summary>
        public readonly string? CustomWindow;
        /// <summary>
        /// Day of the week to be used for maintenance window.
        /// </summary>
        public readonly int? DayOfWeek;
        /// <summary>
        /// Start hour to be used for maintenance window.
        /// </summary>
        public readonly int? StartHour;
        /// <summary>
        /// Start minute to be used for maintenance window.
        /// </summary>
        public readonly int? StartMinute;

        [OutputConstructor]
        private MaintenanceWindowResponse(
            string? customWindow,

            int? dayOfWeek,

            int? startHour,

            int? startMinute)
        {
            CustomWindow = customWindow;
            DayOfWeek = dayOfWeek;
            StartHour = startHour;
            StartMinute = startMinute;
        }
    }
}
