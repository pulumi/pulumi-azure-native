// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Manual trigger configuration for a single execution job. Properties replicaCompletionCount and parallelism would be set to 1 by default
    /// </summary>
    [OutputType]
    public sealed class JobConfigurationResponseManualTriggerConfig
    {
        /// <summary>
        /// Number of parallel replicas of a job that can run at a given time.
        /// </summary>
        public readonly int? Parallelism;
        /// <summary>
        /// Minimum number of successful replica completions before overall job completion.
        /// </summary>
        public readonly int? ReplicaCompletionCount;

        [OutputConstructor]
        private JobConfigurationResponseManualTriggerConfig(
            int? parallelism,

            int? replicaCompletionCount)
        {
            Parallelism = parallelism;
            ReplicaCompletionCount = replicaCompletionCount;
        }
    }
}
