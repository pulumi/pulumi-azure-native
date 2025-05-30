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
    /// Monthly recurrence object.
    /// </summary>
    [OutputType]
    public sealed class MonthlyRecurrenceResponse
    {
        /// <summary>
        /// Specifies the values for monthly recurrence pattern.
        /// </summary>
        public readonly ImmutableArray<int> DaysOfMonth;
        /// <summary>
        /// End time for recurrence.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Specifies when the recurrence should be applied.
        /// Expected value is 'Monthly'.
        /// </summary>
        public readonly string RecurrenceType;
        /// <summary>
        /// Start time for recurrence.
        /// </summary>
        public readonly string? StartTime;

        [OutputConstructor]
        private MonthlyRecurrenceResponse(
            ImmutableArray<int> daysOfMonth,

            string? endTime,

            string recurrenceType,

            string? startTime)
        {
            DaysOfMonth = daysOfMonth;
            EndTime = endTime;
            RecurrenceType = recurrenceType;
            StartTime = startTime;
        }
    }
}
