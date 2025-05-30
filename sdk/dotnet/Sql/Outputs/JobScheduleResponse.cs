// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// Scheduling properties of a job.
    /// </summary>
    [OutputType]
    public sealed class JobScheduleResponse
    {
        /// <summary>
        /// Whether or not the schedule is enabled.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Schedule end time.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// Value of the schedule's recurring interval, if the ScheduleType is recurring. ISO8601 duration format.
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// Schedule start time.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// Schedule interval type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private JobScheduleResponse(
            bool? enabled,

            string? endTime,

            string? interval,

            string? startTime,

            string? type)
        {
            Enabled = enabled;
            EndTime = endTime;
            Interval = interval;
            StartTime = startTime;
            Type = type;
        }
    }
}
