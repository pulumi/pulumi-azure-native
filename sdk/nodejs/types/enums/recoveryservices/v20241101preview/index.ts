// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AcquireStorageAccountLock = {
    Acquire: "Acquire",
    NotAcquire: "NotAcquire",
} as const;

/**
 * Whether storage account lock is to be acquired for this container or not.
 */
export type AcquireStorageAccountLock = (typeof AcquireStorageAccountLock)[keyof typeof AcquireStorageAccountLock];

export const BackupItemType = {
    Invalid: "Invalid",
    VM: "VM",
    FileFolder: "FileFolder",
    AzureSqlDb: "AzureSqlDb",
    SQLDB: "SQLDB",
    Exchange: "Exchange",
    Sharepoint: "Sharepoint",
    VMwareVM: "VMwareVM",
    SystemState: "SystemState",
    Client: "Client",
    GenericDataSource: "GenericDataSource",
    SQLDataBase: "SQLDataBase",
    AzureFileShare: "AzureFileShare",
    SAPHanaDatabase: "SAPHanaDatabase",
    SAPAseDatabase: "SAPAseDatabase",
    SAPHanaDBInstance: "SAPHanaDBInstance",
} as const;

/**
 * Type of backup items associated with this container.
 */
export type BackupItemType = (typeof BackupItemType)[keyof typeof BackupItemType];

export const BackupManagementType = {
    Invalid: "Invalid",
    AzureIaasVM: "AzureIaasVM",
    MAB: "MAB",
    DPM: "DPM",
    AzureBackupServer: "AzureBackupServer",
    AzureSql: "AzureSql",
    AzureStorage: "AzureStorage",
    AzureWorkload: "AzureWorkload",
    DefaultBackup: "DefaultBackup",
} as const;

/**
 * Type of backup management for the backed up item.
 */
export type BackupManagementType = (typeof BackupManagementType)[keyof typeof BackupManagementType];

export const CreateMode = {
    Invalid: "Invalid",
    Default: "Default",
    Recover: "Recover",
} as const;

/**
 * Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
 */
export type CreateMode = (typeof CreateMode)[keyof typeof CreateMode];

export const DayOfWeek = {
    Sunday: "Sunday",
    Monday: "Monday",
    Tuesday: "Tuesday",
    Wednesday: "Wednesday",
    Thursday: "Thursday",
    Friday: "Friday",
    Saturday: "Saturday",
} as const;

export type DayOfWeek = (typeof DayOfWeek)[keyof typeof DayOfWeek];

export const IAASVMPolicyType = {
    Invalid: "Invalid",
    V1: "V1",
    V2: "V2",
} as const;

export type IAASVMPolicyType = (typeof IAASVMPolicyType)[keyof typeof IAASVMPolicyType];

export const IaasVMSnapshotConsistencyType = {
    OnlyCrashConsistent: "OnlyCrashConsistent",
} as const;

export type IaasVMSnapshotConsistencyType = (typeof IaasVMSnapshotConsistencyType)[keyof typeof IaasVMSnapshotConsistencyType];

export const LastBackupStatus = {
    Invalid: "Invalid",
    Healthy: "Healthy",
    Unhealthy: "Unhealthy",
    IRPending: "IRPending",
} as const;

/**
 * Last backup operation status. Possible values: Healthy, Unhealthy.
 */
export type LastBackupStatus = (typeof LastBackupStatus)[keyof typeof LastBackupStatus];

export const MonthOfYear = {
    Invalid: "Invalid",
    January: "January",
    February: "February",
    March: "March",
    April: "April",
    May: "May",
    June: "June",
    July: "July",
    August: "August",
    September: "September",
    October: "October",
    November: "November",
    December: "December",
} as const;

export type MonthOfYear = (typeof MonthOfYear)[keyof typeof MonthOfYear];

export const OperationType = {
    Invalid: "Invalid",
    Register: "Register",
    Reregister: "Reregister",
    Rehydrate: "Rehydrate",
} as const;

/**
 * Re-Do Operation
 */
export type OperationType = (typeof OperationType)[keyof typeof OperationType];

export const PolicyType = {
    Invalid: "Invalid",
    Full: "Full",
    Differential: "Differential",
    Log: "Log",
    CopyOnlyFull: "CopyOnlyFull",
    Incremental: "Incremental",
    SnapshotFull: "SnapshotFull",
    SnapshotCopyOnlyFull: "SnapshotCopyOnlyFull",
} as const;

/**
 * Type of backup policy type
 */
export type PolicyType = (typeof PolicyType)[keyof typeof PolicyType];

