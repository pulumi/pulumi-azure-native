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
    /// For schedules like: 'recur every month on the first Monday' or 'recur every 3 months on last Friday'.
    /// </summary>
    [OutputType]
    public sealed class RelativeMonthlyScheduleResponse
    {
        /// <summary>
        /// Specifies on which day of the week the maintenance occurs.
        /// </summary>
        public readonly string DayOfWeek;
        /// <summary>
        /// Specifies the number of months between each set of occurrences.
        /// </summary>
        public readonly int IntervalMonths;
        /// <summary>
        /// Specifies on which week of the month the dayOfWeek applies.
        /// </summary>
        public readonly string WeekIndex;

        [OutputConstructor]
        private RelativeMonthlyScheduleResponse(
            string dayOfWeek,

            int intervalMonths,

            string weekIndex)
        {
            DayOfWeek = dayOfWeek;
            IntervalMonths = intervalMonths;
            WeekIndex = weekIndex;
        }
    }
}
