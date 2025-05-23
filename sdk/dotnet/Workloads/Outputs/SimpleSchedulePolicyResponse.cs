// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads.Outputs
{

    /// <summary>
    /// Simple policy schedule.
    /// </summary>
    [OutputType]
    public sealed class SimpleSchedulePolicyResponse
    {
        /// <summary>
        /// Hourly Schedule of this Policy
        /// </summary>
        public readonly Outputs.HourlyScheduleResponse? HourlySchedule;
        /// <summary>
        /// This property will be used as the discriminator for deciding the specific types in the polymorphic chain of types.
        /// Expected value is 'SimpleSchedulePolicy'.
        /// </summary>
        public readonly string SchedulePolicyType;
        /// <summary>
        /// List of days of week this schedule has to be run.
        /// </summary>
        public readonly ImmutableArray<string> ScheduleRunDays;
        /// <summary>
        /// Frequency of the schedule operation of this policy.
        /// </summary>
        public readonly string? ScheduleRunFrequency;
        /// <summary>
        /// List of times of day this schedule has to be run.
        /// </summary>
        public readonly ImmutableArray<string> ScheduleRunTimes;
        /// <summary>
        /// At every number weeks this schedule has to be run.
        /// </summary>
        public readonly int? ScheduleWeeklyFrequency;

        [OutputConstructor]
        private SimpleSchedulePolicyResponse(
            Outputs.HourlyScheduleResponse? hourlySchedule,

            string schedulePolicyType,

            ImmutableArray<string> scheduleRunDays,

            string? scheduleRunFrequency,

            ImmutableArray<string> scheduleRunTimes,

            int? scheduleWeeklyFrequency)
        {
            HourlySchedule = hourlySchedule;
            SchedulePolicyType = schedulePolicyType;
            ScheduleRunDays = scheduleRunDays;
            ScheduleRunFrequency = scheduleRunFrequency;
            ScheduleRunTimes = scheduleRunTimes;
            ScheduleWeeklyFrequency = scheduleWeeklyFrequency;
        }
    }
}