export const PrivateEndpointConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
    Disconnected: "Disconnected",
} as const;

/**
 * Gets or sets the status
 */
export type PrivateEndpointConnectionStatus = (typeof PrivateEndpointConnectionStatus)[keyof typeof PrivateEndpointConnectionStatus];

export const ProtectableContainerType = {
    Invalid: "Invalid",
    Unknown: "Unknown",
    IaasVMContainer: "IaasVMContainer",
    IaasVMServiceContainer: "IaasVMServiceContainer",
    DPMContainer: "DPMContainer",
    AzureBackupServerContainer: "AzureBackupServerContainer",
    MABContainer: "MABContainer",
    Cluster: "Cluster",
    AzureSqlContainer: "AzureSqlContainer",
    Windows: "Windows",
    VCenter: "VCenter",
    VMAppContainer: "VMAppContainer",
    SQLAGWorkLoadContainer: "SQLAGWorkLoadContainer",
    StorageContainer: "StorageContainer",
    GenericContainer: "GenericContainer",
    Microsoft_ClassicCompute_virtualMachines: "Microsoft.ClassicCompute/virtualMachines",
    Microsoft_Compute_virtualMachines: "Microsoft.Compute/virtualMachines",
    AzureWorkloadContainer: "AzureWorkloadContainer",
} as const;

/**
 * Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
 * Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
 * Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
 * Backup is VMAppContainer
 */
export type ProtectableContainerType = (typeof ProtectableContainerType)[keyof typeof ProtectableContainerType];

export const ProtectedItemHealthStatus = {
    Invalid: "Invalid",
    Healthy: "Healthy",
    Unhealthy: "Unhealthy",
    NotReachable: "NotReachable",
    IRPending: "IRPending",
} as const;

/**
 * Health status of the backup item, evaluated based on last heartbeat received
 */
export type ProtectedItemHealthStatus = (typeof ProtectedItemHealthStatus)[keyof typeof ProtectedItemHealthStatus];

export const ProtectedItemState = {
    Invalid: "Invalid",
    IRPending: "IRPending",
    Protected: "Protected",
    ProtectionError: "ProtectionError",
    ProtectionStopped: "ProtectionStopped",
    ProtectionPaused: "ProtectionPaused",
    BackupsSuspended: "BackupsSuspended",
} as const;

/**
 * Protection state of the backup engine
 */
export type ProtectedItemState = (typeof ProtectedItemState)[keyof typeof ProtectedItemState];

export const ProtectionIntentItemType = {
    Invalid: "Invalid",
    AzureResourceItem: "AzureResourceItem",
    RecoveryServiceVaultItem: "RecoveryServiceVaultItem",
    AzureWorkloadContainerAutoProtectionIntent: "AzureWorkloadContainerAutoProtectionIntent",
    AzureWorkloadAutoProtectionIntent: "AzureWorkloadAutoProtectionIntent",
    AzureWorkloadSQLAutoProtectionIntent: "AzureWorkloadSQLAutoProtectionIntent",
} as const;

/**
 * backup protectionIntent type.
 */
export type ProtectionIntentItemType = (typeof ProtectionIntentItemType)[keyof typeof ProtectionIntentItemType];

export const ProtectionState = {
    Invalid: "Invalid",
    IRPending: "IRPending",
    Protected: "Protected",
    ProtectionError: "ProtectionError",
    ProtectionStopped: "ProtectionStopped",
    ProtectionPaused: "ProtectionPaused",
    BackupsSuspended: "BackupsSuspended",
} as const;

/**
 * Backup state of this backup item.
 */
export type ProtectionState = (typeof ProtectionState)[keyof typeof ProtectionState];

export const ProtectionStatus = {
    Invalid: "Invalid",
    NotProtected: "NotProtected",
    Protecting: "Protecting",
    Protected: "Protected",
    ProtectionFailed: "ProtectionFailed",
} as const;

/**
 * Backup state of this backup item.
 */
export type ProtectionStatus = (typeof ProtectionStatus)[keyof typeof ProtectionStatus];

export const ProvisioningState = {
    Succeeded: "Succeeded",
    Deleting: "Deleting",
    Failed: "Failed",
    Pending: "Pending",
} as const;

/**
 * Gets or sets provisioning state of the private endpoint connection
 */
export type ProvisioningState = (typeof ProvisioningState)[keyof typeof ProvisioningState];

export const ResourceHealthStatus = {
    Healthy: "Healthy",
    TransientDegraded: "TransientDegraded",
    PersistentDegraded: "PersistentDegraded",
    TransientUnhealthy: "TransientUnhealthy",
    PersistentUnhealthy: "PersistentUnhealthy",
    Invalid: "Invalid",
} as const;

