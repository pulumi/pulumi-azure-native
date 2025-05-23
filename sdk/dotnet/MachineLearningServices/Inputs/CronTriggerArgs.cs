// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class CronTriggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies end time of schedule in ISO 8601, but without a UTC offset. Refer https://en.wikipedia.org/wiki/ISO_8601.
        /// Recommented format would be "2022-06-01T00:00:01"
        /// If not present, the schedule will run indefinitely
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// [Required] Specifies cron expression of schedule.
        /// The expression should follow NCronTab format.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// Specifies start time of schedule in ISO 8601 format, but without a UTC offset.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Specifies time zone in which the schedule runs.
        /// TimeZone should follow Windows time zone format. Refer: https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/default-time-zones?view=windows-11
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// 
        /// Expected value is 'Cron'.
        /// </summary>
        [Input("triggerType", required: true)]
        public Input<string> TriggerType { get; set; } = null!;

        public CronTriggerArgs()
        {
            TimeZone = "UTC";
        }
        public static new CronTriggerArgs Empty => new CronTriggerArgs();
    }
}
