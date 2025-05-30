// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    [OutputType]
    public sealed class MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseErrorResponse
    {
        /// <summary>
        /// Error message
        /// </summary>
        public readonly string? ErrorMessage;
        /// <summary>
        /// List of error events.
        /// </summary>
        public readonly ImmutableArray<Outputs.SyncMigrationDatabaseErrorEventResponse> Events;
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Result type
        /// Expected value is 'DatabaseLevelErrorOutput'.
        /// </summary>
        public readonly string ResultType;

        [OutputConstructor]
        private MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseErrorResponse(
            string? errorMessage,

            ImmutableArray<Outputs.SyncMigrationDatabaseErrorEventResponse> events,

            string id,

            string resultType)
        {
            ErrorMessage = errorMessage;
            Events = events;
            Id = id;
            ResultType = resultType;
        }
    }
}
