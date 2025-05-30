// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// Properties for creating a schedule.
    /// </summary>
    public sealed class ScheduleCreationParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the schedule will occur once each day of the week, specify the daily recurrence.
        /// </summary>
        [Input("dailyRecurrence")]
        public Input<Inputs.DayDetailsArgs>? DailyRecurrence { get; set; }

        /// <summary>
        /// If the schedule will occur multiple times a day, specify the hourly recurrence.
        /// </summary>
        [Input("hourlyRecurrence")]
        public Input<Inputs.HourDetailsArgs>? HourlyRecurrence { get; set; }

        /// <summary>
        /// The name of the virtual machine or environment
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Notification settings.
        /// </summary>
        [Input("notificationSettings")]
        public Input<Inputs.NotificationSettingsArgs>? NotificationSettings { get; set; }

        /// <summary>
        /// The status of the schedule (i.e. Enabled, Disabled)
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.DevTestLab.EnableStatus>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The resource ID to which the schedule belongs
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
        /// </summary>
        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// The time zone ID (e.g. China Standard Time, Greenland Standard Time, Pacific Standard time, etc.). The possible values for this property can be found in `IReadOnlyCollection&lt;string&gt; TimeZoneConverter.TZConvert.KnownWindowsTimeZoneIds` (https://github.com/mattjohnsonpint/TimeZoneConverter/blob/main/README.md)
        /// </summary>
        [Input("timeZoneId")]
        public Input<string>? TimeZoneId { get; set; }

        /// <summary>
        /// If the schedule will occur only some days of the week, specify the weekly recurrence.
        /// </summary>
        [Input("weeklyRecurrence")]
        public Input<Inputs.WeekDetailsArgs>? WeeklyRecurrence { get; set; }

        public ScheduleCreationParameterArgs()
        {
            Status = "Disabled";
        }
        public static new ScheduleCreationParameterArgs Empty => new ScheduleCreationParameterArgs();
    }
}
