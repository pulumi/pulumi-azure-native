// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A database resource.
 *
 * Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
 *
 * Other available API versions: 2014-04-01, 2017-03-01-preview, 2017-10-01-preview, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
     */
    public readonly autoPauseDelay!: pulumi.Output<number | undefined>;
    /**
     * Specifies the availability zone the database is pinned to.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Collation of the metadata catalog.
     */
    public readonly catalogCollation!: pulumi.Output<string | undefined>;
    /**
     * The collation of the database.
     */
    public readonly collation!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the database (ISO8601 format).
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The storage account type used to store backups for this database.
     */
    public /*out*/ readonly currentBackupStorageRedundancy!: pulumi.Output<string>;
    /**
     * The current service level objective name of the database.
     */
    public /*out*/ readonly currentServiceObjectiveName!: pulumi.Output<string>;
    /**
     * The name and tier of the SKU.
     */
    public /*out*/ readonly currentSku!: pulumi.Output<outputs.sql.SkuResponse>;
    /**
     * The ID of the database.
     */
    public /*out*/ readonly databaseId!: pulumi.Output<string>;
    /**
     * The default secondary region for this database.
     */
    public /*out*/ readonly defaultSecondaryLocation!: pulumi.Output<string>;
    /**
     * This records the earliest start date and time that restore is available for this database (ISO8601 format).
     */
    public /*out*/ readonly earliestRestoreDate!: pulumi.Output<string>;
    /**
     * The resource identifier of the elastic pool containing this database.
     */
    public readonly elasticPoolId!: pulumi.Output<string | undefined>;
    /**
     * The azure key vault URI of the database if it's configured with per Database Customer Managed Keys.
     */
    public readonly encryptionProtector!: pulumi.Output<string | undefined>;
    /**
     * The flag to enable or disable auto rotation of database encryption protector AKV key.
     */
    public readonly encryptionProtectorAutoRotation!: pulumi.Output<boolean | undefined>;
    /**
     * Failover Group resource identifier that this database belongs to.
     */
    public /*out*/ readonly failoverGroupId!: pulumi.Output<string>;
    /**
     * The Client id used for cross tenant per database CMK scenario
     */
    public readonly federatedClientId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the behavior when monthly free limits are exhausted for the free database.
     * 
     * AutoPause: The database will be auto paused upon exhaustion of free limits for remainder of the month.
     * 
     * BillForUsage: The database will continue to be online upon exhaustion of free limits and any overage will be billed.
     */
    public readonly freeLimitExhaustionBehavior!: pulumi.Output<string | undefined>;
    /**
     * The number of secondary replicas associated with the Business Critical, Premium, or Hyperscale edition database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
     */
    public readonly highAvailabilityReplicaCount!: pulumi.Output<number | undefined>;
    /**
     * The Azure Active Directory identity of the database.
     */
    public readonly identity!: pulumi.Output<outputs.sql.DatabaseIdentityResponse | undefined>;
    /**
     * Infra encryption is enabled for this database.
     */
    public /*out*/ readonly isInfraEncryptionEnabled!: pulumi.Output<boolean>;
    /**
     * Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
     */
    public readonly isLedgerOn!: pulumi.Output<boolean | undefined>;
    /**
     * The resource ids of the user assigned identities to use
     */
    public readonly keys!: pulumi.Output<{[key: string]: outputs.sql.DatabaseKeyResponse} | undefined>;
    /**
     * Kind of database. This is metadata used for the Azure portal experience.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
     */
    public readonly maintenanceConfigurationId!: pulumi.Output<string | undefined>;
    /**
     * Resource that manages the database.
     */
    public /*out*/ readonly managedBy!: pulumi.Output<string>;
    /**
     * Whether or not customer controlled manual cutover needs to be done during Update Database operation to Hyperscale tier.
     * 
     * This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier.
     * 
     * When manualCutover is specified, the scaling operation will wait for user input to trigger cutover to Hyperscale database.
     * 
     * To trigger cutover, please provide 'performCutover' parameter when the Scaling operation is in Waiting state.
     */
    public readonly manualCutover!: pulumi.Output<boolean | undefined>;
    /**
     * The max log size for this database.
     */
    public /*out*/ readonly maxLogSizeBytes!: pulumi.Output<number>;
    /**
     * The max size of the database expressed in bytes.
     */
    public readonly maxSizeBytes!: pulumi.Output<number | undefined>;
    /**
     * Minimal capacity that database will always have allocated, if not paused
     */
    public readonly minCapacity!: pulumi.Output<number | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
     */
    public /*out*/ readonly pausedDate!: pulumi.Output<string>;
    /**
     * To trigger customer controlled manual cutover during the wait state while Scaling operation is in progress.
     * 
     * This property parameter is only applicable for scaling operations that are initiated along with 'manualCutover' parameter.
     * 
     * This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier is already in progress.
     * 
     * When performCutover is specified, the scaling operation will trigger cutover and perform role-change to Hyperscale database.
     */
    public readonly performCutover!: pulumi.Output<boolean | undefined>;
    /**
     * Type of enclave requested on the database i.e. Default or VBS enclaves.
     */
    public readonly preferredEnclaveType!: pulumi.Output<string | undefined>;
    /**
     * The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
     */
    public readonly readScale!: pulumi.Output<string | undefined>;
    /**
     * The storage account type to be used to store backups for this database.
     */
    public readonly requestedBackupStorageRedundancy!: pulumi.Output<string | undefined>;
    /**
     * The requested service level objective name of the database.
     */
    public /*out*/ readonly requestedServiceObjectiveName!: pulumi.Output<string>;
    /**
     * The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
     */
    public /*out*/ readonly resumedDate!: pulumi.Output<string>;
    /**
     * The secondary type of the database if it is a secondary.  Valid values are Geo, Named and Standby.
     */
    public readonly secondaryType!: pulumi.Output<string | undefined>;
    /**
     * The database SKU.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    public readonly sku!: pulumi.Output<outputs.sql.SkuResponse | undefined>;
    /**
     * The status of the database.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Whether or not the database uses free monthly limits. Allowed on one database in a subscription.
     */
    public readonly useFreeLimit!: pulumi.Output<boolean | undefined>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    public readonly zoneRedundant!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            resourceInputs["autoPauseDelay"] = args ? args.autoPauseDelay : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["catalogCollation"] = args ? args.catalogCollation : undefined;
            resourceInputs["collation"] = args ? args.collation : undefined;
            resourceInputs["createMode"] = args ? args.createMode : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["elasticPoolId"] = args ? args.elasticPoolId : undefined;
            resourceInputs["encryptionProtector"] = args ? args.encryptionProtector : undefined;
            resourceInputs["encryptionProtectorAutoRotation"] = args ? args.encryptionProtectorAutoRotation : undefined;
            resourceInputs["federatedClientId"] = args ? args.federatedClientId : undefined;
            resourceInputs["freeLimitExhaustionBehavior"] = args ? args.freeLimitExhaustionBehavior : undefined;
            resourceInputs["highAvailabilityReplicaCount"] = args ? args.highAvailabilityReplicaCount : undefined;
            resourceInputs["identity"] = args ? args.identity : undefined;
            resourceInputs["isLedgerOn"] = args ? args.isLedgerOn : undefined;
            resourceInputs["keys"] = args ? args.keys : undefined;
            resourceInputs["licenseType"] = args ? args.licenseType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["longTermRetentionBackupResourceId"] = args ? args.longTermRetentionBackupResourceId : undefined;
            resourceInputs["maintenanceConfigurationId"] = args ? args.maintenanceConfigurationId : undefined;
            resourceInputs["manualCutover"] = args ? args.manualCutover : undefined;
            resourceInputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            resourceInputs["minCapacity"] = args ? args.minCapacity : undefined;
            resourceInputs["performCutover"] = args ? args.performCutover : undefined;
            resourceInputs["preferredEnclaveType"] = args ? args.preferredEnclaveType : undefined;
            resourceInputs["readScale"] = args ? args.readScale : undefined;
            resourceInputs["recoverableDatabaseId"] = args ? args.recoverableDatabaseId : undefined;
            resourceInputs["recoveryServicesRecoveryPointId"] = args ? args.recoveryServicesRecoveryPointId : undefined;
            resourceInputs["requestedBackupStorageRedundancy"] = args ? args.requestedBackupStorageRedundancy : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["restorableDroppedDatabaseId"] = args ? args.restorableDroppedDatabaseId : undefined;
            resourceInputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            resourceInputs["sampleName"] = args ? args.sampleName : undefined;
            resourceInputs["secondaryType"] = args ? args.secondaryType : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["sourceDatabaseDeletionDate"] = args ? args.sourceDatabaseDeletionDate : undefined;
            resourceInputs["sourceDatabaseId"] = args ? args.sourceDatabaseId : undefined;
            resourceInputs["sourceResourceId"] = args ? args.sourceResourceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useFreeLimit"] = args ? args.useFreeLimit : undefined;
            resourceInputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["currentBackupStorageRedundancy"] = undefined /*out*/;
            resourceInputs["currentServiceObjectiveName"] = undefined /*out*/;
            resourceInputs["currentSku"] = undefined /*out*/;
            resourceInputs["databaseId"] = undefined /*out*/;
            resourceInputs["defaultSecondaryLocation"] = undefined /*out*/;
            resourceInputs["earliestRestoreDate"] = undefined /*out*/;
            resourceInputs["failoverGroupId"] = undefined /*out*/;
            resourceInputs["isInfraEncryptionEnabled"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["managedBy"] = undefined /*out*/;
            resourceInputs["maxLogSizeBytes"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pausedDate"] = undefined /*out*/;
            resourceInputs["requestedServiceObjectiveName"] = undefined /*out*/;
            resourceInputs["resumedDate"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["autoPauseDelay"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["catalogCollation"] = undefined /*out*/;
            resourceInputs["collation"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["currentBackupStorageRedundancy"] = undefined /*out*/;
            resourceInputs["currentServiceObjectiveName"] = undefined /*out*/;
            resourceInputs["currentSku"] = undefined /*out*/;
            resourceInputs["databaseId"] = undefined /*out*/;
            resourceInputs["defaultSecondaryLocation"] = undefined /*out*/;
            resourceInputs["earliestRestoreDate"] = undefined /*out*/;
            resourceInputs["elasticPoolId"] = undefined /*out*/;
            resourceInputs["encryptionProtector"] = undefined /*out*/;
            resourceInputs["encryptionProtectorAutoRotation"] = undefined /*out*/;
            resourceInputs["failoverGroupId"] = undefined /*out*/;
            resourceInputs["federatedClientId"] = undefined /*out*/;
            resourceInputs["freeLimitExhaustionBehavior"] = undefined /*out*/;
            resourceInputs["highAvailabilityReplicaCount"] = undefined /*out*/;
            resourceInputs["identity"] = undefined /*out*/;
            resourceInputs["isInfraEncryptionEnabled"] = undefined /*out*/;
            resourceInputs["isLedgerOn"] = undefined /*out*/;
            resourceInputs["keys"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["licenseType"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["maintenanceConfigurationId"] = undefined /*out*/;
            resourceInputs["managedBy"] = undefined /*out*/;
            resourceInputs["manualCutover"] = undefined /*out*/;
            resourceInputs["maxLogSizeBytes"] = undefined /*out*/;
            resourceInputs["maxSizeBytes"] = undefined /*out*/;
            resourceInputs["minCapacity"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pausedDate"] = undefined /*out*/;
            resourceInputs["performCutover"] = undefined /*out*/;
            resourceInputs["preferredEnclaveType"] = undefined /*out*/;
            resourceInputs["readScale"] = undefined /*out*/;
            resourceInputs["requestedBackupStorageRedundancy"] = undefined /*out*/;
            resourceInputs["requestedServiceObjectiveName"] = undefined /*out*/;
            resourceInputs["resumedDate"] = undefined /*out*/;
            resourceInputs["secondaryType"] = undefined /*out*/;
            resourceInputs["sku"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["useFreeLimit"] = undefined /*out*/;
            resourceInputs["zoneRedundant"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:sql/v20140401:Database" }, { type: "azure-native:sql/v20170301preview:Database" }, { type: "azure-native:sql/v20171001preview:Database" }, { type: "azure-native:sql/v20190601preview:Database" }, { type: "azure-native:sql/v20200202preview:Database" }, { type: "azure-native:sql/v20200801preview:Database" }, { type: "azure-native:sql/v20201101preview:Database" }, { type: "azure-native:sql/v20210201preview:Database" }, { type: "azure-native:sql/v20210501preview:Database" }, { type: "azure-native:sql/v20210801preview:Database" }, { type: "azure-native:sql/v20211101:Database" }, { type: "azure-native:sql/v20211101preview:Database" }, { type: "azure-native:sql/v20220201preview:Database" }, { type: "azure-native:sql/v20220501preview:Database" }, { type: "azure-native:sql/v20220801preview:Database" }, { type: "azure-native:sql/v20221101preview:Database" }, { type: "azure-native:sql/v20230201preview:Database" }, { type: "azure-native:sql/v20230501preview:Database" }, { type: "azure-native:sql/v20230801:Database" }, { type: "azure-native:sql/v20230801preview:Database" }, { type: "azure-native:sql/v20240501preview:Database" }, { type: "azure-native:sql/v20241101preview:Database" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
     */
    autoPauseDelay?: pulumi.Input<number>;
    /**
     * Specifies the availability zone the database is pinned to.
     */
    availabilityZone?: pulumi.Input<string | enums.sql.AvailabilityZoneType>;
    /**
     * Collation of the metadata catalog.
     */
    catalogCollation?: pulumi.Input<string | enums.sql.CatalogCollationType>;
    /**
     * The collation of the database.
     */
    collation?: pulumi.Input<string>;
    /**
     * Specifies the mode of database creation.
     * 
     * Default: regular database creation.
     * 
     * Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
     * 
     * Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
     * 
     * PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
     * 
     * Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
     * 
     * Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
     * 
     * RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
     * 
     * Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
     */
    createMode?: pulumi.Input<string | enums.sql.CreateMode>;
    /**
     * The name of the database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The resource identifier of the elastic pool containing this database.
     */
    elasticPoolId?: pulumi.Input<string>;
    /**
     * The azure key vault URI of the database if it's configured with per Database Customer Managed Keys.
     */
    encryptionProtector?: pulumi.Input<string>;
    /**
     * The flag to enable or disable auto rotation of database encryption protector AKV key.
     */
    encryptionProtectorAutoRotation?: pulumi.Input<boolean>;
    /**
     * The Client id used for cross tenant per database CMK scenario
     */
    federatedClientId?: pulumi.Input<string>;
    /**
     * Specifies the behavior when monthly free limits are exhausted for the free database.
     * 
     * AutoPause: The database will be auto paused upon exhaustion of free limits for remainder of the month.
     * 
     * BillForUsage: The database will continue to be online upon exhaustion of free limits and any overage will be billed.
     */
    freeLimitExhaustionBehavior?: pulumi.Input<string | enums.sql.FreeLimitExhaustionBehavior>;
    /**
     * The number of secondary replicas associated with the Business Critical, Premium, or Hyperscale edition database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
     */
    highAvailabilityReplicaCount?: pulumi.Input<number>;
    /**
     * The Azure Active Directory identity of the database.
     */
    identity?: pulumi.Input<inputs.sql.DatabaseIdentityArgs>;
    /**
     * Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
     */
    isLedgerOn?: pulumi.Input<boolean>;
    /**
     * The resource ids of the user assigned identities to use
     */
    keys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
     */
    licenseType?: pulumi.Input<string | enums.sql.DatabaseLicenseType>;
    /**
     * Resource location.
     */
    location?: pulumi.Input<string>;
    /**
     * The resource identifier of the long term retention backup associated with create operation of this database.
     */
    longTermRetentionBackupResourceId?: pulumi.Input<string>;
    /**
     * Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
     */
    maintenanceConfigurationId?: pulumi.Input<string>;
    /**
     * Whether or not customer controlled manual cutover needs to be done during Update Database operation to Hyperscale tier.
     * 
     * This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier.
     * 
     * When manualCutover is specified, the scaling operation will wait for user input to trigger cutover to Hyperscale database.
     * 
     * To trigger cutover, please provide 'performCutover' parameter when the Scaling operation is in Waiting state.
     */
    manualCutover?: pulumi.Input<boolean>;
    /**
     * The max size of the database expressed in bytes.
     */
    maxSizeBytes?: pulumi.Input<number>;
    /**
     * Minimal capacity that database will always have allocated, if not paused
     */
    minCapacity?: pulumi.Input<number>;
    /**
     * To trigger customer controlled manual cutover during the wait state while Scaling operation is in progress.
     * 
     * This property parameter is only applicable for scaling operations that are initiated along with 'manualCutover' parameter.
     * 
     * This property is only applicable when scaling database from Business Critical/General Purpose/Premium/Standard tier to Hyperscale tier is already in progress.
     * 
     * When performCutover is specified, the scaling operation will trigger cutover and perform role-change to Hyperscale database.
     */
    performCutover?: pulumi.Input<boolean>;
    /**
     * Type of enclave requested on the database i.e. Default or VBS enclaves.
     */
    preferredEnclaveType?: pulumi.Input<string | enums.sql.AlwaysEncryptedEnclaveType>;
    /**
     * The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
     */
    readScale?: pulumi.Input<string | enums.sql.DatabaseReadScale>;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    recoverableDatabaseId?: pulumi.Input<string>;
    /**
     * The resource identifier of the recovery point associated with create operation of this database.
     */
    recoveryServicesRecoveryPointId?: pulumi.Input<string>;
    /**
     * The storage account type to be used to store backups for this database.
     */
    requestedBackupStorageRedundancy?: pulumi.Input<string | enums.sql.BackupStorageRedundancy>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The resource identifier of the restorable dropped database associated with create operation of this database.
     */
    restorableDroppedDatabaseId?: pulumi.Input<string>;
    /**
     * Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    restorePointInTime?: pulumi.Input<string>;
    /**
     * The name of the sample schema to apply when creating this database.
     */
    sampleName?: pulumi.Input<string | enums.sql.SampleName>;
    /**
     * The secondary type of the database if it is a secondary.  Valid values are Geo, Named and Standby.
     */
    secondaryType?: pulumi.Input<string | enums.sql.SecondaryType>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
    /**
     * The database SKU.
     * 
     * The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
     * 
     * ```azurecli
     * az sql db list-editions -l <location> -o table
     * ````
     * 
     * ```powershell
     * Get-AzSqlServerServiceObjective -Location <location>
     * ````
     */
    sku?: pulumi.Input<inputs.sql.SkuArgs>;
    /**
     * Specifies the time that the database was deleted.
     */
    sourceDatabaseDeletionDate?: pulumi.Input<string>;
    /**
     * The resource identifier of the source database associated with create operation of this database.
     */
    sourceDatabaseId?: pulumi.Input<string>;
    /**
     * The resource identifier of the source associated with the create operation of this database.
     * 
     * This property is only supported for DataWarehouse edition and allows to restore across subscriptions.
     * 
     * When sourceResourceId is specified, sourceDatabaseId, recoverableDatabaseId, restorableDroppedDatabaseId and sourceDatabaseDeletionDate must not be specified and CreateMode must be PointInTimeRestore, Restore or Recover.
     * 
     * When createMode is PointInTimeRestore, sourceResourceId must be the resource ID of the existing database or existing sql pool, and restorePointInTime must be specified.
     * 
     * When createMode is Restore, sourceResourceId must be the resource ID of restorable dropped database or restorable dropped sql pool.
     * 
     * When createMode is Recover, sourceResourceId must be the resource ID of recoverable database or recoverable sql pool.
     * 
     * When source subscription belongs to a different tenant than target subscription, “x-ms-authorization-auxiliary” header must contain authentication token for the source tenant. For more details about “x-ms-authorization-auxiliary” header see https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/authenticate-multi-tenant 
     */
    sourceResourceId?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether or not the database uses free monthly limits. Allowed on one database in a subscription.
     */
    useFreeLimit?: pulumi.Input<boolean>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    zoneRedundant?: pulumi.Input<boolean>;
}
