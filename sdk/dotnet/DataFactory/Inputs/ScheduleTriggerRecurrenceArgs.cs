// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// The workflow trigger recurrence.
    /// </summary>
    public sealed class ScheduleTriggerRecurrenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The frequency.
        /// </summary>
        [Input("frequency")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.RecurrenceFrequency>? Frequency { get; set; }

        /// <summary>
        /// The interval.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The recurrence schedule.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.RecurrenceScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// The start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The time zone.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public ScheduleTriggerRecurrenceArgs()
        {
        }
        public static new ScheduleTriggerRecurrenceArgs Empty => new ScheduleTriggerRecurrenceArgs();
    }
}
