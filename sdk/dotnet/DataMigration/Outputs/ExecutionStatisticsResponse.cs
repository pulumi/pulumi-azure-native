// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    /// <summary>
    /// Description about the errors happen while performing migration validation
    /// </summary>
    [OutputType]
    public sealed class ExecutionStatisticsResponse
    {
        /// <summary>
        /// CPU Time in millisecond(s) for the query execution
        /// </summary>
        public readonly double? CpuTimeMs;
        /// <summary>
        /// Time taken in millisecond(s) for executing the query
        /// </summary>
        public readonly double? ElapsedTimeMs;
        /// <summary>
        /// No. of query executions
        /// </summary>
        public readonly double? ExecutionCount;
        /// <summary>
        /// Indicates whether the query resulted in an error
        /// </summary>
        public readonly bool? HasErrors;
        /// <summary>
        /// List of sql Errors
        /// </summary>
        public readonly ImmutableArray<string> SqlErrors;
        /// <summary>
        /// Dictionary of sql query execution wait types and the respective statistics
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.WaitStatisticsResponse>? WaitStats;

        [OutputConstructor]
        private ExecutionStatisticsResponse(
            double? cpuTimeMs,

            double? elapsedTimeMs,

            double? executionCount,

            bool? hasErrors,

            ImmutableArray<string> sqlErrors,

            ImmutableDictionary<string, Outputs.WaitStatisticsResponse>? waitStats)
        {
            CpuTimeMs = cpuTimeMs;
            ElapsedTimeMs = elapsedTimeMs;
            ExecutionCount = executionCount;
            HasErrors = hasErrors;
            SqlErrors = sqlErrors;
            WaitStats = waitStats;
        }
    }
}
