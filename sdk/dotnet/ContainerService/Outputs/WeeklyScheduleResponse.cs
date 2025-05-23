// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Outputs
{

    /// <summary>
    /// For schedules like: 'recur every Monday' or 'recur every 3 weeks on Wednesday'.
    /// </summary>
    [OutputType]
    public sealed class WeeklyScheduleResponse
    {
        /// <summary>
        /// Specifies on which day of the week the maintenance occurs.
        /// </summary>
        public readonly string DayOfWeek;
        /// <summary>
        /// Specifies the number of weeks between each set of occurrences.
        /// </summary>
        public readonly int IntervalWeeks;

        [OutputConstructor]
        private WeeklyScheduleResponse(
            string dayOfWeek,

            int intervalWeeks)
        {
            DayOfWeek = dayOfWeek;
            IntervalWeeks = intervalWeeks;
        }
    }
}
