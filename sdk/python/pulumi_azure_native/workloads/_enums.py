# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AlertAutoMitigate',
    'AlertRuleStatus',
    'AppServicePlanTier',
    'BackupType',
    'ConditionalOperator',
    'DayOfWeek',
    'DiskSkuName',
    'FileShareConfigurationType',
    'IAASVMPolicyType',
    'ManagedResourcesNetworkAccessType',
    'ManagedServiceIdentityType',
    'MonthOfYear',
    'NamingPatternType',
    'OSType',
    'PolicyType',
    'RetentionDurationType',
    'RetentionScheduleFormat',
    'RoutingPreference',
    'SAPConfigurationType',
    'SAPDatabaseScaleMethod',
    'SAPDatabaseType',
    'SAPDeploymentType',
    'SAPEnvironmentType',
    'SAPHighAvailabilityType',
    'SAPProductType',
    'SAPSoftwareInstallationType',
    'SAPVirtualInstanceIdentityType',
    'ScheduleRunType',
    'SslCryptoProvider',
    'SslPreference',
    'TieringMode',
    'VaultType',
    'WeekOfMonth',
    'WorkloadType',
]


@pulumi.type_token("azure-native:workloads:AlertAutoMitigate")
class AlertAutoMitigate(builtins.str, Enum):
    """
    The value that indicates whether the alert should be automatically resolved or not. The default is Disable.
    """
    ENABLE = "Enable"
    """
    The alert should be automatically resolved.
    """
    DISABLE = "Disable"
    """
    The alert should not be automatically resolved.
    """


@pulumi.type_token("azure-native:workloads:AlertRuleStatus")
class AlertRuleStatus(builtins.str, Enum):
    """
    Indicates whether the alert is in an enabled state.
    """
    ENABLED = "Enabled"
    """
    The scheduled query rule is enabled
    """
    DISABLED = "Disabled"
    """
    The scheduled query rule is disabled
    """


@pulumi.type_token("azure-native:workloads:AppServicePlanTier")
class AppServicePlanTier(builtins.str, Enum):
    """
    The App Service plan tier.
    """
    ELASTIC_PREMIUM = "ElasticPremium"
    """
    Elastic Premium plan
    """
    PREMIUM_V3 = "PremiumV3"
    """
    Dedicated Premium V3 plan
    """


@pulumi.type_token("azure-native:workloads:BackupType")
class BackupType(builtins.str, Enum):
    """
    The type of backup, VM, SQL or HANA.
    """
    VM = "VM"
    SQL = "SQL"
    HANA = "HANA"


@pulumi.type_token("azure-native:workloads:ConditionalOperator")
class ConditionalOperator(builtins.str, Enum):
    """
    The threshold operator of the alert.
    """
    LESS_THAN = "LessThan"
    """
    The value is less than the specified value.
    """
    GREATER_THAN = "GreaterThan"
    """
    The value is greater than the specified value.
    """
    EQUAL = "Equal"
    """
    The value is equal to the specified value.
    """
    GREATER_THAN_OR_EQUAL = "GreaterThanOrEqual"
    """
    The value is greater than or equal to the specified value.
    """
    LESS_THAN_OR_EQUAL = "LessThanOrEqual"
    """
    The value is less than or equal to the specified value.
    """


@pulumi.type_token("azure-native:workloads:DayOfWeek")
class DayOfWeek(builtins.str, Enum):
    SUNDAY = "Sunday"
    MONDAY = "Monday"
    TUESDAY = "Tuesday"
    WEDNESDAY = "Wednesday"
    THURSDAY = "Thursday"
    FRIDAY = "Friday"
    SATURDAY = "Saturday"