/**
 * Resource Health Status
 */
export type ResourceHealthStatus = (typeof ResourceHealthStatus)[keyof typeof ResourceHealthStatus];

export const RetentionDurationType = {
    Invalid: "Invalid",
    Days: "Days",
    Weeks: "Weeks",
    Months: "Months",
    Years: "Years",
} as const;

/**
 * Retention duration type: days/weeks/months/years
 * Used only if TieringMode is set to TierAfter
 */
export type RetentionDurationType = (typeof RetentionDurationType)[keyof typeof RetentionDurationType];

export const RetentionScheduleFormat = {
    Invalid: "Invalid",
    Daily: "Daily",
    Weekly: "Weekly",
} as const;

/**
 * Retention schedule format for yearly retention policy.
 */
export type RetentionScheduleFormat = (typeof RetentionScheduleFormat)[keyof typeof RetentionScheduleFormat];

export const ScheduleRunType = {
    Invalid: "Invalid",
    Daily: "Daily",
    Weekly: "Weekly",
    Hourly: "Hourly",
} as const;

/**
 * Frequency of the schedule operation of this policy.
 */
export type ScheduleRunType = (typeof ScheduleRunType)[keyof typeof ScheduleRunType];

export const TieringMode = {
    Invalid: "Invalid",
    TierRecommended: "TierRecommended",
    TierAfter: "TierAfter",
    DoNotTier: "DoNotTier",
} as const;

/**
 * Tiering Mode to control automatic tiering of recovery points. Supported values are:
 * 1. TierRecommended: Tier all recovery points recommended to be tiered
 * 2. TierAfter: Tier all recovery points after a fixed period, as specified in duration + durationType below.
 * 3. DoNotTier: Do not tier any recovery points
 */
export type TieringMode = (typeof TieringMode)[keyof typeof TieringMode];

export const VMWorkloadPolicyType = {
    Invalid: "Invalid",
    SnapshotV1: "SnapshotV1",
    SnapshotV2: "SnapshotV2",
    Streaming: "Streaming",
} as const;

/**
 * Type of workload for the backup management
 */
export type VMWorkloadPolicyType = (typeof VMWorkloadPolicyType)[keyof typeof VMWorkloadPolicyType];

export const VaultSubResourceType = {
    AzureBackup: "AzureBackup",
    AzureBackup_secondary: "AzureBackup_secondary",
    AzureSiteRecovery: "AzureSiteRecovery",
} as const;

/**
 * GroupId for the PrivateEndpointConnection - AzureBackup, AzureBackup_secondary or AzureSiteRecovery
 */
export type VaultSubResourceType = (typeof VaultSubResourceType)[keyof typeof VaultSubResourceType];

export const WeekOfMonth = {
    First: "First",
    Second: "Second",
    Third: "Third",
    Fourth: "Fourth",
    Last: "Last",
    Invalid: "Invalid",
} as const;

export type WeekOfMonth = (typeof WeekOfMonth)[keyof typeof WeekOfMonth];

export const WorkloadItemType = {
    Invalid: "Invalid",
    SQLInstance: "SQLInstance",
    SQLDataBase: "SQLDataBase",
    SAPHanaSystem: "SAPHanaSystem",
    SAPHanaDatabase: "SAPHanaDatabase",
    SAPAseSystem: "SAPAseSystem",
    SAPAseDatabase: "SAPAseDatabase",
    SAPHanaDBInstance: "SAPHanaDBInstance",
} as const;

/**
 * Workload item type of the item for which intent is to be set
 */
export type WorkloadItemType = (typeof WorkloadItemType)[keyof typeof WorkloadItemType];

export const WorkloadType = {
    Invalid: "Invalid",
    VM: "VM",
    FileFolder: "FileFolder",
    AzureSqlDb: "AzureSqlDb",
    SQLDB: "SQLDB",
    Exchange: "Exchange",
    Sharepoint: "Sharepoint",
    VMwareVM: "VMwareVM",
    SystemState: "SystemState",
    Client: "Client",
    GenericDataSource: "GenericDataSource",
    SQLDataBase: "SQLDataBase",
    AzureFileShare: "AzureFileShare",
    SAPHanaDatabase: "SAPHanaDatabase",
    SAPAseDatabase: "SAPAseDatabase",
    SAPHanaDBInstance: "SAPHanaDBInstance",
} as const;

/**
 * Type of workload for the backup management
 */
export type WorkloadType = (typeof WorkloadType)[keyof typeof WorkloadType];
