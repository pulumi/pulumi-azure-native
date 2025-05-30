// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL.Outputs
{

    /// <summary>
    /// Migration status of an individual database
    /// </summary>
    [OutputType]
    public sealed class DbMigrationStatusResponse
    {
        /// <summary>
        /// CDC applied changes counter
        /// </summary>
        public readonly int? AppliedChanges;
        /// <summary>
        /// CDC delete counter
        /// </summary>
        public readonly int? CdcDeleteCounter;
        /// <summary>
        /// CDC insert counter
        /// </summary>
        public readonly int? CdcInsertCounter;
        /// <summary>
        /// CDC update counter
        /// </summary>
        public readonly int? CdcUpdateCounter;
        /// <summary>
        /// Name of the database
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// End date-time of a migration state
        /// </summary>
        public readonly string? EndedOn;
        /// <summary>
        /// Number of tables loaded during the migration of a DB
        /// </summary>
        public readonly int? FullLoadCompletedTables;
        /// <summary>
        /// Number of tables errored out during the migration of a DB
        /// </summary>
        public readonly int? FullLoadErroredTables;
        /// <summary>
        /// Number of tables loading during the migration of a DB
        /// </summary>
        public readonly int? FullLoadLoadingTables;
        /// <summary>
        /// Number of tables queued for the migration of a DB
        /// </summary>
        public readonly int? FullLoadQueuedTables;
        /// <summary>
        /// CDC incoming changes counter
        /// </summary>
        public readonly int? IncomingChanges;
        /// <summary>
        /// Lag in seconds between source and target during online phase
        /// </summary>
        public readonly int? Latency;
        /// <summary>
        /// Error message, if any, for the migration state
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Migration operation of an individual database
        /// </summary>
        public readonly string? MigrationOperation;
        /// <summary>
        /// Migration db state of an individual database
        /// </summary>
        public readonly string? MigrationState;
        /// <summary>
        /// Start date-time of a migration state
        /// </summary>
        public readonly string? StartedOn;

        [OutputConstructor]
        private DbMigrationStatusResponse(
            int? appliedChanges,

            int? cdcDeleteCounter,

            int? cdcInsertCounter,

            int? cdcUpdateCounter,

            string? databaseName,

            string? endedOn,

            int? fullLoadCompletedTables,

            int? fullLoadErroredTables,

            int? fullLoadLoadingTables,

            int? fullLoadQueuedTables,

            int? incomingChanges,

            int? latency,

            string? message,

            string? migrationOperation,

            string? migrationState,

            string? startedOn)
        {
            AppliedChanges = appliedChanges;
            CdcDeleteCounter = cdcDeleteCounter;
            CdcInsertCounter = cdcInsertCounter;
            CdcUpdateCounter = cdcUpdateCounter;
            DatabaseName = databaseName;
            EndedOn = endedOn;
            FullLoadCompletedTables = fullLoadCompletedTables;
            FullLoadErroredTables = fullLoadErroredTables;
            FullLoadLoadingTables = fullLoadLoadingTables;
            FullLoadQueuedTables = fullLoadQueuedTables;
            IncomingChanges = incomingChanges;
            Latency = latency;
            Message = message;
            MigrationOperation = migrationOperation;
            MigrationState = migrationState;
            StartedOn = startedOn;
        }
    }
}
