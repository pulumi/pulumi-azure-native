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
    /// Database level output for the task that validates connection to SQL Server and also validates source server requirements
    /// </summary>
    [OutputType]
    public sealed class ConnectToSourceSqlServerTaskOutputDatabaseLevelResponse
    {
        /// <summary>
        /// SQL Server compatibility level of database
        /// </summary>
        public readonly string CompatibilityLevel;
        /// <summary>
        /// The list of database files
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseFileInfoResponse> DatabaseFiles;
        /// <summary>
        /// State of the database
        /// </summary>
        public readonly string DatabaseState;
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Database name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of result - database level or task level
        /// Expected value is 'DatabaseLevelOutput'.
        /// </summary>
        public readonly string ResultType;
        /// <summary>
        /// Size of the file in megabytes
        /// </summary>
        public readonly double SizeMB;

        [OutputConstructor]
        private ConnectToSourceSqlServerTaskOutputDatabaseLevelResponse(
            string compatibilityLevel,

            ImmutableArray<Outputs.DatabaseFileInfoResponse> databaseFiles,

            string databaseState,

            string id,

            string name,

            string resultType,

            double sizeMB)
        {
            CompatibilityLevel = compatibilityLevel;
            DatabaseFiles = databaseFiles;
            DatabaseState = databaseState;
            Id = id;
            Name = name;
            ResultType = resultType;
            SizeMB = sizeMB;
        }
    }
}
