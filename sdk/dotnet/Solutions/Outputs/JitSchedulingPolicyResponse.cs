// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions.Outputs
{

    /// <summary>
    /// The JIT scheduling policies.
    /// </summary>
    [OutputType]
    public sealed class JitSchedulingPolicyResponse
    {
        public readonly string Duration;
        /// <summary>
        /// The start time of the request.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The type of JIT schedule.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private JitSchedulingPolicyResponse(
            string duration,

            string startTime,

            string type)
        {
            Duration = duration;
            StartTime = startTime;
            Type = type;
        }
    }
}
