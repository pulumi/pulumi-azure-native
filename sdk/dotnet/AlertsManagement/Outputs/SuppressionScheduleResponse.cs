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
    /// Schedule for a given suppression configuration.
    /// </summary>
    [OutputType]
    public sealed class SuppressionScheduleResponse
    {
        /// <summary>
        /// End date for suppression
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// End date for suppression
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Specifies the values for recurrence pattern
        /// </summary>
        public readonly ImmutableArray<int> RecurrenceValues;
        /// <summary>
        /// Start date for suppression
        /// </summary>
        public readonly string? StartDate;
        /// <summary>
        /// Start time for suppression
        /// </summary>
        public readonly string? StartTime;

        [OutputConstructor]
        private SuppressionScheduleResponse(
            string? endDate,

            string? endTime,

            ImmutableArray<int> recurrenceValues,

            string? startDate,

            string? startTime)
        {
            EndDate = endDate;
            EndTime = endTime;
            RecurrenceValues = recurrenceValues;
            StartDate = startDate;
            StartTime = startTime;
        }
    }
}
