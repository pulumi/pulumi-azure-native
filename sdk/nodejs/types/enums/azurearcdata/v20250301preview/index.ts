// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccountProvisioningMode = {
    Automatic: "automatic",
    Manual: "manual",
} as const;

/**
 * The service account provisioning mode for this Active Directory connector.
 */
export type AccountProvisioningMode = (typeof AccountProvisioningMode)[keyof typeof AccountProvisioningMode];

export const ActivationState = {
    Activated: "Activated",
    Deactivated: "Deactivated",
} as const;

/**
 * The activation state of the license.
 */
export type ActivationState = (typeof ActivationState)[keyof typeof ActivationState];

export const AggregationType = {
    Average: "Average",
    Minimum: "Minimum",
    Maximum: "Maximum",
    Sum: "Sum",
    Count: "Count",
} as const;

/**
 * The aggregation type to use for the numerical columns in the dataset.
 */
export type AggregationType = (typeof AggregationType)[keyof typeof AggregationType];

export const ArcSqlManagedInstanceLicenseType = {
    BasePrice: "BasePrice",
    LicenseIncluded: "LicenseIncluded",
    DisasterRecovery: "DisasterRecovery",
} as const;

/**
 * The license type to apply for this managed instance.
 */
export type ArcSqlManagedInstanceLicenseType = (typeof ArcSqlManagedInstanceLicenseType)[keyof typeof ArcSqlManagedInstanceLicenseType];

export const ArcSqlServerAvailabilityMode = {
    SYNCHRONOUS_COMMIT: "SYNCHRONOUS_COMMIT",
    ASYNCHRONOUS_COMMIT: "ASYNCHRONOUS_COMMIT",
} as const;

/**
 * Property that determines whether a given availability replica can run in synchronous-commit mode
 */
export type ArcSqlServerAvailabilityMode = (typeof ArcSqlServerAvailabilityMode)[keyof typeof ArcSqlServerAvailabilityMode];

export const ArcSqlServerFailoverMode = {
    AUTOMATIC: "AUTOMATIC",
    MANUAL: "MANUAL",
    EXTERNAL: "EXTERNAL",
    NONE: "NONE",
} as const;

/**
 * Property to set the failover mode of the availability group replica
 */
export type ArcSqlServerFailoverMode = (typeof ArcSqlServerFailoverMode)[keyof typeof ArcSqlServerFailoverMode];

export const BillingPlan = {
    PAYG: "PAYG",
    Paid: "Paid",
} as const;

/**
 * SQL Server license type.
 */
export type BillingPlan = (typeof BillingPlan)[keyof typeof BillingPlan];

export const ConnectionAuth = {
    Windows_NTLM: "Windows_NTLM",
    Windows_Kerberos: "Windows_Kerberos",
    Windows_Negotiate: "Windows_Negotiate",
    Certificate: "Certificate",
    Windows_NTLM_Certificate: "Windows_NTLM_Certificate",
    Windows_Kerberos_Certificate: "Windows_Kerberos_Certificate",
    Windows_Negotiate_Certificate: "Windows_Negotiate_Certificate",
    Certificate_Windows_NTLM: "Certificate_Windows_NTLM",
    Certificate_Windows_Kerberos: "Certificate_Windows_Kerberos",
    Certificate_Windows_Negotiate: "Certificate_Windows_Negotiate",
} as const;

/**
 * Permitted authentication modes for the mirroring endpoint.
 */
export type ConnectionAuth = (typeof ConnectionAuth)[keyof typeof ConnectionAuth];

export const DatabaseCreateMode = {
    Default: "Default",
    PointInTimeRestore: "PointInTimeRestore",
} as const;

/**
 * Database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. sourceDatabaseId and restorePointInTime must be specified.
 */
export type DatabaseCreateMode = (typeof DatabaseCreateMode)[keyof typeof DatabaseCreateMode];

export const DatabaseState = {
    Online: "Online",
    Restoring: "Restoring",
    Recovering: "Recovering",
    RecoveryPending: "RecoveryPending",
    Suspect: "Suspect",
    Emergency: "Emergency",
    Offline: "Offline",
    Copying: "Copying",
    OfflineSecondary: "OfflineSecondary",
} as const;

/**
 * State of the database.
 */
