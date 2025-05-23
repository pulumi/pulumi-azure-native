// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SqlVirtualMachine.Inputs
{

    /// <summary>
    /// Set assessment schedule for SQL Server.
    /// </summary>
    public sealed class ScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of the week to run assessment.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<Pulumi.AzureNative.SqlVirtualMachine.AssessmentDayOfWeek>? DayOfWeek { get; set; }

        /// <summary>
        /// Enable or disable assessment schedule on SQL virtual machine.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Occurrence of the DayOfWeek day within a month to schedule assessment. Takes values: 1,2,3,4 and -1. Use -1 for last DayOfWeek day of the month
        /// </summary>
        [Input("monthlyOccurrence")]
        public Input<int>? MonthlyOccurrence { get; set; }

        /// <summary>
        /// Time of the day in HH:mm format. Eg. 17:30
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Number of weeks to schedule between 2 assessment runs. Takes value from 1-6
        /// </summary>
        [Input("weeklyInterval")]
        public Input<int>? WeeklyInterval { get; set; }

        public ScheduleArgs()
        {
        }
        public static new ScheduleArgs Empty => new ScheduleArgs();
    }
}
