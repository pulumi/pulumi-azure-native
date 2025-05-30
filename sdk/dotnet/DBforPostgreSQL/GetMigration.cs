// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DBforPostgreSQL
{
    public static class GetMigration
    {
        /// <summary>
        /// Gets details of a migration.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetMigrationResult> InvokeAsync(GetMigrationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMigrationResult>("azure-native:dbforpostgresql:getMigration", args ?? new GetMigrationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a migration.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMigrationResult> Invoke(GetMigrationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMigrationResult>("azure-native:dbforpostgresql:getMigration", args ?? new GetMigrationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a migration.
        /// 
        /// Uses Azure REST API version 2024-08-01.
        /// 
        /// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetMigrationResult> Invoke(GetMigrationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMigrationResult>("azure-native:dbforpostgresql:getMigration", args ?? new GetMigrationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMigrationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the migration.
        /// </summary>
        [Input("migrationName", required: true)]
        public string MigrationName { get; set; } = null!;

        /// <summary>
        /// The resource group name of the target database server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The subscription ID of the target database server.
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        /// <summary>
        /// The name of the target database server.
        /// </summary>
        [Input("targetDbServerName", required: true)]
        public string TargetDbServerName { get; set; } = null!;

        public GetMigrationArgs()
        {
        }
        public static new GetMigrationArgs Empty => new GetMigrationArgs();
    }

    public sealed class GetMigrationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the migration.
        /// </summary>
        [Input("migrationName", required: true)]
        public Input<string> MigrationName { get; set; } = null!;

        /// <summary>
        /// The resource group name of the target database server.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The subscription ID of the target database server.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The name of the target database server.
        /// </summary>
        [Input("targetDbServerName", required: true)]
        public Input<string> TargetDbServerName { get; set; } = null!;

        public GetMigrationInvokeArgs()
        {
        }
        public static new GetMigrationInvokeArgs Empty => new GetMigrationInvokeArgs();
    }


    [OutputType]
    public sealed class GetMigrationResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// To trigger cancel for entire migration we need to send this flag as True
        /// </summary>
        public readonly string? Cancel;
        /// <summary>
        /// Current status of migration
        /// </summary>
        public readonly Outputs.MigrationStatusResponse CurrentStatus;
        /// <summary>
        /// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
        /// </summary>
        public readonly ImmutableArray<string> DbsToCancelMigrationOn;
        /// <summary>
        /// Number of databases to migrate
        /// </summary>
        public readonly ImmutableArray<string> DbsToMigrate;
        /// <summary>
        /// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
        /// </summary>
        public readonly ImmutableArray<string> DbsToTriggerCutoverOn;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// To migrate roles and permissions we need to send this flag as True
        /// </summary>
        public readonly string? MigrateRoles;
        /// <summary>
        /// ID for migration, a GUID.
        /// </summary>
        public readonly string MigrationId;
        /// <summary>
        /// ResourceId of the private endpoint migration instance
        /// </summary>
        public readonly string? MigrationInstanceResourceId;
        /// <summary>
        /// There are two types of migration modes Online and Offline
        /// </summary>
        public readonly string? MigrationMode;
        /// <summary>
        /// This indicates the supported Migration option for the migration
        /// </summary>
        public readonly string? MigrationOption;
        /// <summary>
        /// End time in UTC for migration window
        /// </summary>
        public readonly string? MigrationWindowEndTimeInUtc;
        /// <summary>
        /// Start time in UTC for migration window
        /// </summary>
        public readonly string? MigrationWindowStartTimeInUtc;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
        /// </summary>
        public readonly string? OverwriteDbsInTarget;
        /// <summary>
        /// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
        /// </summary>
        public readonly string? SetupLogicalReplicationOnSourceDbIfNeeded;
        /// <summary>
        /// Source server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
        /// </summary>
        public readonly string? SourceDbServerFullyQualifiedDomainName;
        /// <summary>
        /// Metadata of the source database server
        /// </summary>
        public readonly Outputs.DbServerMetadataResponse SourceDbServerMetadata;
        /// <summary>
        /// ResourceId of the source database server in case the sourceType is PostgreSQLSingleServer. For other source types this should be ipaddress:port@username or hostname:port@username
        /// </summary>
        public readonly string? SourceDbServerResourceId;
        /// <summary>
        /// migration source server type : OnPremises, AWS, GCP, AzureVM, PostgreSQLSingleServer, AWS_RDS, AWS_AURORA, AWS_EC2, GCP_CloudSQL, GCP_AlloyDB, GCP_Compute, or EDB
        /// </summary>
        public readonly string? SourceType;
        /// <summary>
        /// SSL modes for migration. Default SSL mode for PostgreSQLSingleServer is VerifyFull and Prefer for other source types
        /// </summary>
        public readonly string? SslMode;
        /// <summary>
        /// Indicates whether the data migration should start right away
        /// </summary>
        public readonly string? StartDataMigration;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Target server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
        /// </summary>
        public readonly string? TargetDbServerFullyQualifiedDomainName;
        /// <summary>
        /// Metadata of the target database server
        /// </summary>
        public readonly Outputs.DbServerMetadataResponse TargetDbServerMetadata;
        /// <summary>
        /// ResourceId of the source database server
        /// </summary>
        public readonly string TargetDbServerResourceId;
        /// <summary>
        /// To trigger cutover for entire migration we need to send this flag as True
        /// </summary>
        public readonly string? TriggerCutover;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMigrationResult(
            string azureApiVersion,

            string? cancel,

            Outputs.MigrationStatusResponse currentStatus,

            ImmutableArray<string> dbsToCancelMigrationOn,

            ImmutableArray<string> dbsToMigrate,

            ImmutableArray<string> dbsToTriggerCutoverOn,

            string id,

            string location,

            string? migrateRoles,

            string migrationId,

            string? migrationInstanceResourceId,

            string? migrationMode,

            string? migrationOption,

            string? migrationWindowEndTimeInUtc,

            string? migrationWindowStartTimeInUtc,

            string name,

            string? overwriteDbsInTarget,

            string? setupLogicalReplicationOnSourceDbIfNeeded,

            string? sourceDbServerFullyQualifiedDomainName,

            Outputs.DbServerMetadataResponse sourceDbServerMetadata,

            string? sourceDbServerResourceId,

            string? sourceType,

            string? sslMode,

            string? startDataMigration,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string? targetDbServerFullyQualifiedDomainName,

            Outputs.DbServerMetadataResponse targetDbServerMetadata,

            string targetDbServerResourceId,

            string? triggerCutover,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Cancel = cancel;
            CurrentStatus = currentStatus;
            DbsToCancelMigrationOn = dbsToCancelMigrationOn;
            DbsToMigrate = dbsToMigrate;
            DbsToTriggerCutoverOn = dbsToTriggerCutoverOn;
            Id = id;
            Location = location;
            MigrateRoles = migrateRoles;
            MigrationId = migrationId;
            MigrationInstanceResourceId = migrationInstanceResourceId;
            MigrationMode = migrationMode;
            MigrationOption = migrationOption;
            MigrationWindowEndTimeInUtc = migrationWindowEndTimeInUtc;
            MigrationWindowStartTimeInUtc = migrationWindowStartTimeInUtc;
            Name = name;
            OverwriteDbsInTarget = overwriteDbsInTarget;
            SetupLogicalReplicationOnSourceDbIfNeeded = setupLogicalReplicationOnSourceDbIfNeeded;
            SourceDbServerFullyQualifiedDomainName = sourceDbServerFullyQualifiedDomainName;
            SourceDbServerMetadata = sourceDbServerMetadata;
            SourceDbServerResourceId = sourceDbServerResourceId;
            SourceType = sourceType;
            SslMode = sslMode;
            StartDataMigration = startDataMigration;
            SystemData = systemData;
            Tags = tags;
            TargetDbServerFullyQualifiedDomainName = targetDbServerFullyQualifiedDomainName;
            TargetDbServerMetadata = targetDbServerMetadata;
            TargetDbServerResourceId = targetDbServerResourceId;
            TriggerCutover = triggerCutover;
            Type = type;
        }
    }
}
