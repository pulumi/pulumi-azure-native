// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.Outputs
{

    /// <summary>
    /// The workflow trigger recurrence.
    /// </summary>
    [OutputType]
    public sealed class WorkflowTriggerRecurrenceResponse
    {
        /// <summary>
        /// The end time.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// The frequency.
        /// </summary>
        public readonly string? Frequency;
        /// <summary>
        /// The interval.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// The recurrence schedule.
        /// </summary>
        public readonly Outputs.RecurrenceScheduleResponse? Schedule;
        /// <summary>
        /// The start time.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// The time zone.
        /// </summary>
        public readonly string? TimeZone;

        [OutputConstructor]
        private WorkflowTriggerRecurrenceResponse(
            string? endTime,

            string? frequency,

            int? interval,

            Outputs.RecurrenceScheduleResponse? schedule,

            string? startTime,

            string? timeZone)
        {
            EndTime = endTime;
            Frequency = frequency;
            Interval = interval;
            Schedule = schedule;
            StartTime = startTime;
            TimeZone = timeZone;
        }
    }
}