@pulumi.type_token("azure-native:workloads:DiskSkuName")
class DiskSkuName(builtins.str, Enum):
    """
    Defines the disk sku name.
    """
    STANDARD_LRS = "Standard_LRS"
    """
    Standard LRS Disk SKU.
    """
    PREMIUM_LRS = "Premium_LRS"
    """
    Premium_LRS Disk SKU.
    """
    STANDARD_SS_D_LRS = "StandardSSD_LRS"
    """
    StandardSSD_LRS Disk SKU.
    """
    ULTRA_SS_D_LRS = "UltraSSD_LRS"
    """
    UltraSSD_LRS Disk SKU.
    """
    PREMIUM_ZRS = "Premium_ZRS"
    """
    Premium_ZRS Disk SKU.
    """
    STANDARD_SS_D_ZRS = "StandardSSD_ZRS"
    """
    StandardSSD_ZRS Disk SKU.
    """
    PREMIUM_V2_LRS = "PremiumV2_LRS"
    """
    PremiumV2_LRS Disk SKU.
    """


@pulumi.type_token("azure-native:workloads:FileShareConfigurationType")
class FileShareConfigurationType(builtins.str, Enum):
    """
    The type of file share config, eg: Mount/CreateAndMount/Skip.
    """
    SKIP = "Skip"
    """
    Skip creating the file share.
    """
    CREATE_AND_MOUNT = "CreateAndMount"
    """
    Fileshare will be created and mounted by service.
    """
    MOUNT = "Mount"
    """
    Existing fileshare provided will be mounted by service.
    """


@pulumi.type_token("azure-native:workloads:IAASVMPolicyType")
class IAASVMPolicyType(builtins.str, Enum):
    """
    The policy type.
    """
    INVALID = "Invalid"
    V1 = "V1"
    V2 = "V2"


@pulumi.type_token("azure-native:workloads:ManagedResourcesNetworkAccessType")
class ManagedResourcesNetworkAccessType(builtins.str, Enum):
    """
    Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
    """
    PUBLIC = "Public"
    """
    Managed resources will be deployed with public network access enabled.
    """
    PRIVATE = "Private"
    """
    Managed resources will be deployed with public network access disabled.
    """


