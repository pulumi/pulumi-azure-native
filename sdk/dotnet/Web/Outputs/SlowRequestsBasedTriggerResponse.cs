// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Outputs
{

    /// <summary>
    /// Trigger based on request execution time.
    /// </summary>
    [OutputType]
    public sealed class SlowRequestsBasedTriggerResponse
    {
        /// <summary>
        /// Request Count.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// Request Path.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Time interval.
        /// </summary>
        public readonly string? TimeInterval;
        /// <summary>
        /// Time taken.
        /// </summary>
        public readonly string? TimeTaken;

        [OutputConstructor]
        private SlowRequestsBasedTriggerResponse(
            int? count,

            string? path,

            string? timeInterval,

            string? timeTaken)
        {
            Count = count;
            Path = path;
            TimeInterval = timeInterval;
            TimeTaken = timeTaken;
        }
    }
}