export type DatabaseState = (typeof DatabaseState)[keyof typeof DatabaseState];

export const EditionType = {
    Evaluation: "Evaluation",
    Enterprise: "Enterprise",
    Standard: "Standard",
    Web: "Web",
    Developer: "Developer",
    Express: "Express",
    Business_Intelligence: "Business Intelligence",
} as const;

/**
 * SQL Server edition.
 */
export type EditionType = (typeof EditionType)[keyof typeof EditionType];

export const ExtendedLocationTypes = {
    CustomLocation: "CustomLocation",
} as const;

/**
 * The type of the extended location.
 */
export type ExtendedLocationTypes = (typeof ExtendedLocationTypes)[keyof typeof ExtendedLocationTypes];

export const FailoverGroupPartnerSyncMode = {
    Async: "async",
    Sync: "sync",
} as const;

/**
 * The partner sync mode of the SQL managed instance.
 */
export type FailoverGroupPartnerSyncMode = (typeof FailoverGroupPartnerSyncMode)[keyof typeof FailoverGroupPartnerSyncMode];

export const HostType = {
    Azure_Virtual_Machine: "Azure Virtual Machine",
    Azure_VMWare_Virtual_Machine: "Azure VMWare Virtual Machine",
    Azure_Kubernetes_Service: "Azure Kubernetes Service",
    AWS_VMWare_Virtual_Machine: "AWS VMWare Virtual Machine",
    AWS_Kubernetes_Service: "AWS Kubernetes Service",
    GCP_VMWare_Virtual_Machine: "GCP VMWare Virtual Machine",
    GCP_Kubernetes_Service: "GCP Kubernetes Service",
    Container: "Container",
    Virtual_Machine: "Virtual Machine",
    Physical_Server: "Physical Server",
    AWS_Virtual_Machine: "AWS Virtual Machine",
    GCP_Virtual_Machine: "GCP Virtual Machine",
    Other: "Other",
} as const;

/**
 * Type of host for Azure Arc SQL Server
 */
export type HostType = (typeof HostType)[keyof typeof HostType];

export const IdentityType = {
    /**
     * System Assigned Managed Identity
     */
    SystemAssignedManagedIdentity: "SystemAssignedManagedIdentity",
    /**
     * User Assigned Managed Identity
     */
    UserAssignedManagedIdentity: "UserAssignedManagedIdentity",
} as const;

/**
 * The method used for Entra authentication
 */
export type IdentityType = (typeof IdentityType)[keyof typeof IdentityType];

export const Infrastructure = {
    Azure: "azure",
    Gcp: "gcp",
    Aws: "aws",
    Alibaba: "alibaba",
    Onpremises: "onpremises",
    Other: "other",
} as const;

/**
 * The infrastructure the data controller is running on.
 */
export type Infrastructure = (typeof Infrastructure)[keyof typeof Infrastructure];

export const InstanceFailoverGroupRole = {
    Primary: "primary",
    Secondary: "secondary",
    Force_primary_allow_data_loss: "force-primary-allow-data-loss",
    Force_secondary: "force-secondary",
} as const;

/**
 * The role of the SQL managed instance in this failover group.
 */
export type InstanceFailoverGroupRole = (typeof InstanceFailoverGroupRole)[keyof typeof InstanceFailoverGroupRole];

export const LicenseCategory = {
    Core: "Core",
} as const;

/**
 * This property represents the choice between SQL Server Core and ESU licenses.
 */
export type LicenseCategory = (typeof LicenseCategory)[keyof typeof LicenseCategory];

export const Mode = {
    /**
     * Mixed mode authentication for SQL Server which includes windows and SQL Authentication.
     */
    Mixed: "Mixed",
    /**
     * Windows Authentication for SQL Server.
     */
    Windows: "Windows",
    /**
     * Used for scenarios were the mode cannot be determined.
     */
    Undefined: "Undefined",
} as const;

/**
 * Mode of authentication in SqlServer.
 */
export type Mode = (typeof Mode)[keyof typeof Mode];

export const PostgresInstanceSkuTier = {
    Hyperscale: "Hyperscale",
} as const;

/**
 * This field is required to be implemented by the Resource Provider if the service has more than one tier.
 */
