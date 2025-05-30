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
    /// Database Migration Resource properties for SQL database.
    /// </summary>
    [OutputType]
    public sealed class DatabaseMigrationPropertiesSqlDbResponse
    {
        /// <summary>
        /// Database migration end time.
        /// </summary>
        public readonly string EndedOn;
        /// <summary>
        /// 
        /// Expected value is 'SqlDb'.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Error details in case of migration failure.
        /// </summary>
        public readonly Outputs.ErrorInfoResponse MigrationFailureError;
        /// <summary>
        /// ID for current migration operation.
        /// </summary>
        public readonly string? MigrationOperationId;
        /// <summary>
        /// Resource Id of the Migration Service.
        /// </summary>
        public readonly string? MigrationService;
        /// <summary>
        /// Migration status.
        /// </summary>
        public readonly string MigrationStatus;
        /// <summary>
        /// Detailed migration status. Not included by default.
        /// </summary>
        public readonly Outputs.SqlDbMigrationStatusDetailsResponse MigrationStatusDetails;
        /// <summary>
        /// Offline configuration.
        /// </summary>
        public readonly Outputs.SqlDbOfflineConfigurationResponse OfflineConfiguration;
        /// <summary>
        /// Error message for migration provisioning failure, if any.
        /// </summary>
        public readonly string? ProvisioningError;
        /// <summary>
        /// Provisioning State of migration. ProvisioningState as Succeeded implies that validations have been performed and migration has started.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource Id of the target resource.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Name of the source database.
        /// </summary>
        public readonly string? SourceDatabaseName;
        /// <summary>
        /// Name of the source sql server.
        /// </summary>
        public readonly string SourceServerName;
        /// <summary>
        /// Source SQL Server connection details.
        /// </summary>
        public readonly Outputs.SqlConnectionInformationResponse? SourceSqlConnection;
        /// <summary>
        /// Database migration start time.
        /// </summary>
        public readonly string StartedOn;
        /// <summary>
        /// List of tables to copy.
        /// </summary>
        public readonly ImmutableArray<string> TableList;
        /// <summary>
        /// Database collation to be used for the target database.
        /// </summary>
        public readonly string? TargetDatabaseCollation;
        /// <summary>
        /// Target SQL DB connection details.
        /// </summary>
        public readonly Outputs.SqlConnectionInformationResponse? TargetSqlConnection;

        [OutputConstructor]
        private DatabaseMigrationPropertiesSqlDbResponse(
            string endedOn,

            string kind,

            Outputs.ErrorInfoResponse migrationFailureError,

            string? migrationOperationId,

            string? migrationService,

            string migrationStatus,

            Outputs.SqlDbMigrationStatusDetailsResponse migrationStatusDetails,

            Outputs.SqlDbOfflineConfigurationResponse offlineConfiguration,

            string? provisioningError,

            string provisioningState,

            string? scope,

            string? sourceDatabaseName,

            string sourceServerName,

            Outputs.SqlConnectionInformationResponse? sourceSqlConnection,

            string startedOn,

            ImmutableArray<string> tableList,

            string? targetDatabaseCollation,

            Outputs.SqlConnectionInformationResponse? targetSqlConnection)
        {
            EndedOn = endedOn;
            Kind = kind;
            MigrationFailureError = migrationFailureError;
            MigrationOperationId = migrationOperationId;
            MigrationService = migrationService;
            MigrationStatus = migrationStatus;
            MigrationStatusDetails = migrationStatusDetails;
            OfflineConfiguration = offlineConfiguration;
            ProvisioningError = provisioningError;
            ProvisioningState = provisioningState;
            Scope = scope;
            SourceDatabaseName = sourceDatabaseName;
            SourceServerName = sourceServerName;
            SourceSqlConnection = sourceSqlConnection;
            StartedOn = startedOn;
            TableList = tableList;
            TargetDatabaseCollation = targetDatabaseCollation;
            TargetSqlConnection = targetSqlConnection;
        }
    }
}
