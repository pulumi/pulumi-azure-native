// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.Inputs
{

    /// <summary>
    /// For schedules like: 'recur every month on the first Monday' or 'recur every 3 months on last Friday'.
    /// </summary>
    public sealed class RelativeMonthlyScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies on which day of the week the maintenance occurs.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.WeekDay> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// Specifies the number of months between each set of occurrences.
        /// </summary>
        [Input("intervalMonths", required: true)]
        public Input<int> IntervalMonths { get; set; } = null!;

        /// <summary>
        /// Specifies on which week of the month the dayOfWeek applies.
        /// </summary>
        [Input("weekIndex", required: true)]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.Type> WeekIndex { get; set; } = null!;

        public RelativeMonthlyScheduleArgs()
        {
        }
        public static new RelativeMonthlyScheduleArgs Empty => new RelativeMonthlyScheduleArgs();
    }
}