@pulumi.type_token("azure-native:workloads:ManagedServiceIdentityType")
class ManagedServiceIdentityType(builtins.str, Enum):
    """
    Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"
    SYSTEM_ASSIGNED_USER_ASSIGNED = "SystemAssigned,UserAssigned"


@pulumi.type_token("azure-native:workloads:MonthOfYear")
class MonthOfYear(builtins.str, Enum):
    INVALID = "Invalid"
    JANUARY = "January"
    FEBRUARY = "February"
    MARCH = "March"
    APRIL = "April"
    MAY = "May"
    JUNE = "June"
    JULY = "July"
    AUGUST = "August"
    SEPTEMBER = "September"
    OCTOBER = "October"
    NOVEMBER = "November"
    DECEMBER = "December"


@pulumi.type_token("azure-native:workloads:NamingPatternType")
class NamingPatternType(builtins.str, Enum):
    """
    The pattern type to be used for resource naming.
    """
    FULL_RESOURCE_NAME = "FullResourceName"
    """
    Full resource names that will be created by service.
    """


@pulumi.type_token("azure-native:workloads:OSType")
class OSType(builtins.str, Enum):
    """
    The OS Type
    """
    LINUX = "Linux"
    """
    Linux OS Type.
    """
    WINDOWS = "Windows"
    """
    Windows OS Type.
    """


@pulumi.type_token("azure-native:workloads:PolicyType")
class PolicyType(builtins.str, Enum):
    """
    Type of backup policy type
    """
    INVALID = "Invalid"
    FULL = "Full"
    DIFFERENTIAL = "Differential"
    LOG = "Log"
    COPY_ONLY_FULL = "CopyOnlyFull"
    INCREMENTAL = "Incremental"
    SNAPSHOT_FULL = "SnapshotFull"
    SNAPSHOT_COPY_ONLY_FULL = "SnapshotCopyOnlyFull"


@pulumi.type_token("azure-native:workloads:RetentionDurationType")
class RetentionDurationType(builtins.str, Enum):
    """
    Retention duration type: days/weeks/months/years
    Used only if TieringMode is set to TierAfter
    """
    INVALID = "Invalid"
    DAYS = "Days"
    WEEKS = "Weeks"
    MONTHS = "Months"
    YEARS = "Years"


@pulumi.type_token("azure-native:workloads:RetentionScheduleFormat")
class RetentionScheduleFormat(builtins.str, Enum):
    """
    Retention schedule format for yearly retention policy.
    """
    INVALID = "Invalid"
    DAILY = "Daily"
    WEEKLY = "Weekly"


@pulumi.type_token("azure-native:workloads:RoutingPreference")
class RoutingPreference(builtins.str, Enum):
    """
    Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
    """
    DEFAULT = "Default"
    """
    Default routing preference. Only RFC1918 traffic is routed to the customer VNET.
    """
    ROUTE_ALL = "RouteAll"
    """
    Route all traffic to the customer VNET.
    """


@pulumi.type_token("azure-native:workloads:SAPConfigurationType")
class SAPConfigurationType(builtins.str, Enum):
    """
    The configuration type. Eg: Deployment/Discovery
    """
    DEPLOYMENT = "Deployment"
    """
    SAP system will be deployed by service. No OS configurations will be done.
    """
    DISCOVERY = "Discovery"
    """
    Existing SAP system will be registered.
    """
    DEPLOYMENT_WITH_OS_CONFIG = "DeploymentWithOSConfig"
    """
    SAP system will be deployed by service. OS configurations will be done.
    """


@pulumi.type_token("azure-native:workloads:SAPDatabaseScaleMethod")
class SAPDatabaseScaleMethod(builtins.str, Enum):
    """
    The DB scale method.
    """
    SCALE_UP = "ScaleUp"
    """
    ScaleUp Hana Database deployment type
    """


@pulumi.type_token("azure-native:workloads:SAPDatabaseType")
class SAPDatabaseType(builtins.str, Enum):
    """
    The database type.
    """
    HANA = "HANA"
    """
    HANA Database type of SAP system.
    """
    DB2 = "DB2"
    """
    DB2 database type of the SAP system.
    """


@pulumi.type_token("azure-native:workloads:SAPDeploymentType")
class SAPDeploymentType(builtins.str, Enum):
    """
    The deployment type. Eg: SingleServer/ThreeTier
    """
    SINGLE_SERVER = "SingleServer"
    """
    SAP Single server deployment type.
    """
    THREE_TIER = "ThreeTier"
    """
    SAP Distributed deployment type.
    """


@pulumi.type_token("azure-native:workloads:SAPEnvironmentType")
class SAPEnvironmentType(builtins.str, Enum):
    """
    Defines the environment type - Production/Non Production.
    """
    NON_PROD = "NonProd"
    """
    Non Production SAP system.
    """
    PROD = "Prod"
    """
    Production SAP system.
    """


@pulumi.type_token("azure-native:workloads:SAPHighAvailabilityType")
class SAPHighAvailabilityType(builtins.str, Enum):
    """
    The high availability type.
    """
    AVAILABILITY_SET = "AvailabilitySet"
    """
    HA deployment with availability sets.
    """
    AVAILABILITY_ZONE = "AvailabilityZone"
    """
    HA deployment with availability zones.
    """


@pulumi.type_token("azure-native:workloads:SAPProductType")
class SAPProductType(builtins.str, Enum):
    """
    Defines the SAP Product type.
    """
    ECC = "ECC"
    """
    SAP Product ECC.
    """
    S4_HANA = "S4HANA"
    """
    SAP Product S4HANA.
    """
    OTHER = "Other"
    """
    SAP Products other than the ones listed.
    """


@pulumi.type_token("azure-native:workloads:SAPSoftwareInstallationType")
class SAPSoftwareInstallationType(builtins.str, Enum):
    """
    The SAP software installation type.
    """
    SERVICE_INITIATED = "ServiceInitiated"
    """
    SAP Install managed by service.
    """
    SAP_INSTALL_WITHOUT_OS_CONFIG = "SAPInstallWithoutOSConfig"
    """
    SAP Install without OS Config.
    """
    EXTERNAL = "External"
    """
    External software installation type.
    """


@pulumi.type_token("azure-native:workloads:SAPVirtualInstanceIdentityType")
class SAPVirtualInstanceIdentityType(builtins.str, Enum):
    """
    The type of managed identity assigned to this resource.
    """
    NONE = "None"
    """
    No managed identity.
    """
    USER_ASSIGNED = "UserAssigned"
    """
    User assigned managed identity.
    """


@pulumi.type_token("azure-native:workloads:ScheduleRunType")
class ScheduleRunType(builtins.str, Enum):
    """
    Frequency of the schedule operation of this policy.
    """
    INVALID = "Invalid"
    DAILY = "Daily"
    WEEKLY = "Weekly"
    HOURLY = "Hourly"


@pulumi.type_token("azure-native:workloads:SslCryptoProvider")
class SslCryptoProvider(builtins.str, Enum):
    """
    Specify the crypto provider being used (commoncrypto/openssl). If this argument is not provided, it is automatically determined by searching in the configuration files.
    """
    COMMONCRYPTO = "commoncrypto"
    OPENSSL = "openssl"


@pulumi.type_token("azure-native:workloads:SslPreference")
class SslPreference(builtins.str, Enum):
    """
    Gets or sets certificate preference if secure communication is enabled.
    """
    DISABLED = "Disabled"
    """
    Secure communication is disabled.
    """
    ROOT_CERTIFICATE = "RootCertificate"
    """
    Secure communication is enabled with root certificate.
    """
    SERVER_CERTIFICATE = "ServerCertificate"
    """
    Secure communication is enabled with server certificate.
    """


@pulumi.type_token("azure-native:workloads:TieringMode")
class TieringMode(builtins.str, Enum):
    """
    Tiering Mode to control automatic tiering of recovery points. Supported values are:
    1. TierRecommended: Tier all recovery points recommended to be tiered
    2. TierAfter: Tier all recovery points after a fixed period, as specified in duration + durationType below.
    3. DoNotTier: Do not tier any recovery points
    """
    INVALID = "Invalid"
    TIER_RECOMMENDED = "TierRecommended"
    TIER_AFTER = "TierAfter"
    DO_NOT_TIER = "DoNotTier"


@pulumi.type_token("azure-native:workloads:VaultType")
class VaultType(builtins.str, Enum):
    """
    The vault type, whether it is existing or has to be created.
    """
    EXISTING = "Existing"
    NEW = "New"


@pulumi.type_token("azure-native:workloads:WeekOfMonth")
class WeekOfMonth(builtins.str, Enum):
    FIRST = "First"
    SECOND = "Second"
    THIRD = "Third"
    FOURTH = "Fourth"
    LAST = "Last"
    INVALID = "Invalid"


@pulumi.type_token("azure-native:workloads:WorkloadType")
class WorkloadType(builtins.str, Enum):
    """
    Type of workload for the backup management
    """
    INVALID = "Invalid"
    VM = "VM"
    FILE_FOLDER = "FileFolder"
    AZURE_SQL_DB = "AzureSqlDb"
    SQLDB = "SQLDB"
    EXCHANGE = "Exchange"
    SHAREPOINT = "Sharepoint"
    V_MWARE_VM = "VMwareVM"
    SYSTEM_STATE = "SystemState"
    CLIENT = "Client"
    GENERIC_DATA_SOURCE = "GenericDataSource"
    SQL_DATA_BASE = "SQLDataBase"
    AZURE_FILE_SHARE = "AzureFileShare"
    SAP_HANA_DATABASE = "SAPHanaDatabase"
    SAPASE_DATABASE = "SAPAseDatabase"
    SAP_HANA_DB_INSTANCE = "SAPHanaDBInstance"
