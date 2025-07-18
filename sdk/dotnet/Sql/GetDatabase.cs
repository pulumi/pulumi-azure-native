// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    public static class GetDatabase
    {
        /// <summary>
        /// Gets a database.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2014-04-01, 2017-03-01-preview, 2017-10-01-preview, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("azure-native:sql:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a database.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2014-04-01, 2017-03-01-preview, 2017-10-01-preview, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("azure-native:sql:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a database.
        /// 
        /// Uses Azure REST API version 2023-08-01.
        /// 
        /// Other available API versions: 2014-04-01, 2017-03-01-preview, 2017-10-01-preview, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("azure-native:sql:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// An OData filter expression that filters elements in the collection.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
        public static new GetDatabaseArgs Empty => new GetDatabaseArgs();
    }

    public sealed class GetDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The child resources to include in the response.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// An OData filter expression that filters elements in the collection.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public GetDatabaseInvokeArgs()
        {
        }
        public static new GetDatabaseInvokeArgs Empty => new GetDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
        /// </summary>
        public readonly int? AutoPauseDelay;
        /// <summary>
        /// Specifies the availability zone the database is pinned to.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Collation of the metadata catalog.
        /// </summary>
        public readonly string? CatalogCollation;
        /// <summary>
        /// The collation of the database.
        /// </summary>
        public readonly string? Collation;
        /// <summary>
        /// The creation date of the database (ISO8601 format).
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The storage account type used to store backups for this database.
        /// </summary>
        public readonly string CurrentBackupStorageRedundancy;
        /// <summary>
        /// The current service level objective name of the database.
        /// </summary>
        public readonly string CurrentServiceObjectiveName;
        /// <summary>
        /// The name and tier of the SKU.
        /// </summary>
        public readonly Outputs.SkuResponse CurrentSku;
        /// <summary>
        /// The ID of the database.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// The default secondary region for this database.
        /// </summary>
        public readonly string DefaultSecondaryLocation;
        /// <summary>
        /// This records the earliest start date and time that restore is available for this database (ISO8601 format).
        /// </summary>
        public readonly string EarliestRestoreDate;
        /// <summary>
        /// The resource identifier of the elastic pool containing this database.
        /// </summary>
        public readonly string? ElasticPoolId;
        /// <summary>
        /// The azure key vault URI of the database if it's configured with per Database Customer Managed Keys.
        /// </summary>
        public readonly string? EncryptionProtector;
        /// <summary>
        /// The flag to enable or disable auto rotation of database encryption protector AKV key.
        /// </summary>
        public readonly bool? EncryptionProtectorAutoRotation;
        /// <summary>
        /// Failover Group resource identifier that this database belongs to.
        /// </summary>
        public readonly string FailoverGroupId;
        /// <summary>
        /// The Client id used for cross tenant per database CMK scenario
        /// </summary>
        public readonly string? FederatedClientId;
        /// <summary>
        /// Specifies the behavior when monthly free limits are exhausted for the free database.
        /// 
        /// AutoPause: The database will be auto paused upon exhaustion of free limits for remainder of the month.
        /// 
        /// BillForUsage: The database will continue to be online upon exhaustion of free limits and any overage will be billed.
        /// </summary>
        public readonly string? FreeLimitExhaustionBehavior;
        /// <summary>
        /// The number of secondary replicas associated with the Business Critical, Premium, or Hyperscale edition database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
        /// </summary>
        public readonly int? HighAvailabilityReplicaCount;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Active Directory identity of the database.
        /// </summary>
        public readonly Outputs.DatabaseIdentityResponse? Identity;
        /// <summary>
        /// Infra encryption is enabled for this database.
        /// </summary>
        public readonly bool IsInfraEncryptionEnabled;
        /// <summary>
        /// Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
        /// </summary>
        public readonly bool? IsLedgerOn;
        /// <summary>
        /// The resource ids of the user assigned identities to use
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.DatabaseKeyResponse>? Keys;
        /// <summary>
        /// Kind of database. This is metadata used for the Azure portal experience.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
        /// </summary>
        public readonly string? MaintenanceConfigurationId;
        /// <summary>
        /// Resource that manages the database.
        /// </summary>
        public readonly string ManagedBy;
        /// <summary>
        /// Whether or not customer controlled manual cutover needs to be done during Update Database operation to Hyperscale tier.
        /// 
        /// This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier.
        /// 
        /// When manualCutover is specified, the scaling operation will wait for user input to trigger cutover to Hyperscale database.
        /// 
        /// To trigger cutover, please provide 'performCutover' parameter when the Scaling operation is in Waiting state.
        /// </summary>
        public readonly bool? ManualCutover;
        /// <summary>
        /// The max log size for this database.
        /// </summary>
        public readonly double MaxLogSizeBytes;
        /// <summary>
        /// The max size of the database expressed in bytes.
        /// </summary>
        public readonly double? MaxSizeBytes;
        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused
        /// </summary>
        public readonly double? MinCapacity;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
        /// </summary>
        public readonly string PausedDate;
        /// <summary>
        /// To trigger customer controlled manual cutover during the wait state while Scaling operation is in progress.
        /// 
        /// This property parameter is only applicable for scaling operations that are initiated along with 'manualCutover' parameter.
        /// 
        /// This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier is already in progress.
        /// 
        /// When performCutover is specified, the scaling operation will trigger cutover and perform role-change to Hyperscale database.
        /// </summary>
        public readonly bool? PerformCutover;
        /// <summary>
        /// Type of enclave requested on the database i.e. Default or VBS enclaves.
        /// </summary>
        public readonly string? PreferredEnclaveType;
        /// <summary>
        /// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
        /// </summary>
        public readonly string? ReadScale;
        /// <summary>
        /// The storage account type to be used to store backups for this database.
        /// </summary>
        public readonly string? RequestedBackupStorageRedundancy;
        /// <summary>
        /// The requested service level objective name of the database.
        /// </summary>
        public readonly string RequestedServiceObjectiveName;
        /// <summary>
        /// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
        /// </summary>
        public readonly string ResumedDate;
        /// <summary>
        /// The secondary type of the database if it is a secondary.  Valid values are Geo, Named and Standby.
        /// </summary>
        public readonly string? SecondaryType;
        /// <summary>
        /// The database SKU.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        /// 
        /// ```azurecli
        /// az sql db list-editions -l &lt;location&gt; -o table
        /// ````
        /// 
        /// ```powershell
        /// Get-AzSqlServerServiceObjective -Location &lt;location&gt;
        /// ````
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The status of the database.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Whether or not the database uses free monthly limits. Allowed on one database in a subscription.
        /// </summary>
        public readonly bool? UseFreeLimit;
        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        /// </summary>
        public readonly bool? ZoneRedundant;

        [OutputConstructor]
        private GetDatabaseResult(
            int? autoPauseDelay,

            string? availabilityZone,

            string azureApiVersion,

            string? catalogCollation,

            string? collation,

            string creationDate,

            string currentBackupStorageRedundancy,

            string currentServiceObjectiveName,

            Outputs.SkuResponse currentSku,

            string databaseId,

            string defaultSecondaryLocation,

            string earliestRestoreDate,

            string? elasticPoolId,

            string? encryptionProtector,

            bool? encryptionProtectorAutoRotation,

            string failoverGroupId,

            string? federatedClientId,

            string? freeLimitExhaustionBehavior,

            int? highAvailabilityReplicaCount,

            string id,

            Outputs.DatabaseIdentityResponse? identity,

            bool isInfraEncryptionEnabled,

            bool? isLedgerOn,

            ImmutableDictionary<string, Outputs.DatabaseKeyResponse>? keys,

            string kind,

            string? licenseType,

            string location,

            string? maintenanceConfigurationId,

            string managedBy,

            bool? manualCutover,

            double maxLogSizeBytes,

            double? maxSizeBytes,

            double? minCapacity,

            string name,

            string pausedDate,

            bool? performCutover,

            string? preferredEnclaveType,

            string? readScale,

            string? requestedBackupStorageRedundancy,

            string requestedServiceObjectiveName,

            string resumedDate,

            string? secondaryType,

            Outputs.SkuResponse? sku,

            string status,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? useFreeLimit,

            bool? zoneRedundant)
        {
            AutoPauseDelay = autoPauseDelay;
            AvailabilityZone = availabilityZone;
            AzureApiVersion = azureApiVersion;
            CatalogCollation = catalogCollation;
            Collation = collation;
            CreationDate = creationDate;
            CurrentBackupStorageRedundancy = currentBackupStorageRedundancy;
            CurrentServiceObjectiveName = currentServiceObjectiveName;
            CurrentSku = currentSku;
            DatabaseId = databaseId;
            DefaultSecondaryLocation = defaultSecondaryLocation;
            EarliestRestoreDate = earliestRestoreDate;
            ElasticPoolId = elasticPoolId;
            EncryptionProtector = encryptionProtector;
            EncryptionProtectorAutoRotation = encryptionProtectorAutoRotation;
            FailoverGroupId = failoverGroupId;
            FederatedClientId = federatedClientId;
            FreeLimitExhaustionBehavior = freeLimitExhaustionBehavior;
            HighAvailabilityReplicaCount = highAvailabilityReplicaCount;
            Id = id;
            Identity = identity;
            IsInfraEncryptionEnabled = isInfraEncryptionEnabled;
            IsLedgerOn = isLedgerOn;
            Keys = keys;
            Kind = kind;
            LicenseType = licenseType;
            Location = location;
            MaintenanceConfigurationId = maintenanceConfigurationId;
            ManagedBy = managedBy;
            ManualCutover = manualCutover;
            MaxLogSizeBytes = maxLogSizeBytes;
            MaxSizeBytes = maxSizeBytes;
            MinCapacity = minCapacity;
            Name = name;
            PausedDate = pausedDate;
            PerformCutover = performCutover;
            PreferredEnclaveType = preferredEnclaveType;
            ReadScale = readScale;
            RequestedBackupStorageRedundancy = requestedBackupStorageRedundancy;
            RequestedServiceObjectiveName = requestedServiceObjectiveName;
            ResumedDate = resumedDate;
            SecondaryType = secondaryType;
            Sku = sku;
            Status = status;
            Tags = tags;
            Type = type;
            UseFreeLimit = useFreeLimit;
            ZoneRedundant = zoneRedundant;
        }
    }
}
