// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// The execution options of a job step.
    /// </summary>
    [OutputType]
    public sealed class JobStepExecutionOptionsResponse
    {
        /// <summary>
        /// Initial delay between retries for job step execution.
        /// </summary>
        public readonly int? InitialRetryIntervalSeconds;
        /// <summary>
        /// The maximum amount of time to wait between retries for job step execution.
        /// </summary>
        public readonly int? MaximumRetryIntervalSeconds;
        /// <summary>
        /// Maximum number of times the job step will be reattempted if the first attempt fails.
        /// </summary>
        public readonly int? RetryAttempts;
        /// <summary>
        /// The backoff multiplier for the time between retries.
        /// </summary>
        public readonly double? RetryIntervalBackoffMultiplier;
        /// <summary>
        /// Execution timeout for the job step.
        /// </summary>
        public readonly int? TimeoutSeconds;

        [OutputConstructor]
        private JobStepExecutionOptionsResponse(
            int? initialRetryIntervalSeconds,

            int? maximumRetryIntervalSeconds,

            int? retryAttempts,

            double? retryIntervalBackoffMultiplier,

            int? timeoutSeconds)
        {
            InitialRetryIntervalSeconds = initialRetryIntervalSeconds;
            MaximumRetryIntervalSeconds = maximumRetryIntervalSeconds;
            RetryAttempts = retryAttempts;
            RetryIntervalBackoffMultiplier = retryIntervalBackoffMultiplier;
            TimeoutSeconds = timeoutSeconds;
        }
    }
}
