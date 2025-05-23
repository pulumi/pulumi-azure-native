// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scheduler.Outputs
{

    [OutputType]
    public sealed class JobMaxRecurrenceResponse
    {
        /// <summary>
        /// Gets or sets the frequency of recurrence (second, minute, hour, day, week, month).
        /// </summary>
        public readonly string? Frequency;
        /// <summary>
        /// Gets or sets the interval between retries.
        /// </summary>
        public readonly int? Interval;

        [OutputConstructor]
        private JobMaxRecurrenceResponse(
            string? frequency,

            int? interval)
        {
            Frequency = frequency;
            Interval = interval;
        }
    }
}
