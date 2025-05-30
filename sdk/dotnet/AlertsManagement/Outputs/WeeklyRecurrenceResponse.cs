// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AlertsManagement.Outputs
{

    /// <summary>
    /// Weekly recurrence object.
    /// </summary>
    [OutputType]
    public sealed class WeeklyRecurrenceResponse
    {
        /// <summary>
        /// Specifies the values for weekly recurrence pattern.
        /// </summary>
        public readonly ImmutableArray<string> DaysOfWeek;
        /// <summary>
        /// End time for recurrence.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Specifies when the recurrence should be applied.
        /// Expected value is 'Weekly'.
        /// </summary>
        public readonly string RecurrenceType;
        /// <summary>
        /// Start time for recurrence.
        /// </summary>
        public readonly string? StartTime;

        [OutputConstructor]
        private WeeklyRecurrenceResponse(
            ImmutableArray<string> daysOfWeek,

            string? endTime,

            string recurrenceType,

            string? startTime)
        {
            DaysOfWeek = daysOfWeek;
            EndTime = endTime;
            RecurrenceType = recurrenceType;
            StartTime = startTime;
        }
    }
}
