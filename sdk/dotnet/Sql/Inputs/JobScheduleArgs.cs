// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Inputs
{

    /// <summary>
    /// Scheduling properties of a job.
    /// </summary>
    public sealed class JobScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the schedule is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Schedule end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Value of the schedule's recurring interval, if the ScheduleType is recurring. ISO8601 duration format.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// Schedule start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Schedule interval type
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNative.Sql.JobScheduleType>? Type { get; set; }

        public JobScheduleArgs()
        {
            EndTime = "9999-12-31T17:29:59+05:30";
            StartTime = "0001-01-01T05:30:00+05:30";
            Type = Pulumi.AzureNative.Sql.JobScheduleType.Once;
        }
        public static new JobScheduleArgs Empty => new JobScheduleArgs();
    }
}