export type PostgresInstanceSkuTier = (typeof PostgresInstanceSkuTier)[keyof typeof PostgresInstanceSkuTier];

export const PrimaryAllowConnections = {
    ALL: "ALL",
    READ_WRITE: "READ_WRITE",
} as const;

/**
 * Whether the primary replica should allow all connections or only READ_WRITE connections (disallowing ReadOnly connections)
 */
export type PrimaryAllowConnections = (typeof PrimaryAllowConnections)[keyof typeof PrimaryAllowConnections];

export const RecoveryMode = {
    Full: "Full",
    Bulk_logged: "Bulk-logged",
    Simple: "Simple",
} as const;

/**
 * Status of the database.
 */
export type RecoveryMode = (typeof RecoveryMode)[keyof typeof RecoveryMode];

export const ScopeType = {
    Tenant: "Tenant",
    Subscription: "Subscription",
    ResourceGroup: "ResourceGroup",
} as const;

/**
 * The Azure scope to which the license will apply.
 */
export type ScopeType = (typeof ScopeType)[keyof typeof ScopeType];

export const SecondaryAllowConnections = {
    NO: "NO",
    ALL: "ALL",
    READ_ONLY: "READ_ONLY",
} as const;

/**
 * Whether the secondary replica should allow all connections, no connections, or only ReadOnly connections.
 */
export type SecondaryAllowConnections = (typeof SecondaryAllowConnections)[keyof typeof SecondaryAllowConnections];

export const SeedingMode = {
    AUTOMATIC: "AUTOMATIC",
    MANUAL: "MANUAL",
} as const;

/**
 * Specifies how the secondary replica will be initially seeded. AUTOMATIC enables direct seeding. This method will seed the secondary replica over the network. This method does not require you to backup and restore a copy of the primary database on the replica. MANUAL specifies manual seeding (default). This method requires you to create a backup of the database on the primary replica and manually restore that backup on the secondary replica.
 */
export type SeedingMode = (typeof SeedingMode)[keyof typeof SeedingMode];

export const ServiceType = {
    /**
     * SQL Server Database Services.
     */
    Engine: "Engine",
    /**
     * SQL Server Reporting Services.
     */
    SSRS: "SSRS",
    /**
     * SQL Server Analysis Services.
     */
    SSAS: "SSAS",
    /**
     * SQL Server Integration Services.
     */
    SSIS: "SSIS",
    /**
     * Power BI Report Server.
     */
    PBIRS: "PBIRS",
} as const;

/**
 * Indicates if the resource represents a SQL Server engine or a SQL Server component service installed on the host.
 */
export type ServiceType = (typeof ServiceType)[keyof typeof ServiceType];

export const SqlManagedInstanceSkuName = {
    VCore: "vCore",
} as const;

/**
 * The name of the SKU.
 */
export type SqlManagedInstanceSkuName = (typeof SqlManagedInstanceSkuName)[keyof typeof SqlManagedInstanceSkuName];

export const SqlManagedInstanceSkuTier = {
    GeneralPurpose: "GeneralPurpose",
    BusinessCritical: "BusinessCritical",
} as const;

/**
 * The pricing tier for the instance.
 */
export type SqlManagedInstanceSkuTier = (typeof SqlManagedInstanceSkuTier)[keyof typeof SqlManagedInstanceSkuTier];

export const SqlVersion = {
    SQL_Server_2012: "SQL Server 2012",
    SQL_Server_2014: "SQL Server 2014",
    SQL_Server_2016: "SQL Server 2016",
    SQL_Server_2017: "SQL Server 2017",
    SQL_Server_2019: "SQL Server 2019",
    SQL_Server_2022: "SQL Server 2022",
    Unknown: "Unknown",
} as const;

/**
 * SQL Server version.
 */
export type SqlVersion = (typeof SqlVersion)[keyof typeof SqlVersion];

export const State = {
    Inactive: "Inactive",
    Active: "Active",
    Terminated: "Terminated",
} as const;

/**
 * The activation state of the license.
 */
export type State = (typeof State)[keyof typeof State];

export const Version = {
    SQL_Server_2012: "SQL Server 2012",
    SQL_Server_2014: "SQL Server 2014",
} as const;

/**
 * The SQL Server version the license covers.
 */
export type Version = (typeof Version)[keyof typeof Version];
