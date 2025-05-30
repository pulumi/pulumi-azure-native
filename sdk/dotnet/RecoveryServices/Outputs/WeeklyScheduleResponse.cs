// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    [OutputType]
    public sealed class WeeklyScheduleResponse
    {
        public readonly ImmutableArray<string> ScheduleRunDays;
        /// <summary>
        /// List of times of day this schedule has to be run.
        /// </summary>
        public readonly ImmutableArray<string> ScheduleRunTimes;

        [OutputConstructor]
        private WeeklyScheduleResponse(
            ImmutableArray<string> scheduleRunDays,

            ImmutableArray<string> scheduleRunTimes)
        {
            ScheduleRunDays = scheduleRunDays;
            ScheduleRunTimes = scheduleRunTimes;
        }
    }
}
